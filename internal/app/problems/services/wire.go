// +build wireinject

package services

import (
	"github.com/Infinity-OJ/Server/internal/app/problems/repositories"
	"github.com/Infinity-OJ/Server/internal/pkg/config"
	"github.com/Infinity-OJ/Server/internal/pkg/database"
	"github.com/Infinity-OJ/Server/internal/pkg/log"
	"github.com/google/wire"
)

var testProviderSet = wire.NewSet(
	log.ProviderSet,
	config.ProviderSet,
	database.ProviderSet,
	ProviderSet,
)

func CreateUsersService(cf string, sto repositories.ProblemRepository) (ProblemsService, error) {
	panic(wire.Build(testProviderSet))
}
