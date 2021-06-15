package routes

import (
	"fmt"
	"log"

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

	products := []models.Product{}
	for _, item := range data {
		product := models.Product{
			Id:           item.ProductID,
			Name:         item.ProductName,
			Description:  item.Description,
			Price:        item.ProductPrice,
			ProviderCnpj: item.ProviderCNPJ,
		}
		products = append(products, product)
	}
	log.Println("Listando dados ...")
	return c.JSON(products)

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
