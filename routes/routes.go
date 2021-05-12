package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/vinny1892/ecommerce_api_gateway/models"
)

func usersAll(c *fiber.Ctx) error {
	users := []models.User{
		{
			Name:     "Vinicius",
			Email:    "vinny1892@gmail.com",
			Cpf:      "12341234",
			Endereco: "Seila",
			Cartao:   "awefiawoefç",
			Role:     "funcionario",
		},
	}
	seila := models.User{}
	users = append(users, seila)
	return c.JSON(users)
}
func user(c *fiber.Ctx) error {
	fmt.Println(c.Query("seila"))
	user := models.User{
		Name:     "Vinicius",
		Email:    "vinny1892@gmail.com",
		Cpf:      "12341234",
		Endereco: "Seila",
		Cartao:   "awefiawoefç",
		Role:     "funcionario",
	}
	return c.JSON(user)
}

func userCreate(c *fiber.Ctx) error {
	fmt.Println(string(c.Body()))
	return c.SendString(string(c.Body()))
}
func userUpdate(c *fiber.Ctx) error {
	fmt.Println(c.Body())
	return c.SendString("teste")
}
func userDelete(c *fiber.Ctx) error {
	seila := string(c.Body())
	return c.SendString(seila)
}

func SetupRoutes(app *fiber.App) {
	app.Get("/user", usersAll)
	app.Get("/user/:id", user)
	app.Post("/user", userCreate)
	app.Put("/user", userUpdate)
	app.Delete("/user", userDelete)
}
