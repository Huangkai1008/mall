// +build wireinject

package main

import (
	"github.com/google/wire"
	"mall/internal/app/v1/user/handler"
	"mall/internal/app/v1/user/repository"
	"mall/internal/app/v1/user/router"
	"mall/internal/app/v1/user/service"
	"mall/internal/pkg/database/gorm"
	"mall/internal/pkg/storage/minio"

	"mall/internal/app/v1/user"
	"mall/internal/pkg/application"
	"mall/internal/pkg/config"
	"mall/internal/pkg/logging"
	"mall/internal/pkg/transport/http"
)

var providerSet = wire.NewSet(
	user.ProviderSet,
	config.ProviderSet,
	logging.ProviderSet,
	http.ProviderSet,
	minio.ProviderSet,
	gorm.ProviderSet,
	router.ProviderSet,
	handler.ProviderSet,
	repository.ProviderSet,
	service.ProviderSet,
)

// CreateApp creates an app by wire.
func CreateApp(cf string) (*application.Application, error) {
	panic(wire.Build(providerSet))
}
