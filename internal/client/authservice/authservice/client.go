package authservice

import (
	"github.com/s0vunia/chat_microservice/internal/client/authservice"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Client структура клиента, реализующая интерфейс AuthService
type Client struct {
	conn *grpc.ClientConn
}

// NewClient создает нового клиента, подключаясь к gRPC серверу
func NewClient(address string) (authservice.AuthService, error) {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return &Client{conn: conn}, nil
}

// Close закрывает соединение с gRPC сервером
func (c *Client) Close() {
	_ = c.conn.Close()
}
