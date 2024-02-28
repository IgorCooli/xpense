package user

import (
	"context"

	"github.com/IgorCooli/xpense/internal/business/model"
	"github.com/IgorCooli/xpense/internal/business/service/helpers/password"
	"github.com/IgorCooli/xpense/internal/repository/user"
	"github.com/google/uuid"
)

type Service interface {
	RegisterUser(ctx context.Context, user model.User) error
}

type service struct {
	repository      user.Respository
	passwordService password.PasswordService
}

func NewService(repository user.Respository, passwordService password.PasswordService) Service {
	return service{
		repository:      repository,
		passwordService: passwordService,
	}
}

func (s service) RegisterUser(ctx context.Context, user model.User) error {
	buildUserId(&user)
	encryptPassword(&user, s)

	return s.repository.InsertOne(ctx, user)
}

func buildUserId(user *model.User) {
	UUID, err := uuid.NewUUID()

	if err != nil {
		panic("Could not generate uuid")
	}

	userId := UUID.String()

	user.ID = userId
}

func encryptPassword(user *model.User, s service) {
	user.Password = s.passwordService.EncryptPassword(user.Password)
}
