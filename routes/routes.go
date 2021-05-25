package routes

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/vinny1892/ecommerce_api_gateway/models"
	pb "github.com/vinny1892/ecommerce_api_gateway/protos"
	"google.golang.org/grpc"

	jwt "github.com/form3tech-oss/jwt-go"
)

func usersAll(c *fiber.Ctx) error {

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "servidor fora do ar",
		})
	}
	defer conn.Close()
	client := pb.NewUserServiceClient(conn)
	req := &pb.Empty{}
	res, err := client.SelectAll(context.Background(), req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "conexão recusada com o servidor",
		})
	}
	fmt.Println(res.Users)
	return c.JSON(res.Users)
}

func login(c *fiber.Ctx) error {
	loginRequest := new(models.LoginRequest)

	if err := c.BodyParser(loginRequest); err != nil {
		return err
	}
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "servidor fora do ar",
		})
	}
	defer conn.Close()
	client := pb.NewUserServiceClient(conn)
	request := &pb.LoginRequest{Email: loginRequest.Email,
		Password: loginRequest.Password}
	res, err := client.SelectByEmail(context.Background(), request)
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
		log.Printf("token.SignedString: %v", err)
		fmt.Println("DEU MERDA")
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
	fmt.Println(seila)
	return c.SendString(seila)
}

func SetupRoutes(app *fiber.App) {
	app.Post("/login", login)
	app.Post("/user", userCreate)
	app.Delete("/user", userDelete)

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("secret"),
	}))

	app.Get("/user", usersAll)
	app.Get("/user/:id", user)
	app.Put("/user", userUpdate)
}
