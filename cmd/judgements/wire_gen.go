// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"github.com/infinity-oj/server/internal/app/judgements"
	"github.com/infinity-oj/server/internal/app/judgements/controllers"
	"github.com/infinity-oj/server/internal/app/judgements/grpcclients"
	"github.com/infinity-oj/server/internal/app/judgements/grpcservers"
	"github.com/infinity-oj/server/internal/app/judgements/repositories"
	"github.com/infinity-oj/server/internal/app/judgements/services"
	"github.com/infinity-oj/server/internal/pkg/app"
	"github.com/infinity-oj/server/internal/pkg/config"
	"github.com/infinity-oj/server/internal/pkg/consul"
	"github.com/infinity-oj/server/internal/pkg/database"
	"github.com/infinity-oj/server/internal/pkg/jaeger"
	"github.com/infinity-oj/server/internal/pkg/log"
	"github.com/infinity-oj/server/internal/pkg/transports/grpc"
	"github.com/infinity-oj/server/internal/pkg/transports/http"
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
	judgementsOptions, err := judgements.NewOptions(viper, logger)
	if err != nil {
		return nil, err
	}
	httpOptions, err := http.NewOptions(viper)
	if err != nil {
		return nil, err
	}
	databaseOptions, err := database.NewOptions(viper, logger)
	if err != nil {
		return nil, err
	}
	db, err := database.New(databaseOptions)
	if err != nil {
		return nil, err
	}
	judgementsRepository := repositories.NewMysqlJudgementsRepository(logger, db)
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
	filesClient, err := grpcclients.NewFilesClient(client)
	if err != nil {
		return nil, err
	}
	filesService := services.NewFilesService(filesClient)
	judgementsService := services.NewJudgementsService(logger, judgementsRepository, filesService)
	judgementController := controllers.NewJudgementsController(logger, judgementsService)
	initControllers := controllers.CreateInitControllersFn(judgementController)
	engine := http.NewRouter(httpOptions, logger, initControllers, tracer)
	apiClient, err := consul.New(consulOptions)
	if err != nil {
		return nil, err
	}
	server, err := http.New(httpOptions, logger, engine, apiClient)
	if err != nil {
		return nil, err
	}
	serverOptions, err := grpc.NewServerOptions(viper)
	if err != nil {
		return nil, err
	}
	grpcserversJudgementsService, err := grpcservers.NewJudgementsServer(logger, judgementsService)
	if err != nil {
		return nil, err
	}
	initServers := grpcservers.CreateInitServersFn(grpcserversJudgementsService)
	grpcServer, err := grpc.NewServer(serverOptions, logger, initServers, apiClient, tracer)
	if err != nil {
		return nil, err
	}
	application, err := judgements.NewApp(judgementsOptions, logger, server, grpcServer)
	if err != nil {
		return nil, err
	}
	return application, nil
}

// wire.go:

var providerSet = wire.NewSet(log.ProviderSet, config.ProviderSet, consul.ProviderSet, jaeger.ProviderSet, database.ProviderSet, http.ProviderSet, grpc.ProviderSet, grpcservers.ProviderSet, grpcclients.ProviderSet, repositories.ProviderSet, controllers.ProviderSet, services.ProviderSet, judgements.ProviderSet)
