// +build wireinject

package main

import (
	"github.com/google/wire"

	"github.com/Huangkai1008/mall/internal/app/v1/account"
	"github.com/Huangkai1008/mall/internal/app/v1/account/handler"
	"github.com/Huangkai1008/mall/internal/app/v1/account/repository"
	"github.com/Huangkai1008/mall/internal/app/v1/account/router"
	"github.com/Huangkai1008/mall/internal/app/v1/account/service"
	"github.com/Huangkai1008/mall/internal/pkg/config"
	"github.com/Huangkai1008/mall/internal/pkg/provider/auth"
	"github.com/Huangkai1008/mall/internal/pkg/provider/database"
	"github.com/Huangkai1008/mall/internal/pkg/provider/http"
	"github.com/Huangkai1008/mall/internal/pkg/provider/logging"
	"github.com/Huangkai1008/mall/internal/pkg/provider/registry"

	"github.com/Huangkai1008/micro-kit/pkg/application"
)

var providerSet = wire.NewSet(
	account.ProviderSet,
	config.ProviderSet,
	registry.ProviderSet,
	logging.ProviderSet,
	auth.ProviderSet,
	router.ProviderSet,
	http.ProviderSet,
	database.ProviderSet,
	repository.ProviderSet,
	service.ProviderSet,
	handler.ProviderSet,
)

// CreateApp creates an app by wire.
func CreateApp(cf string) (*application.Application, error) {
	panic(wire.Build(providerSet))
}
