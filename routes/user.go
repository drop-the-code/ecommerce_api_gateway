package routes

import (
	"fmt"
	"log"
	"time"

	jwt "github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/vinny1892/ecommerce_api_gateway/models"
	UserRepository "github.com/vinny1892/ecommerce_api_gateway/repositories"
)

func usersAll(c *fiber.Ctx) error {
	data, err := UserRepository.ListAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "conexão recusada com o servidor",
		})
	}
	users := []models.User{}
	for _, item := range data {
		user := models.User{
			Id:      item.Id,
			Name:    item.Name,
			Email:   item.Email,
			Cpf:     item.Cpf,
			Address: item.Address,
			Card: models.Card{
				Id:           item.Card.Id,
				Name:         item.Card.Name,
				SecurityCode: item.Card.SecurityCode,
				ValidThru:    item.Card.ValidThru,
				Number:       item.Card.Number,
			},
			Role:     item.Role.String(),
			Password: item.Password,
		}
		users = append(users, user)
	}

	return c.JSON(users)
}

func login(c *fiber.Ctx) error {
	loginRequest := new(models.LoginRequest)
	if err := c.BodyParser(loginRequest); err != nil {
		return err
	}
	res, err := UserRepository.Login(loginRequest)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "conexão recusada com o servidor",
		})
	}
	if !res.MatchPassword {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = res.User.Id
	claims["email"] = res.User.Email
	claims["name"] = res.User.Name
	claims["role"] = res.User.Role
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		log.Printf("error: %v", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	user := models.User{
		Id:      res.User.Id,
		Name:    res.User.Name,
		Email:   res.User.Email,
		Address: res.User.Address,
		Card: models.Card{
			Id:           res.User.Card.Id,
			Name:         res.User.Card.Name,
			Number:       res.User.Card.Number,
			SecurityCode: res.User.Card.SecurityCode,
			ValidThru:    res.User.Card.ValidThru,
		},
		Password: res.User.Password,
		Role:     res.User.Role.String(),
		Cpf:      res.User.Cpf,
		Token:    t,
	}
	return c.JSON(user)
}
func user(c *fiber.Ctx) error {
	id := c.Params("id")
	res, err := UserRepository.SelectByID(id)
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "conexão recusada com o servidor",
		})
	}
	user := models.User{
		Id:      res.Id,
		Name:    res.Name,
		Email:   res.Email,
		Address: res.Address,
		Card: models.Card{
			Id:           res.Card.Id,
			Name:         res.Card.Name,
			Number:       res.Card.Number,
			SecurityCode: res.Card.SecurityCode,
			ValidThru:    res.Card.ValidThru,
		},
		Password: res.Password,
		Role:     res.Role.String(),
		Cpf:      res.Cpf,
	}
	return c.JSON(user)
}

func userCreate(c *fiber.Ctx) error {
	user := new(models.User)
	errors := models.ValidateStruct(*user)
	if errors != nil {
		return c.JSON(errors)
	}
	if err := c.BodyParser(user); err != nil {
		return err
	}
	res, err := UserRepository.CreateUser(user)
	if err != nil {
		fmt.Println(err)
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "conexão recusada com o servidor",
		})
	}
	return c.JSON(res)

}

func userUpdate(c *fiber.Ctx) error {
	fmt.Println(c.Body())
	return c.SendString("teste")
}
func userDelete(c *fiber.Ctx) error {
	id := c.Params("id")
	res, err := UserRepository.DeleteUser(id)
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "erro ao excluir usuario",
		})
	}
	user := models.User{
		Id:      res.Id,
		Name:    res.Name,
		Email:   res.Email,
		Address: res.Address,
		Card: models.Card{
			Id:           res.Card.Id,
			Name:         res.Card.Name,
			Number:       res.Card.Number,
			SecurityCode: res.Card.SecurityCode,
			ValidThru:    res.Card.ValidThru,
		},
		Password: res.Password,
		Role:     res.Role.String(),
		Cpf:      res.Cpf,
	}
	return c.JSON(user)
}
