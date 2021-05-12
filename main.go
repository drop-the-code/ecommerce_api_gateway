package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vinny1892/ecommerce_api_gateway/routes"
)

func main() {
	app := fiber.New()
	routes.SetupRoutes(app)
	app.Listen(":3000")
}
