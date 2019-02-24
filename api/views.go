package api

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/riquellopes/store/camunda"
)

// Buy -
func Buy(client *camunda.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		client.Send("store")
		return c.String(http.StatusOK, "Send to zeebe")
	}
}
