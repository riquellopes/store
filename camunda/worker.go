package camunda

import (
	"log"

	"github.com/zeebe-io/zeebe/clients/go/worker"
)

// Worker -
type Worker struct {
	Client Client
}

// Add -
func (w *Worker) Add(name string, jobFn worker.JobHandler) *Worker {
	/*
		w := Worker{}
		w.add("task-pagamento", func (client worker.JobClient, job entities.Job)  {
		    do something
		})
	*/
	log.Printf("Task register -> '%s'.", name)
	w.Client.handler.NewJobWorker().JobType(name).Handler(jobFn).Open()
	return w
}
