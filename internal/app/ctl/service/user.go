package service

import (
	"context"
	"fmt"

	proto "github.com/Infinity-OJ/Server/api/protobuf-spec"
	"github.com/Infinity-OJ/Server/internal/pkg/models"
	"github.com/pkg/errors"
)

type UserService interface {
	Create(username, password, email string) (*models.User, error)
}

type DefaultUserService struct {
	userSrv proto.UsersClient
}

func NewUserService(userSrv proto.UsersClient) UserService {
	return &DefaultUserService{
		userSrv: userSrv,
	}
}

func (s *DefaultUserService) Create(username, password, email string) (*models.User, error) {
	// get detail
	req := &proto.RegisterRequest{
		Username: username,
		Email:    email,
		Password: password,
	}

	pd, err := s.userSrv.CreateUser(context.TODO(), req)
	if err != nil {
		return nil, errors.Wrap(err, "get rating error")
	}

	fmt.Println(pd.User.Password)
	return nil, nil
}
