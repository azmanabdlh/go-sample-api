package provider

import (
	"context"
	"errors"

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
