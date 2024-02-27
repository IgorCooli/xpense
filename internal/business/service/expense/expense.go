package service

import (
	"context"

	"github.com/IgorCooli/xpense/internal/business/model"
	repository "github.com/IgorCooli/xpense/internal/repository/expense"
)

type Service interface {
	InsertOne(ctx context.Context, expense model.Expense) error
}

type service struct {
	repository repository.Repository
}

func NewService(repository repository.Repository) Service {
	return service{
		repository: repository,
	}
}

func (s service) InsertOne(ctx context.Context, expense model.Expense) error {
	return s.repository.InsertOne(ctx, expense)
}
