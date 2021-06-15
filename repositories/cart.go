package repositories

import (
	"context"
	"fmt"

	"github.com/vinny1892/ecommerce_api_gateway/clientGRPC"
	"github.com/vinny1892/ecommerce_api_gateway/config"
	"github.com/vinny1892/ecommerce_api_gateway/models"
	pb "github.com/vinny1892/ecommerce_api_gateway/protos/cart"
)

var uri_connection_microservice_cart string = config.Environment().GetString("IP_ECOMMERCE_CART") + ":" + config.Environment().GetString("GRPC_PORT_ECOMMERCE_CART")

func CreateCart(cart *models.Cart) (*pb.CartResponse, error) {
	conn, err := clientGRPC.Connect(uri_connection_microservice_cart)
	if err != nil {
		return nil, err
	}
	defer clientGRPC.Close_Connect(conn)
	client := pb.NewCartServiceClient(conn)
	request := &pb.CartRequest{
		Cart: &pb.Cart{
			ClientId:      cart.ClientId,
			Status:        cart.Status,
			ProductListId: cart.ProductList,
		},
	}
	res, err := client.CreateCart(context.Background(), request)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func GetCart(cartID string) (*pb.CartResponse, error) {
	conn, err := clientGRPC.Connect(uri_connection_microservice_cart)
	if err != nil {
		return nil, err
	}
	defer clientGRPC.Close_Connect(conn)
	client := pb.NewCartServiceClient(conn)
	request := &pb.CartIdRequest{
		Id: cartID,
	}
	res, err := client.GetCart(context.Background(), request)
	if err != nil {
		return nil, err
	}
	return res, nil

}

func UpdateAddOneProduct(cartRequest *models.Cart) (*pb.CartResponse, error) {

	conn, err := clientGRPC.Connect(uri_connection_microservice_cart)
	if err != nil {
		return nil, err
	}
	defer clientGRPC.Close_Connect(conn)
	client := pb.NewCartServiceClient(conn)
	request := &pb.CartRequest{
		Cart: &pb.Cart{
			Id:            cartRequest.Id,
			UpdatedAt:     cartRequest.UpdatedAt,
			ClientId:      cartRequest.ClientId,
			Status:        cartRequest.Status,
			ProductListId: cartRequest.ProductList,
		},
	}
	res, err := client.UpdateAddOneProduct(context.Background(), request)
	if err != nil {
		return nil, err
	}
	return res, nil

}

func GetCartByClientId(clientID string) (*pb.CartResponse, error) {

	conn, err := clientGRPC.Connect(uri_connection_microservice_cart)
	if err != nil {
		return nil, err
	}
	defer clientGRPC.Close_Connect(conn)
	client := pb.NewCartServiceClient(conn)
	request := &pb.ClientIdRequest{
		ClientId: clientID,
	}
	res, err := client.GetCartByClientId(context.Background(), request)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func UpdateStatus(cartRequest *models.Cart) (*pb.CartResponse, error) {
	conn, err := clientGRPC.Connect(uri_connection_microservice_cart)
	fmt.Println(cartRequest.Id)
	if err != nil {
		return nil, err
	}
	defer clientGRPC.Close_Connect(conn)
	client := pb.NewCartServiceClient(conn)
	request := &pb.CartRequest{
		Cart: &pb.Cart{
			Id:            cartRequest.Id,
			UpdatedAt:     cartRequest.UpdatedAt,
			ClientId:      cartRequest.ClientId,
			Status:        cartRequest.Status,
			ProductListId: cartRequest.ProductList,
		},
	}
	res, err := client.UpdateStatus(context.Background(), request)
	if err != nil {
		return nil, err
	}
	return res, nil
}
