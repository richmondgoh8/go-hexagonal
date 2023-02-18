package services

import (
	"context"
	"github.com/richmondgoh8/boilerplate/internal/core/domain"
	custommiddleware "github.com/richmondgoh8/boilerplate/pkg/middleware"
)

type TokenSvc struct{}

type TokenSvcImpl interface {
	GetJWTToken(ctx context.Context) (domain.Token, error)
}

// NewTokenSvc Takes in Interface, Return Struct
func NewTokenSvc() *TokenSvc {
	return &TokenSvc{}
}

func (srv *TokenSvc) GetJWTToken(ctx context.Context) (domain.Token, error) {
	tokenString, err := custommiddleware.GenerateToken(custommiddleware.JWTPayload{
		ID:   15,
		Role: "admin",
	})
	if err != nil {
		return domain.Token{}, err
	}

	return domain.Token{
		Token: tokenString,
	}, nil
}
