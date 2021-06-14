package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/vinny1892/ecommerce_api_gateway/models"
	ProductRepository "github.com/vinny1892/ecommerce_api_gateway/repositories"
)

func productAll(c *fiber.Ctx) error {
	data, err := ProductRepository.ListAllProduct()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "conexão recusada com o servidor",
		})
	}
	// users := []models.User{}
	// for _, item := range data {
	// 	user := models.User{
	// 		Id:      item.Id,
	// 		Name:    item.Name,
	// 		Email:   item.Email,
	// 		Cpf:     item.Cpf,
	// 		Address: item.Address,
	// 		Card: models.Card{
	// 			Id:           item.Card.Id,
	// 			Name:         item.Card.Name,
	// 			SecurityCode: item.Card.SecurityCode,
	// 			ValidThru:    item.Card.ValidThru,
	// 			Number:       item.Card.Number,
	// 		},
	// 		Role:     item.Role.String(),
	// 		Password: item.Password,
	// 	}
	// 	users = append(users, user)
	// }
	fmt.Println(data)
	return c.JSON(data)

}

func productCreate(c *fiber.Ctx) error {
	product := new(models.Product)

	if err := c.BodyParser(product); err != nil {
		return err
	}

	errors := models.ValidateStructProduct(*product)
	if errors != nil {
		return c.JSON(errors)
	}

	data, err := ProductRepository.CreateProduct(product)
	if err != nil {
		fmt.Println(err)
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "conexão recusada com o servidor",
		})
	}
	return c.JSON(data)

}
