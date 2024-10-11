package main

import (
	"event-driven/job_processor/worker"
	_ "event-driven/job_processor/worker/domain1/jobs"
	_ "event-driven/job_processor/worker/domain2/jobs"
	"log"
)

func main() {
	log.Println("Job processor starting...")
	worker.StartWorker()
}
