package config

import (
	"errors"
	"net"
	"os"
	"strconv"
	"time"
)

// PrometheusConfig - config for prometheus
type PrometheusConfig interface {
	Address() string
	ReadTimeout() time.Duration
}

const (
	prometheusHostEnvName        = "PROMETHEUS_HOST"
	prometheusPortEnvName        = "PROMETHEUS_PORT"
	prometheusReadTimeoutEnvName = "PROMETHEUS_READ_TIMEOUT"
)

type prometheusConfig struct {
	host        string
	port        string
	readTimeout time.Duration
}

// NewPrometheusConfig - creates new prometheus config
func NewPrometheusConfig() (PrometheusConfig, error) {
	host := os.Getenv(prometheusHostEnvName)
	if len(host) == 0 {
		return nil, errors.New("prometheus host not found")
	}

	port := os.Getenv(prometheusPortEnvName)
	if len(port) == 0 {
		return nil, errors.New("prometheus port not found")
	}

	timeout, err := strconv.Atoi(os.Getenv(prometheusReadTimeoutEnvName))
	if len(port) == 0 && err != nil {
		return nil, errors.New("read header timeout not found")
	}
	readHeaderTimeout := time.Second * time.Duration(timeout)

	return &prometheusConfig{
		host:        host,
		port:        port,
		readTimeout: readHeaderTimeout,
	}, nil
}

// Address - creates address from host and port
func (cfg *prometheusConfig) Address() string {
	return net.JoinHostPort(cfg.host, cfg.port)
}

// ReadTimeout - returns read timeout
func (cfg *prometheusConfig) ReadTimeout() time.Duration {
	return cfg.readTimeout
}
