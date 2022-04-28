package controller

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/ahmetcanaydemir/thinksurance-ahmet-can-aydemir/pkg/configs"
	"github.com/gofiber/fiber/v2"
)

type GatewayController struct {
}

func NewGatewayController() {
	ct := &GatewayController{}
	configs.Server.App.Get("/persons", ct.GetPersons)
	configs.Server.App.Post("/find-position", ct.PostFindPosition)
}

func (ct GatewayController) GetPersons(c *fiber.Ctx) error {
	resp, err := http.Get(configs.Server.Config.JsonServiceUrl + "/persons")
	if err != nil {
		log.Panicln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panicln(err)
	}

	c.Status(resp.StatusCode)
	return c.Send(body)
}

type FindPositionRequest struct {
	Array  []int `json:"array" xml:"array" form:"array"`
	Search int   `json:"search" xml:"search" form:"search"`
}

func (ct GatewayController) PostFindPosition(c *fiber.Ctx) error {
	resp, err := http.Post(configs.Server.Config.AlgorithmServiceUrl+"/find-position", fiber.MIMEApplicationJSON, bytes.NewBuffer(c.Body()))
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	c.Status(resp.StatusCode)
	return c.Send(body)
}
