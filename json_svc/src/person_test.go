package src_test

import (
	"encoding/json"
	"github.com/ahmetcanaydemir/thinksurance-ahmet-can-aydemir/src"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_GetPersonsController(t *testing.T) {
	// Arrange
	src.PersonsDir = "../persons"
	app := fiber.New()
	src.NewPersonController(app)

	// Act
	resp, err := app.Test(httptest.NewRequest("GET", "/persons", nil))
	defer resp.Body.Close()
	var bodyString string
	if resp.StatusCode == http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		bodyString = string(bodyBytes)
	}
	var response []src.Person
	json.Unmarshal([]byte(bodyString), &response)

	// Assert
	utils.AssertEqual(t, nil, err, "Error")
	utils.AssertEqual(t, 50, len(response), "Person Count")
	utils.AssertEqual(t, 200, resp.StatusCode, "Status code")
}
