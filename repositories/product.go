package repositories

import (
	"context"

	"github.com/vinny1892/ecommerce_api_gateway/clientGRPC"
	"github.com/vinny1892/ecommerce_api_gateway/config"
	"github.com/vinny1892/ecommerce_api_gateway/models"
	pb "github.com/vinny1892/ecommerce_api_gateway/protos/product"
)

var uri_connection_microservice_product string = config.Environment().GetString("IP_ECOMMERCE_PRODUCT") + ":" + config.Environment().GetString("GRPC_PORT_ECOMMERCE_PRODUCT")

func ListAllProduct() ([]*pb.Product, error) {
	conn, err := clientGRPC.Connect(uri_connection_microservice_product)
	if err != nil {
		return nil, err
	}

	defer clientGRPC.Close_Connect(conn)
	client := pb.NewProductCRUDClient(conn)
	req := &pb.Empty{}
	res, err := client.SelectAll(context.Background(), req)
	if err != nil {
		return nil, err
	}
	return res.Items, err
}

func SelectByIDProduct(id int32) (*pb.Product, error) {
	conn, err := clientGRPC.Connect(uri_connection_microservice_product)
	if err != nil {
		return nil, err
	}
	defer clientGRPC.Close_Connect(conn)
	client := pb.NewProductCRUDClient(conn)

	request := &pb.ProductID{
		ProductID: id,
	}
	res, err := client.SelectByID(context.Background(), request)
	if err != nil {
		return nil, err
	}
	return res, err
}

func DeleteProduct(id int32) (*pb.Empty, error) {
	conn, err := clientGRPC.Connect(uri_connection_microservice_product)
	if err != nil {
		return nil, err
	}
	defer clientGRPC.Close_Connect(conn)
	client := pb.NewProductCRUDClient(conn)
	request := &pb.ProductID{
		ProductID: id,
	}
	res, err := client.Delete(context.Background(), request)
	if err != nil {
		return nil, err
	}
	return res, err
}

func CreateProduct(productRequest *models.Product) (*pb.Empty, error) {
	conn, err := clientGRPC.Connect(uri_connection_microservice_product)
	if err != nil {
		return nil, err
	}
	defer clientGRPC.Close_Connect(conn)
	client := pb.NewProductCRUDClient(conn)
	request := &pb.Product{
		ProductName:  productRequest.Name,
		ProductPrice: productRequest.Price,
		ProviderCNPJ: productRequest.ProviderCnpj,
		Description:  productRequest.Description,
	}
	res, err := client.Insert(context.Background(), request)
	if err != nil {
		return nil, err
	}
	return res, err
}

func UpdateProduct(productRequest *models.Product) (*pb.Empty, error) {
	conn, err := clientGRPC.Connect(uri_connection_microservice_product)
	if err != nil {
		return nil, err
	}
	defer clientGRPC.Close_Connect(conn)
	request := &pb.Product{
		ProductID:    productRequest.Id,
		ProductName:  productRequest.Name,
		ProviderCNPJ: productRequest.ProviderCnpj,
		ProductPrice: productRequest.Price,
		Description:  productRequest.Description,
	}
	client := pb.NewProductCRUDClient(conn)

	res, err := client.Update(context.Background(), request)
	if err != nil {
		return nil, err
	}
	return res, nil

}
