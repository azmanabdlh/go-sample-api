package provider

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

type JsonWebToken struct {
	secret string
}

func NewJsonWebTokenProvider(secret string) *JsonWebToken {
	return &JsonWebToken{
		secret: secret,
	}
}

func (j *JsonWebToken) ValidateToken(
	ctx context.Context,
	token string,
) error {
	_, err := jwt.Parse(
		token,
		func(token *jwt.Token) (any, error) {
			return []byte(j.secret), nil
		},
	)

	if err != nil {
		return errors.New("invalid jwt token")
	}

	return nil
}

func (j *JsonWebToken) GenerateToken(
	userID string,
) (string, error) {

	now := time.Now()

	claims := jwt.MapClaims{
		"sub": userID,
		"iat": now.Unix(),
		"exp": now.Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	return token.SignedString(
		[]byte(j.secret),
	)
}
