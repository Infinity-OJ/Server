// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"github.com/Infinity-OJ/Server/internal/app/server"
	"github.com/Infinity-OJ/Server/internal/app/server/controllers"
	"github.com/Infinity-OJ/Server/internal/app/server/grpcclients"
	"github.com/Infinity-OJ/Server/internal/app/server/services"
	"github.com/Infinity-OJ/Server/internal/pkg/app"
	"github.com/Infinity-OJ/Server/internal/pkg/config"
	"github.com/Infinity-OJ/Server/internal/pkg/consul"
	"github.com/Infinity-OJ/Server/internal/pkg/jaeger"
	"github.com/Infinity-OJ/Server/internal/pkg/log"
	"github.com/Infinity-OJ/Server/internal/pkg/transports/grpc"
	"github.com/Infinity-OJ/Server/internal/pkg/transports/http"
	"github.com/google/wire"
)

// Injectors from wire.go:

func CreateApp(cf string) (*app.Application, error) {
	viper, err := config.New(cf)
	if err != nil {
		return nil, err
	}
	options, err := log.NewOptions(viper)
	if err != nil {
		return nil, err
	}
	logger, err := log.New(options)
	if err != nil {
		return nil, err
	}
	productsOptions, err := products.NewOptions(viper, logger)
	if err != nil {
		return nil, err
	}
	httpOptions, err := http.NewOptions(viper)
	if err != nil {
		return nil, err
	}
	consulOptions, err := consul.NewOptions(viper)
	if err != nil {
		return nil, err
	}
	configuration, err := jaeger.NewConfiguration(viper, logger)
	if err != nil {
		return nil, err
	}
	tracer, err := jaeger.New(configuration)
	if err != nil {
		return nil, err
	}
	clientOptions, err := grpc.NewClientOptions(viper, tracer)
	if err != nil {
		return nil, err
	}
	client, err := grpc.NewClient(consulOptions, clientOptions)
	if err != nil {
		return nil, err
	}
	usersClient, err := grpcclients.NewUsersClient(client)
	if err != nil {
		return nil, err
	}
	productsService := services.NewProductService(logger, usersClient)
	productsController := controllers.NewProductsController(logger, productsService)
	userService := services.NewUserService(logger, usersClient)
	usersController := controllers.NewUsersController(logger, userService)
	initControllers := controllers.CreateInitControllersFn(productsController, usersController)
	engine := http.NewRouter(httpOptions, logger, initControllers, tracer)
	apiClient, err := consul.New(consulOptions)
	if err != nil {
		return nil, err
	}
	server, err := http.New(httpOptions, logger, engine, apiClient)
	if err != nil {
		return nil, err
	}
	application, err := products.NewApp(productsOptions, logger, server)
	if err != nil {
		return nil, err
	}
	return application, nil
}

// wire.go:

var providerSet = wire.NewSet(log.ProviderSet, config.ProviderSet, consul.ProviderSet, jaeger.ProviderSet, http.ProviderSet, grpc.ProviderSet, grpcclients.ProviderSet, controllers.ProviderSet, services.ProviderSet, products.ProviderSet)
