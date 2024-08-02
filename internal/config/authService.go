package config

import (
	"net"
	"os"

	"github.com/pkg/errors"
)

const (
	authHostEnvName = "AUTH_HOST"
	authPortEnvName = "AUTH_PORT"
)

// AuthServiceConfig config for gRPC server
type AuthServiceConfig interface {
	Address() string
}

type authConfig struct {
	host string
	port string
}

// NewAuthServiceConfig initializes a gRPC configuration.
func NewAuthServiceConfig() (AuthServiceConfig, error) {
	host := os.Getenv(authHostEnvName)
	if len(host) == 0 {
		return nil, errors.New("grpc host not found")
	}

	port := os.Getenv(authPortEnvName)
	if len(port) == 0 {
		return nil, errors.New("grpc port not found")
	}

	return &authConfig{
		host: host,
		port: port,
	}, nil
}

func (cfg *authConfig) Address() string {
	return net.JoinHostPort(cfg.host, cfg.port)
}
