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

// @title Zamazon API
// @version 1.0
// @description Doc api zamazon
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email kaiomudkt@gmail.compatible
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host https://api.zamazon.gorillas.ninja
// @BasePath /
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
