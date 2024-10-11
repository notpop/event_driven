package jobs

import (
	"event-driven/common/model"
	"fmt"
	"time"
)

type JobTypeA struct {
	Payload map[string]interface{}
}

func (j *JobTypeA) Process() error {
	time.Sleep(3 * time.Second)
	fmt.Println("JobTypeA: ", j.Payload)

	return nil
}

func init() {
	model.RegisterJob("Domain1JobTypeA", func(payload interface{}) model.JobInterface {
		return &JobTypeA{Payload: payload.(map[string]interface{})}
	})
}
