// +build wireinject

package services

import (
	"github.com/infinity-oj/server/internal/app/files/repositories"
	"github.com/infinity-oj/server/internal/pkg/config"
	"github.com/infinity-oj/server/internal/pkg/database"
	"github.com/infinity-oj/server/internal/pkg/log"
	"github.com/google/wire"
)

var testProviderSet = wire.NewSet(
	log.ProviderSet,
	config.ProviderSet,
	database.ProviderSet,
	ProviderSet,
)

func CreateFilesService(cf string, sto repositories.FilesRepository) (FilesService, error) {
	panic(wire.Build(testProviderSet))
}
