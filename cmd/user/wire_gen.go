// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/google/wire"
	"mall/internal/app/v1/user"
	"mall/internal/app/v1/user/handler"
	"mall/internal/app/v1/user/repository"
	"mall/internal/app/v1/user/router"
	"mall/internal/app/v1/user/service"
	"mall/internal/pkg/application"
	"mall/internal/pkg/config"
	"mall/internal/pkg/database/gorm"
	"mall/internal/pkg/logging"
	"mall/internal/pkg/storage/minio"
	"mall/internal/pkg/transport/http"
)

// Injectors from wire.go:

// CreateApp creates an app by wire.
func CreateApp(cf string) (*application.Application, error) {
	viper, err := config.New(cf)
	if err != nil {
		return nil, err
	}
	options, err := logging.NewOptions(viper)
	if err != nil {
		return nil, err
	}
	logger, err := logging.New(options)
	if err != nil {
		return nil, err
	}
	userOptions, err := user.NewOptions(viper, logger)
	if err != nil {
		return nil, err
	}
	httpOptions, err := http.NewOptions(viper)
	if err != nil {
		return nil, err
	}
	gormOptions, err := gorm.NewOptions(viper)
	if err != nil {
		return nil, err
	}
	db, err := gorm.New(gormOptions)
	if err != nil {
		return nil, err
	}
	repositoryRepository := repository.NewRepository(logger, db)
	serviceService := service.NewService(logger, repositoryRepository)
	handlerHandler := handler.NewHandler(logger, serviceService)
	group := router.NewRouter(handlerHandler)
	engine := http.NewRouter(httpOptions, logger, group)
	server := http.New(httpOptions, logger, engine)
	minioOptions, err := minio.NewOptions(viper)
	if err != nil {
		return nil, err
	}
	client, err := minio.New(minioOptions)
	if err != nil {
		return nil, err
	}
	applicationApplication, err := user.New(userOptions, logger, server, client)
	if err != nil {
		return nil, err
	}
	return applicationApplication, nil
}

// wire.go:

var providerSet = wire.NewSet(user.ProviderSet, config.ProviderSet, logging.ProviderSet, http.ProviderSet, minio.ProviderSet, gorm.ProviderSet, router.ProviderSet, handler.ProviderSet, repository.ProviderSet, service.ProviderSet)
