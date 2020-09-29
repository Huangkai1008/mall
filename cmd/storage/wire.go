// +build wireinject

package main

import (
	"github.com/google/wire"

	"github.com/Huangkai1008/micro-kit/pkg/application"

	"github.com/Huangkai1008/mall/internal/app/v1/account/repository"
	storagev1 "github.com/Huangkai1008/mall/internal/app/v1/storage"
	"github.com/Huangkai1008/mall/internal/app/v1/storage/handler"
	"github.com/Huangkai1008/mall/internal/app/v1/storage/router"
	"github.com/Huangkai1008/mall/internal/app/v1/storage/service"
	"github.com/Huangkai1008/mall/internal/pkg/config"
	"github.com/Huangkai1008/mall/internal/pkg/provider/database"
	"github.com/Huangkai1008/mall/internal/pkg/provider/http"
	"github.com/Huangkai1008/mall/internal/pkg/provider/logging"
	"github.com/Huangkai1008/mall/internal/pkg/provider/registry"
	"github.com/Huangkai1008/mall/internal/pkg/provider/storage"
)

var providerSet = wire.NewSet(
	storagev1.ProviderSet,
	config.ProviderSet,
	registry.ProviderSet,
	storage.ProviderSet,
	logging.ProviderSet,
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
