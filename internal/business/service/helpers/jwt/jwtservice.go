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

func (s JwtService) ParseJwt(tokenString string) (*jwt.Token, string, error) {
	var signingMethod = jwt.SigningMethodHS256

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if token.Method != signingMethod {
			return nil, errors.New("error opening jwt")
		}
		return []byte(s.secret), nil
	})
	if err != nil {
		return nil, "", err
	}

	issuerClaim := token.Claims.(jwt.MapClaims)["issuer"].(string)

	return token, issuerClaim, nil
}
