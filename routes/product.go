package routes

import (
	"fmt"
	"log"
	"strconv"

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

func productUpdate(c *fiber.Ctx) error {
	product := new(models.Product)
	if err := c.BodyParser(product); err != nil {
		return err
	}
	_, err := ProductRepository.UpdateProduct(product)
	if err != nil {
		fmt.Println(err)
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "conexão recusada com o servidor",
		})
	}
	return c.JSON(fiber.Map{
		"message": "atualizado com sucesso",
	})
}

func productDelete(c *fiber.Ctx) error {
	id := c.Params("id")
	id_int, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"error": "deve ser numero inteiro",
		})
	}
	_, err = ProductRepository.DeleteProduct(int32(id_int))

	if err != nil {
		fmt.Println(err)
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "conexão recusada com o servidor",
		})
	}
	return c.JSON(fiber.Map{
		"message": "usuario deletado com sucesso",
	})

}

// func productById(c *fiber.Ctx) error {
// 	id := c.Params("id")
// 	res, err := ProductRepository.GetCartByClientId(id)
// 	if err != nil {
// 		log.Println(err)
// 		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"error": "conexão recusada com o servidor",
// 		})
// 	}
// 	data := models.Cart{
// 		Id:          res.Cart.Id,
// 		ClientId:    res.Cart.ClientId,
// 		UpdatedAt:   res.Cart.UpdatedAt,
// 		Status:      res.Cart.Status,
// 		ProductList: res.Cart.ProductListId,
// 	}
// 	return c.JSON(data)
// }
