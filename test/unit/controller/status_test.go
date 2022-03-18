package test_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ahmetcanaydemir/go-rest-clean-boilerplate/pkg/api/controller"
	"github.com/ahmetcanaydemir/go-rest-clean-boilerplate/pkg/configs"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func Test_StatusEndpoint(t *testing.T) {
	// Arrange
	expectedResponse := controller.Status{
		Ping: controller.Ping{
			Error:     nil,
			IsHealthy: true,
			Message:   nil,
		},
	}
	app := fiber.New()
	configs.Server.App = app
	controller.NewStatusController()

	// Act
	resp, err := app.Test(httptest.NewRequest("GET", "/status", nil))
	defer resp.Body.Close()

	var bodyString string
	if resp.StatusCode == http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		bodyString = string(bodyBytes)
	}
	var response controller.Status
	json.Unmarshal([]byte(bodyString), &response)

	// Assert
	utils.AssertEqual(t, expectedResponse, response, "Body")
	utils.AssertEqual(t, nil, err, "Error")
	utils.AssertEqual(t, 200, resp.StatusCode, "Status code")
}
