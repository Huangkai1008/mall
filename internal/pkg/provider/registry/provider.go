package registry

import (
	"github.com/google/wire"

	"github.com/Huangkai1008/micro-kit/pkg/registry"
	"github.com/Huangkai1008/micro-kit/pkg/registry/consul"
)

var ProviderSet = wire.NewSet(NewConsulRegistrar, NewConsulClient, wire.Bind(new(registry.Registrar), new(*consul.Registry)))
