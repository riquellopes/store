package tasks

import (
	"log"

	"github.com/zeebe-io/zeebe/clients/go/entities"
	"github.com/zeebe-io/zeebe/clients/go/worker"
)

// FinishOrder -
func FinishOrder(client worker.JobClient, job entities.Job) {
	jobKey := job.GetKey()

	headers, err := job.GetCustomHeadersAsMap()
	if err != nil {
		// failed to handle job as we require the custom job headers
		failJob(client, job)
		return
	}

	payload, err := job.GetPayloadAsMap()
	if err != nil {
		// failed to handle job as we require the payload
		failJob(client, job)
		return
	}

	request, err := client.NewCompleteJobCommand().JobKey(jobKey).PayloadFromMap(payload)
	if err != nil {
		// failed to set the updated payload
		failJob(client, job)
		return
	}

	log.Println("Complete FinishOrder", jobKey, "of type", job.Type)
	log.Println("Processing order:", payload["id"])
	log.Println("Collect money using payment method:", headers["method"])

	request.Send()
}
