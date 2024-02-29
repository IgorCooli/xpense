package jwt

import (
	"errors"
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
		"issuer": userId,
		"exp":    time.Now().Add(time.Hour * 1).Unix(),
	})

	secretKey := []byte(s.secret)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s JwtService) ParseJwt(tokenString string) (*jwt.Token, error) {
	// Defina o método de assinatura esperado
	var signingMethod = jwt.SigningMethodHS256

	// Faça o parsing do token JWT, especificando o método de assinatura
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if token.Method != signingMethod {
			return nil, errors.New("error opening jwt")
		}
		return []byte(s.secret), nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}
