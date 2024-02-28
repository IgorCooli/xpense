package password

import (
	"golang.org/x/crypto/bcrypt"
)

type PasswordService struct {
}

func NewPasswordService() PasswordService {
	return PasswordService{}
}

func (s PasswordService) EncryptPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic("Erro ao criar a senha criptografada")
	}

	return string(hashedPassword)
}

func (s PasswordService) ValidatePassword(hashedPassword string, actualPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(actualPassword))
	if err != nil {
		return err
	}

	return nil
}
