// +build wireinject

package main

import (
	"github.com/google/wire"

	"mall/internal/app/v1/storage"
	"mall/internal/app/v1/storage/handler"
	"mall/internal/app/v1/storage/router"
	"mall/internal/app/v1/storage/service"
	"mall/internal/pkg/application"
	"mall/internal/pkg/config"
	"mall/internal/pkg/database/gorm"
	"mall/internal/pkg/logging"
	"mall/internal/pkg/registry/consul"
	minioCli "mall/internal/pkg/storage/minio"
	"mall/internal/pkg/transport/http"
	"mall/internal/pkg/validators"
)

var providerSet = wire.NewSet(
	storage.ProviderSet,
	config.ProviderSet,
	logging.ProviderSet,
	minioCli.ProviderSet,
	http.ProviderSet,
	gorm.ProviderSet,
	router.ProviderSet,
	handler.ProviderSet,
	service.ProviderSet,
	validators.ProviderSet,
	consul.ProviderSet,
)

// CreateApp creates an app by wire.
func CreateApp(cf string) (*application.Application, error) {
	panic(wire.Build(providerSet))
}
