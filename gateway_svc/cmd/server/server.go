package cmd

import (
	"fmt"
	"os"

	"github.com/ahmetcanaydemir/thinksurance-ahmet-can-aydemir/db"
	"github.com/ahmetcanaydemir/thinksurance-ahmet-can-aydemir/pkg/api/controller"
	"github.com/ahmetcanaydemir/thinksurance-ahmet-can-aydemir/pkg/api/middleware"
	"github.com/ahmetcanaydemir/thinksurance-ahmet-can-aydemir/pkg/configs"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go run main.go",
	Short: "Go Rest Clean Boilerplate Service",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize()
	configs.Server.Config.JsonServiceUrl = os.Getenv("JSON_SERVICE_URL")
	configs.Server.Config.AlgorithmServiceUrl = os.Getenv("ALGORITHM_SERVICE_URL")
	rootCmd.Flags().StringVarP(&configs.Server.Config.Port, "port", "p", "8080", "api port")

	rootCmd.RunE = start
}

func start(_ *cobra.Command, _ []string) error {
	// region Fiber ####
	app := fiber.New(fiber.Config{
		ErrorHandler: middleware.ErrorHandler,
	})
	configs.Server.App = app

	// Middlewares
	app.Use(cors.New())
	app.Use(recover.New())
	app.Use(basicauth.New(basicauth.Config{
		Users: db.Users,
		Unauthorized: func(c *fiber.Ctx) error {
			c.Status(fiber.StatusForbidden)
			return c.SendString("You need BasicAuth authentication to access this API.")
		},
	}))
	// endregion

	// region Controllers ####
	controller.NewGatewayController()
	// endregion

	configs.Server.App.Use(middleware.NotFound)
	err := configs.Server.App.Listen(":" + configs.Server.Config.Port)

	return err
}
