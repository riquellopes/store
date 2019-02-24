package camunda

import (
	"fmt"
	"log"
	"time"

	"github.com/riquellopes/store/camunda/deployer"
	"github.com/zeebe-io/zeebe/clients/go/zbc"
)

// BaseClient -
type BaseClient interface {
	Connect()
	Deploy()
	Send()
}

// Connection -
type Connection interface {
	zbc.ZBClient
}

// Client -
type Client struct {
	Address string
	Port    string
	handler Connection
}

// Publish -
func (c *Client) publish(bpmn deployer.Bpmn) error {
	response, err := c.handler.NewDeployWorkflowCommand().AddResourceFile(bpmn.Path).Send()

	log.Println(response.String())
	return err
}

// Connect -
func (c *Client) Connect() *Client {
	log.Println("Zeebe connection was started.")

	// @TODO For solve a crazy docker problem. I'll search for solution after.
	time.Sleep(10 * time.Second)

	client, err := zbc.NewZBClient(fmt.Sprintf("%s:%s", c.Address, c.Port))
	c.handler = client

	if err != nil {
		panic(err)
	}

	log.Println("Zeebe connection is ready.")
	return c
}

// Deploy -
func (c *Client) Deploy() *Client {
	log.Println("Zeebe deploy was started.")
	workflows := deployer.New()

	for _, bpmn := range workflows.Itens() {
		err := c.publish(bpmn)

		if err != nil {
			panic(err)
		}

		log.Printf("Workflow -> %s publish.", bpmn.Path)
	}

	log.Println("Zeebe deploy is ready.")
	return c
}

// Send -
func (c *Client) Send(processID string) {
	log.Printf("Starting the processID '%s'.", processID)

	// @TODO define a payload
	payload := make(map[string]interface{})

	request, err := c.handler.NewCreateInstanceCommand().BPMNProcessId(processID).LatestVersion().PayloadFromMap(payload)

	if err != nil {
		log.Panic(err)
	}

	result, err := request.Send()
	if err != nil {
		log.Panic(err)
	}

	log.Println(result.String())
}
