package registry

import (
	"github.com/Huangkai1008/micro-kit/pkg/registry/consul"
	"go.uber.org/zap"

	"github.com/Huangkai1008/mall/internal/pkg/config"
)

// NewConsulClient returns a new consul client.
func NewConsulClient(c *config.Config, logger *zap.Logger) (*consul.Client, error) {
	return consul.NewClient(
		logger,
		consul.WithAddr(c.Consul.Addr),
	)
}

// NewConsulRegistrar returns new Registrar with Consul Client.
func NewConsulRegistrar(client *consul.Client) *consul.Registry {
	return consul.New(client)
}
