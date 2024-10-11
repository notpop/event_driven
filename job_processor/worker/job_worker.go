package worker

import (
	"bytes"
	"encoding/json"
	"event-driven/common/model"
	"event-driven/queue"
	"log"
	"net/http"
	"time"
)

func StartWorker() {
	log.Println("Job worker started...")
	dispatcher := initDispatcher()

	for {
		rawJob, err := queue.GetJobFromQueue()
		if err != nil {
			log.Println("Error getting job from queue:", err)
			time.Sleep(5 * time.Second)
			continue
		}
		if rawJob.ID == "" {
			time.Sleep(1 * time.Second)
			continue
		}

		go handleJob(rawJob, dispatcher)
	}
}

func handleJob(rawJob model.Job, dispatcher *JobDispatcher) {
	updateJobStatus(rawJob.ID, "processing")

	job, err := dispatcher.Dispatch(rawJob.Type, rawJob.Payload)
	if err != nil {
		log.Println("Error dispatching job:", err)
		updateJobStatus(rawJob.ID, "failed")
		return
	}

	log.Printf("Received job: ID=%s, Payload=%s", rawJob.ID, rawJob.Payload)
	processJob(job, rawJob.ID)
}

func processJob(job model.JobInterface, jobID string) {
	err := job.Process()
	if err != nil {
		log.Println("Error processing job:", err)
		updateJobStatus(jobID, "failed")
	} else {
		updateJobStatus(jobID, "completed")
	}
}

func updateJobStatus(jobID string, status string) {
	err := queue.UpdateJobStatus(jobID, status)
	if err != nil {
		log.Printf("Failed to update job status in Redis: %v", err)
	}

	message := map[string]string{"jobId": jobID, "status": status}
	jsonMessage, _ := json.Marshal(message)

	_, err = http.Post("http://localhost:8080/job/status/update", "application/json", bytes.NewBuffer(jsonMessage))
	if err != nil {
		log.Printf("Failed to send job status update: %v", err)
	}
}
