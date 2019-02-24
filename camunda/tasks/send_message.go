package tasks

import (
	"log"

	"github.com/zeebe-io/zeebe/clients/go/entities"
	"github.com/zeebe-io/zeebe/clients/go/worker"
)

// SendMessage -
func SendMessage(client worker.JobClient, job entities.Job) {
	jobKey := job.GetKey()

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

	log.Println("Complete SendMessage", jobKey, "of type", job.Type)
	request.Send()
}
