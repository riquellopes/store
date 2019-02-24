package main

import (
	"fmt"
	"os"

	"github.com/labstack/echo"
	"github.com/riquellopes/store/api"
	"github.com/riquellopes/store/camunda"
	"github.com/riquellopes/store/camunda/tasks"
)

func main() {
	zee := camunda.Client{
		Address: os.Getenv("CLIENT_ADDRESS"),
		Port:    os.Getenv("CLIENT_PORT"),
	}

	zee.Connect().Deploy()

	go func() {
		w := camunda.Worker{
			Client: zee,
		}
		w.Add("check-availability", tasks.CheckAvailability).
			Add("finish-order", tasks.FinishOrder).
			Add("send-message", tasks.SendMessage)
	}()

	// Register service
	app := echo.New()
	app.GET("/buy/", api.Buy(&zee))

	port := os.Getenv("PORT")
	app.Start(fmt.Sprintf(":%s", port))
}
