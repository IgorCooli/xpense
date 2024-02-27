package card

import (
	"context"

	"github.com/IgorCooli/xpense/internal/business/model"
	"github.com/IgorCooli/xpense/internal/repository/card"
	"github.com/google/uuid"
)

type Service interface {
	RegisterCard(ctx context.Context, card model.Card) error
}

type service struct {
	repository card.Respository
}

func NewService(repository card.Respository) Service {
	return service{
		repository: repository,
	}
}

func (s service) RegisterCard(ctx context.Context, card model.Card) error {
	buildCardId(&card)
	return s.repository.InsertOne(ctx, card)
}

func buildCardId(user *model.Card) {
	UUID, err := uuid.NewUUID()

	if err != nil {
		panic("Could not generate uuid")
	}

	userId := UUID.String()

	user.ID = userId
}
