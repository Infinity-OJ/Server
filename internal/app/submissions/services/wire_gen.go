// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package services

import (
	"github.com/google/wire"
	"github.com/infinity-oj/api/protobuf-spec"
	"github.com/infinity-oj/server/internal/app/submissions/repositories"
	"github.com/infinity-oj/server/internal/pkg/config"
	"github.com/infinity-oj/server/internal/pkg/database"
	"github.com/infinity-oj/server/internal/pkg/log"
)

// Injectors from wire.go:

func CreateSubmissionsService(cf string, sto repositories.SubmissionRepository, problemsClient protobuf_spec.ProblemsClient, filesClient protobuf_spec.FilesClient, judgementsClient protobuf_spec.JudgementsClient) (SubmissionsService, error) {
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
	problemsService := NewProblemService(problemsClient)
	filesService := NewFilesService(filesClient)
	judgementsService := NewJudgementsService(judgementsClient)
	submissionsService := NewSubmissionService(logger, problemsService, sto, filesService, judgementsService)
	return submissionsService, nil
}

// wire.go:

var testProviderSet = wire.NewSet(log.ProviderSet, config.ProviderSet, database.ProviderSet, ProviderSet)
