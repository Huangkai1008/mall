// +build wireinject

package main

import (
	"github.com/google/wire"

	"mall/internal/app/v1/account"
	"mall/internal/app/v1/account/handler"
	"mall/internal/app/v1/account/repository"
	"mall/internal/app/v1/account/router"
	"mall/internal/app/v1/account/service"
	"mall/internal/pkg/application"
	"mall/internal/pkg/config"
	"mall/internal/pkg/database/gorm"
	"mall/internal/pkg/logging"
	"mall/internal/pkg/registry/consul"
	"mall/internal/pkg/transport/http"
	"mall/internal/pkg/validators"
)

var providerSet = wire.NewSet(
	account.ProviderSet,
	config.ProviderSet,
	logging.ProviderSet,
	http.ProviderSet,
	gorm.ProviderSet,
	router.ProviderSet,
	handler.ProviderSet,
	repository.ProviderSet,
	service.ProviderSet,
	validators.ProviderSet,
	consul.ProviderSet,
)

// CreateApp creates an app by wire.
func CreateApp(cf string) (*application.Application, error) {
	panic(wire.Build(providerSet))
}
