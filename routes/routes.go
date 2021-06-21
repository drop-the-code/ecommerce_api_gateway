package routes

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
)

func SetupRoutes(app *fiber.App) {

	v1 := app.Group("/v1")
	v1.Post("/login", login)
	v1.Post("/user", userCreate)

	app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
		URL:         "/public/swagger.json",
		DeepLinking: false,
	}))
	app.Static("/public", "./docs")
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("secret"),
	}))
	v1.Get("/logout", logout)
	v1.Get("/product", productAll)
	v1.Get("/product/:id", productById)
	v1.Post("/product", productCreate)
	v1.Put("/product/:id", productUpdate)
	v1.Delete("/product/:id", productDelete)
	v1.Delete("/user/:id", userDelete)
	v1.Get("/user", usersAll)
	v1.Get("/user/:id", user)
	v1.Put("/user/:id", userUpdate)
	v1.Post("/cart", updateAddOneProduct)
	v1.Put("/cart/status", updateStatus)
	v1.Get("/cart/client/:id", getCartByClientId)
}
