package tasks

import (
	"log"
	"math/rand"

	"github.com/zeebe-io/zeebe/clients/go/entities"
	"github.com/zeebe-io/zeebe/clients/go/worker"
)

//CheckAvailability -
func CheckAvailability(client worker.JobClient, job entities.Job) {
	jobKey := job.GetKey()

	payload, err := job.GetPayloadAsMap()
	if err != nil {
		// failed to handle job as we require the payload
		failJob(client, job)
		return
	}

	payload["available"] = rand.Int()%2 == 0
	request, err := client.NewCompleteJobCommand().JobKey(jobKey).PayloadFromMap(payload)
	if err != nil {
		// failed to set the updated payload
		failJob(client, job)
		return
	}

	log.Println("Complete job", jobKey, "of type", job.Type)
	request.Send()
}

func failJob(client worker.JobClient, job entities.Job) {
	log.Println("Failed to complete job", job.GetKey())
	client.NewFailJobCommand().JobKey(job.GetKey()).Retries(job.Retries - 1).Send()
}
