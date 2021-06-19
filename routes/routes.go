package routes

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/cart", createCart)
	app.Post("/login", login)
	app.Post("/user", userCreate)

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("secret"),
	}))
	app.Get("/product", productAll)
	app.Post("/product", productCreate)

	app.Delete("/user/:id", userDelete)
	app.Get("/user", usersAll)
	app.Get("/user/:id", user)
	app.Put("/user/:id", userUpdate)
	
	app.Post("/cart", createCart)
	app.Put("/cart/:id", updateAddOneProduct)
	app.Put("/cart/status", updateStatus)
	app.Get("/cart/client/:id", getCartByClientId)
}
