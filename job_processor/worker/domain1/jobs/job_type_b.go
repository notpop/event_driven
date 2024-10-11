package jobs

import (
	"event-driven/common/model"
	"fmt"
)

type JobTypeB struct {
	Payload map[string]interface{}
}

func (j *JobTypeB) Process() error {
	fmt.Println("JobTypeB: ", j.Payload)

	return nil
}

func init() {
	model.RegisterJob("Domain1JobTypeB", func(payload interface{}) model.JobInterface {
		return &JobTypeB{Payload: payload.(map[string]interface{})}
	})
}
