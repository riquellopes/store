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

// Client -
type Client struct {
	Address string
	Port    string
	handler zbc.ZBClient
	logger  log.Logger
}

// Publish -
func (c *Client) publish(bpmn deployer.Bpmn) error {
	response, err := c.handler.NewDeployWorkflowCommand().AddResourceFile(bpmn.Path).Send()

	fmt.Println(response.String())
	return err
}

// Connect -
func (c *Client) Connect() *Client {
	fmt.Println("Zeebe connection was started.")

	// @TODO For solve a crazy docker problem. I'll search for solution after.
	time.Sleep(10 * time.Second)

	client, err := zbc.NewZBClient(fmt.Sprintf("%s:%s", c.Address, c.Port))
	c.handler = client

	if err != nil {
		panic(err)
	}

	fmt.Println("Zeebe connection is ready.")
	return c
}

// Deploy -
func (c *Client) Deploy() *Client {
	fmt.Println("Zeebe deploy was started.")
	workflows := deployer.New()

	for _, bpmn := range workflows.Itens() {
		err := c.publish(bpmn)

		if err != nil {
			panic(err)
		}

		fmt.Println()
		fmt.Printf("Workflow -> %s publish.", bpmn.Path)
		fmt.Println()
	}

	fmt.Println("Zeebe deploy is ready.")
	return c
}

// Send -
// func (z *Client) Send(task string) error {
// 	return nil
// }
