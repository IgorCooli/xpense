package user

import (
	"context"

	"github.com/IgorCooli/xpense/internal/business/model"
	"github.com/IgorCooli/xpense/internal/repository/user"
)

type Service interface {
	RegisterUser(ctx context.Context, user model.User) error
}

type service struct {
	repository user.Respository
}

func NewService(repository user.Respository) Service {
	return service{
		repository: repository,
	}
}

func (s service) RegisterUser(ctx context.Context, user model.User) error {
	return s.repository.InsertOne(ctx, user)
}
