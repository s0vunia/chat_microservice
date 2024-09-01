package authservice

import (
	"context"
)

// AuthService interface
type AuthService interface {
	IsUserExists(ctx context.Context, userIDs []int64) (bool, error)
	Check(ctx context.Context, endpointAddress string) error
}
