// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"github.com/google/wire"
	"github.com/infinity-oj/server/internal/app/submissions"
	"github.com/infinity-oj/server/internal/app/submissions/controllers"
	"github.com/infinity-oj/server/internal/app/submissions/grpcclients"
	"github.com/infinity-oj/server/internal/app/submissions/grpcservers"
	"github.com/infinity-oj/server/internal/app/submissions/repositories"
	"github.com/infinity-oj/server/internal/app/submissions/services"
	"github.com/infinity-oj/server/internal/pkg/app"
	"github.com/infinity-oj/server/internal/pkg/config"
	"github.com/infinity-oj/server/internal/pkg/consul"
	"github.com/infinity-oj/server/internal/pkg/database"
	"github.com/infinity-oj/server/internal/pkg/jaeger"
	"github.com/infinity-oj/server/internal/pkg/log"
	"github.com/infinity-oj/server/internal/pkg/transports/grpc"
	"github.com/infinity-oj/server/internal/pkg/transports/http"
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
	submissionsOptions, err := submissions.NewOptions(viper, logger)
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
	problemsClient, err := grpcclients.NewProblemsClient(client)
	if err != nil {
		return nil, err
	}
	problemsService := services.NewProblemService(problemsClient)
	databaseOptions, err := database.NewOptions(viper, logger)
	if err != nil {
		return nil, err
	}
	db, err := database.New(databaseOptions)
	if err != nil {
		return nil, err
	}
	submissionRepository := repositories.NewMysqlSubmissionsRepository(logger, db)
	filesClient, err := grpcclients.NewFilesClient(client)
	if err != nil {
		return nil, err
	}
	filesService := services.NewFilesService(filesClient)
	judgementsClient, err := grpcclients.NewJudgementsClient(client)
	if err != nil {
		return nil, err
	}
	judgementsService := services.NewJudgementsService(judgementsClient)
	submissionsService := services.NewSubmissionService(logger, problemsService, submissionRepository, filesService, judgementsService)
	submissionController := controllers.NewSubmissionsController(logger, submissionsService)
	initControllers := controllers.CreateInitControllersFn(submissionController)
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
	submissionService, err := grpcservers.NewSubmissionsServer(logger, submissionsService)
	if err != nil {
		return nil, err
	}
	initServers := grpcservers.CreateInitServersFn(submissionService)
	grpcServer, err := grpc.NewServer(serverOptions, logger, initServers, apiClient, tracer)
	if err != nil {
		return nil, err
	}
	application, err := submissions.NewApp(submissionsOptions, logger, server, grpcServer)
	if err != nil {
		return nil, err
	}
	return application, nil
}

// wire.go:

var providerSet = wire.NewSet(log.ProviderSet, config.ProviderSet, consul.ProviderSet, database.ProviderSet, jaeger.ProviderSet, http.ProviderSet, repositories.ProviderSet, grpc.ProviderSet, grpcservers.ProviderSet, grpcclients.ProviderSet, controllers.ProviderSet, services.ProviderSet, submissions.ProviderSet)
