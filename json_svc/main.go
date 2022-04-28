package main

import (
	"log"

	"github.com/ahmetcanaydemir/thinksurance-ahmet-can-aydemir/src"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	src.NewPersonController(app)

	err := app.Listen(":8080")
	log.Fatal(err)
}
