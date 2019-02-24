package api

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/riquellopes/store/camunda"
)

// Buy -
func Buy(client camunda.BaseClient) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Sending information for camunda zeebe.
		client.Send("store")
		return c.String(http.StatusOK, "Send to zeebe")
	}
}
