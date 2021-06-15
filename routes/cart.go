package routes

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/vinny1892/ecommerce_api_gateway/models"
	CartRepository "github.com/vinny1892/ecommerce_api_gateway/repositories"
)

func createCart(c *fiber.Ctx) error {
	cart := new(models.Cart)

	if err := c.BodyParser(cart); err != nil {
		return err
	}

	res, err := CartRepository.CreateCart(cart)

	if err != nil {
		log.Println(err)
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "conexão recusada com o servidor",
		})
	}
	data := models.Cart{
		Id:          res.Cart.Id,
		ClientId:    res.Cart.ClientId,
		UpdatedAt:   res.Cart.UpdatedAt,
		Status:      res.Cart.Status,
		ProductList: res.Cart.ProductListId,
	}

	return c.JSON(data)
}

func updateAddOneProduct(c *fiber.Ctx) error {
	cart := new(models.Cart)
	if err := c.BodyParser(cart); err != nil {
		return err
	}
	res, err := CartRepository.UpdateAddOneProduct(cart)
	if err != nil {
		log.Println(err)
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "conexão recusada com o servidor",
		})
	}
	data := models.Cart{
		Id:          res.Cart.Id,
		ClientId:    res.Cart.ClientId,
		UpdatedAt:   res.Cart.UpdatedAt,
		Status:      res.Cart.Status,
		ProductList: res.Cart.ProductListId,
	}

	return c.JSON(data)

}

func getCartByClientId(c *fiber.Ctx) error {
	id := c.Params("id")
	res, err := CartRepository.GetCartByClientId(id)
	if err != nil {
		log.Println(err)
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "conexão recusada com o servidor",
		})
	}
	data := models.Cart{
		Id:          res.Cart.Id,
		ClientId:    res.Cart.ClientId,
		UpdatedAt:   res.Cart.UpdatedAt,
		Status:      res.Cart.Status,
		ProductList: res.Cart.ProductListId,
	}
	return c.JSON(data)
}

func updateStatus(c *fiber.Ctx) error {
	cart := new(models.Cart)
	if err := c.BodyParser(cart); err != nil {
		return err
	}
	log.Println("aqui")
	fmt.Println(cart.Id)
	res, err := CartRepository.UpdateStatus(cart)
	if err != nil {
		log.Println(err)
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "conexão recusada com o servidor",
		})
	}
	data := models.Cart{
		Id:          res.Cart.Id,
		ClientId:    res.Cart.ClientId,
		UpdatedAt:   res.Cart.UpdatedAt,
		Status:      res.Cart.Status,
		ProductList: res.Cart.ProductListId,
	}
	return c.JSON(data)

}
