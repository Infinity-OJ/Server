package services

import (
	"context"
	"fmt"

	proto "github.com/infinity-oj/api/protobuf-spec"
	"github.com/infinity-oj/server/internal/pkg/models"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type ProductsService interface {
	Get(c context.Context, ID uint64) (*models.User, error)
}

type DefaultProductsService struct {
	logger   *zap.Logger
	usersSrv proto.UsersClient
}

func NewProductService(logger *zap.Logger, usersSrv proto.UsersClient) ProductsService {
	return &DefaultProductsService{
		logger:   logger.With(zap.String("type", "DefaultProductsService")),
		usersSrv: usersSrv,
	}
}

func (s *DefaultProductsService) Get(c context.Context, productID uint64) (p *models.User, err error) {
	// get detail
	req := &proto.RegisterRequest{
		Username: "wycer",
		Email:    "wycers@gmail.com",
		Password: "123",
	}

	pd, err := s.usersSrv.CreateUser(c, req)
	if err != nil {
		return nil, errors.Wrap(err, "get rating error")
	}

	fmt.Println(pd.User.Password)

	return nil, nil
}
