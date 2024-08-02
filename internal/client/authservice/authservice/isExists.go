package authservice

import (
	"context"

	desc "github.com/s0vunia/auth_microservice_proto/auth_v1/gen"
)

// IsUserExists реализует метод интерфейса AuthService
func (c *Client) IsUserExists(ctx context.Context, userIDs []int64) (bool, error) {
	client := desc.NewAuthV1Client(c.conn)
	resp, err := client.IsExists(ctx, &desc.IsExistsRequest{Ids: userIDs})
	if err != nil {
		return false, err
	}
	return resp.Exists, nil
}
