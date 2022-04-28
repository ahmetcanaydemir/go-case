package configs

import (
	"github.com/gofiber/fiber/v2"
)

type server struct {
	App    *fiber.App
	Config struct {
		Port                string
		JsonServiceUrl      string
		AlgorithmServiceUrl string
	}
}

var Server server
