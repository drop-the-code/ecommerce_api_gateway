package clientGRPC

import (
	"google.golang.org/grpc"
)

func Connect(uri_connection string) (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(uri_connection, grpc.WithInsecure())
	if err != nil {
		return conn, err
	}
	return conn, err
}

func Close_Connect(conn *grpc.ClientConn) {
	conn.Close()
}
