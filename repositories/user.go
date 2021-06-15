package repositories

import (
	"context"

	"github.com/vinny1892/ecommerce_api_gateway/config"
	"github.com/vinny1892/ecommerce_api_gateway/models"
	pb "github.com/vinny1892/ecommerce_api_gateway/protos/user"

	"github.com/vinny1892/ecommerce_api_gateway/clientGRPC"
)

var uri_connection_microservice_user string = config.Environment().GetString("IP_ECOMMERCE_USER") + ":" + config.Environment().GetString("GRPC_PORT_ECOMMERCE_USER")

func ListAll() ([]*pb.User, error) {
	conn, err := clientGRPC.Connect(uri_connection_microservice_user)
	if err != nil {
		return nil, err
	}

	defer clientGRPC.Close_Connect(conn)
	client := pb.NewUserServiceClient(conn)
	req := &pb.Empty{}
	res, err := client.SelectAll(context.Background(), req)
	if err != nil {
		return nil, err
	}
	return res.Users, err
}

func Login(loginRequest *models.LoginRequest) (*pb.LoginResponse, error) {
	conn, err := clientGRPC.Connect(uri_connection_microservice_user)
	if err != nil {
		return nil, err
	}
	defer clientGRPC.Close_Connect(conn)
	client := pb.NewUserServiceClient(conn)
	request := &pb.LoginRequest{Email: loginRequest.Email, Password: loginRequest.Password}
	res, err := client.SelectByEmail(context.Background(), request)
	if err != nil {
		return nil, err
	}
	return res, err
}

func SelectByID(id string) (*pb.User, error) {
	conn, err := clientGRPC.Connect(uri_connection_microservice_user)
	if err != nil {
		return nil, err
	}
	defer clientGRPC.Close_Connect(conn)
	client := pb.NewUserServiceClient(conn)
	request := &pb.UserID{Id: id}
	res, err := client.SelectById(context.Background(), request)
	if err != nil {
		return nil, err
	}
	return res, err
}

func DeleteUser(id string) (*pb.User, error) {
	conn, err := clientGRPC.Connect(uri_connection_microservice_user)
	if err != nil {
		return nil, err
	}
	defer clientGRPC.Close_Connect(conn)
	client := pb.NewUserServiceClient(conn)
	request := &pb.UserID{Id: id}
	res, err := client.Delete(context.Background(), request)
	if err != nil {
		return nil, err
	}
	return res, err
}

func CreateUser(userRequest *models.User) (*pb.User, error) {
	conn, err := clientGRPC.Connect(uri_connection_microservice_user)
	if err != nil {
		return nil, err
	}
	defer clientGRPC.Close_Connect(conn)
	client := pb.NewUserServiceClient(conn)
	var role pb.Role = pb.Role_funcionario /* funcionario */
	if userRequest.Role == "cliente" {
		role = pb.Role_cliente /* cliente */
	}
	request := &pb.UserRequest{
		User: &pb.User{
			Name:     userRequest.Name,
			Email:    userRequest.Email,
			Cpf:      userRequest.Cpf,
			Address:  userRequest.Address,
			Role:     role,
			Password: userRequest.Password,
			Card: &pb.Card{
				Name:         userRequest.Card.Name,
				SecurityCode: userRequest.Card.SecurityCode,
				ValidThru:    userRequest.Card.ValidThru,
				Number:       userRequest.Card.Number,
			},
		},
	}
	res, err := client.Create(context.Background(), request)
	if err != nil {
		return nil, err
	}
	return res, err
}
