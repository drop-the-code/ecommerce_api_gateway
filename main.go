package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/vinny1892/ecommerce_api_gateway/routes"
)

func main() {

	app := fiber.New()
	app.Use(logger.New())
	routes.SetupRoutes(app)
	log.Fatal(app.Listen(":3000"))
}
