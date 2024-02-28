package user

import (
	"context"

	"github.com/IgorCooli/xpense/internal/business/model"
	"github.com/IgorCooli/xpense/internal/business/model/request"
	"github.com/IgorCooli/xpense/internal/business/service/helpers/jwt"
	"github.com/IgorCooli/xpense/internal/business/service/helpers/password"
	"github.com/IgorCooli/xpense/internal/repository/user"
	"github.com/google/uuid"
)

type Service interface {
	RegisterUser(ctx context.Context, user model.User) error
	AuthenticateUser(ctx context.Context, credentials request.Credentials) (string, error)
}

type service struct {
	repository      user.Respository
	passwordService password.PasswordService
	jwtService      jwt.JwtService
}

func NewService(repository user.Respository, passwordService password.PasswordService, jwtService jwt.JwtService) Service {
	return service{
		repository:      repository,
		passwordService: passwordService,
		jwtService:      jwtService,
	}
}

func (s service) RegisterUser(ctx context.Context, user model.User) error {
	buildUserId(&user)
	encryptPassword(&user, s)

	return s.repository.InsertOne(ctx, user)
}

func (s service) AuthenticateUser(ctx context.Context, credentials request.Credentials) (string, error) {
	user, err := s.repository.FindByUsername(ctx, credentials.Username)

	if err != nil {
		return "", err
	}

	passwordErr := s.passwordService.ValidatePassword(user.Password, credentials.Password)

	if passwordErr != nil {
		return "", passwordErr
	}

	return s.jwtService.GenerateJwt(user.ID)
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
