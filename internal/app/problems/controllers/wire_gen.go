// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package controllers

import (
	"github.com/google/wire"
	"github.com/infinity-oj/api/protobuf-spec"
	"github.com/infinity-oj/server/internal/app/problems/repositories"
	"github.com/infinity-oj/server/internal/app/problems/services"
	"github.com/infinity-oj/server/internal/pkg/config"
	"github.com/infinity-oj/server/internal/pkg/database"
	"github.com/infinity-oj/server/internal/pkg/log"
)

// Injectors from wire.go:

func CreateUsersController(cf string, sto repositories.ProblemRepository, client proto.FilesClient) (*ProblemController, error) {
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
	fileService := services.NewFileService(client)
	problemsService := services.NewProblemService(logger, sto, fileService)
	problemController := NewProblemsController(logger, problemsService)
	return problemController, nil
}

// wire.go:

var testProviderSet = wire.NewSet(log.ProviderSet, config.ProviderSet, database.ProviderSet, services.ProviderSet, ProviderSet)
