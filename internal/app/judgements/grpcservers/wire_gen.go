// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package grpcservers

import (
	"github.com/infinity-oj/server/internal/app/judgements/services"
	"github.com/infinity-oj/server/internal/pkg/config"
	"github.com/infinity-oj/server/internal/pkg/database"
	"github.com/infinity-oj/server/internal/pkg/log"
	"github.com/google/wire"
)

// Injectors from wire.go:

func CreateJudgementsServer(cf string, service services.JudgementsService) (*JudgementsService, error) {
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
	judgementsService, err := NewJudgementsServer(logger, service)
	if err != nil {
		return nil, err
	}
	return judgementsService, nil
}

// wire.go:

var testProviderSet = wire.NewSet(log.ProviderSet, config.ProviderSet, database.ProviderSet, ProviderSet)
