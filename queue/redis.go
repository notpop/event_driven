package queue

import (
	"context"
	"encoding/json"
	"event-driven/common/model"
	"log"

	"github.com/go-redis/redis/v8"
)

var (
	ctx    = context.Background()
	client *redis.Client
)

func init() {
	client = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	_, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}
}

func AddJobToQueue(job model.Job) error {
	data, err := json.Marshal(job)
	if err != nil {
		return err
	}
	err = client.LPush(ctx, "jobQueue", data).Err()
	if err != nil {
		return err
	}

	jobStatus := map[string]interface{}{
		"status": "queued",
		"job":    job,
	}
	statusData, err := json.Marshal(jobStatus)
	if err != nil {
		return err
	}

	return client.Set(ctx, "jobStatus:"+job.ID, statusData, 0).Err()
}

func GetJobFromQueue() (model.Job, error) {
	var job model.Job
	data, err := client.RPop(ctx, "jobQueue").Result()
	if err == redis.Nil {
		return job, nil
	} else if err != nil {
		return job, err
	}
	err = json.Unmarshal([]byte(data), &job)
	return job, err
}

func GetJobStatus(jobID string) (map[string]interface{}, error) {
	var jobStatus map[string]interface{}
	statusData, err := client.Get(ctx, "jobStatus:"+jobID).Result()
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(statusData), &jobStatus)
	return jobStatus, err
}

func GetAllJobs() ([]model.Job, error) {
	var jobs []model.Job
	data, err := client.LRange(ctx, "jobQueue", 0, -1).Result()
	if err != nil {
		return nil, err
	}
	for _, item := range data {
		var job model.Job
		err := json.Unmarshal([]byte(item), &job)
		if err != nil {
			return nil, err
		}
		jobs = append(jobs, job)
	}
	return jobs, nil
}

func UpdateJobStatus(jobID string, status string) error {
	var jobStatus map[string]interface{}
	statusData, err := client.Get(ctx, "jobStatus:"+jobID).Result()
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(statusData), &jobStatus)
	if err != nil {
		return err
	}

	jobStatus["status"] = status
	newStatusData, err := json.Marshal(jobStatus)
	if err != nil {
		return err
	}

	return client.Set(ctx, "jobStatus:"+jobID, newStatusData, 0).Err()
}
