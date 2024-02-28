package jwt

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type JwtService struct {
	secret string
}

func NewJwtService(secretKey string) JwtService {
	return JwtService{
		secret: secretKey,
	}
}

func (s JwtService) GenerateJwt(userId string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 1).Unix(),
	})

	secretKey := []byte(s.secret)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
