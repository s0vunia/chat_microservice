package authservice

import (
	"context"

	desc "github.com/s0vunia/auth_microservice_proto/access_v1/gen"
)

// Check реализует метод интерфейса AuthService
func (c *Client) Check(ctx context.Context, endpointAddress string) error {
	client := desc.NewAccessV1Client(c.conn)
	_, err := client.Check(ctx, &desc.CheckRequest{EndpointAddress: endpointAddress})
	if err != nil {
		return err
	}
	return nil
}
