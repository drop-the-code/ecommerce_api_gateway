package routes

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/vinny1892/ecommerce_api_gateway/models"
	pb "github.com/vinny1892/ecommerce_api_gateway/protos"
	"google.golang.org/grpc"
)

func usersAll(c *fiber.Ctx) error {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "servidor fora do ar",
		})
	}
	defer conn.Close()
	client := pb.NewUserServiceClient(conn)
	req := &pb.Empty{}
	res, err := client.SelectAll(context.Background(), req)
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "conexão recusada com o servidor",
		})
	}
	fmt.Println(res.Users)
	return c.JSON(res.Users)
}
func user(c *fiber.Ctx) error {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "servidor fora do ar",
		})
	}
	defer conn.Close()
	client := pb.NewUserServiceClient(conn)
	id := c.Params("id")
	req := &pb.UserID{
		Id: id,
	}
	res, err := client.SelectById(context.Background(), req)
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
