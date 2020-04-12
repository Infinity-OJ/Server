// +build wireinject

package grpcservers

import (
	"github.com/Infinity-OJ/Server/internal/app/judgements/services"
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

func CreateJudgementsServer(cf string, service services.JudgementsService) (*JudgementsService, error) {
	panic(wire.Build(testProviderSet))
}
