package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/riquellopes/store/camunda"
	"github.com/stretchr/testify/assert"
)

type MockClient struct {
}

func (m *MockClient) Connect() camunda.BaseClient {
	return m
}
func (m *MockClient) Deploy() camunda.BaseClient {
	return m
}
func (m *MockClient) Send(processID string) {
}

func Test_should_get_status_ok(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/buy/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	client := new(MockClient)

	assert.NoError(t, Buy(client)(c))
	assert.Equal(t, http.StatusOK, rec.Code)
}
