// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package services

import (
	"github.com/google/wire"
	"github.com/infinity-oj/api/protobuf-spec"
	"github.com/infinity-oj/server/internal/pkg/config"
	"github.com/infinity-oj/server/internal/pkg/log"
)

// Injectors from wire.go:

func CreateProductsService(cf string, usersClients protobuf_spec.UsersClient) (ProductsService, error) {
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
	productsService := NewProductService(logger, usersClients)
	return productsService, nil
}

// wire.go:

var testProviderSet = wire.NewSet(log.ProviderSet, config.ProviderSet, ProviderSet)
