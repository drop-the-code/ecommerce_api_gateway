package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/vinny1892/ecommerce_api_gateway/config"
	"github.com/vinny1892/ecommerce_api_gateway/routes"
)

func main() {
	v1 := config.Environment()
	fmt.Println(v1.GetInt("PORT"))
	app := fiber.New()
	app.Use(cors.New())

	app.Use(logger.New())
	routes.SetupRoutes(app)
	var port_string string = ":" + v1.GetString("PORT")
	log.Fatal(app.Listen(port_string))
}
