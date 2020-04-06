// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"github.com/Infinity-OJ/Server/internal/app/files"
	"github.com/Infinity-OJ/Server/internal/app/files/controllers"
	"github.com/Infinity-OJ/Server/internal/app/files/grpcservers"
	"github.com/Infinity-OJ/Server/internal/app/files/repositories"
	"github.com/Infinity-OJ/Server/internal/app/files/services"
	"github.com/Infinity-OJ/Server/internal/pkg/app"
	"github.com/Infinity-OJ/Server/internal/pkg/config"
	"github.com/Infinity-OJ/Server/internal/pkg/consul"
	"github.com/Infinity-OJ/Server/internal/pkg/database"
	files2 "github.com/Infinity-OJ/Server/internal/pkg/files"
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
	filesOptions, err := files.NewOptions(viper, logger)
	if err != nil {
		return nil, err
	}
	serverOptions, err := grpc.NewServerOptions(viper)
	if err != nil {
		return nil, err
	}
	options2, err := files2.NewOptions(viper, logger)
	if err != nil {
		return nil, err
	}
	fileManager, err := files2.New(options2)
	if err != nil {
		return nil, err
	}
	filesRepository := repositories.NewFileManager(logger, fileManager)
	filesService := services.NewFileService(logger, filesRepository)
	filesServer, err := grpcservers.NewFilesServer(logger, filesService)
	if err != nil {
		return nil, err
	}
	initServers := grpcservers.CreateInitServersFn(filesServer)
	consulOptions, err := consul.NewOptions(viper)
	if err != nil {
		return nil, err
	}
	client, err := consul.New(consulOptions)
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
	server, err := grpc.NewServer(serverOptions, logger, initServers, client, tracer)
	if err != nil {
		return nil, err
	}
	application, err := files.NewApp(filesOptions, logger, server)
	if err != nil {
		return nil, err
	}
	return application, nil
}

// wire.go:

var providerSet = wire.NewSet(log.ProviderSet, config.ProviderSet, database.ProviderSet, services.ProviderSet, repositories.ProviderSet, consul.ProviderSet, jaeger.ProviderSet, http.ProviderSet, grpc.ProviderSet, files.ProviderSet, controllers.ProviderSet, grpcservers.ProviderSet, files2.ProviderSet)