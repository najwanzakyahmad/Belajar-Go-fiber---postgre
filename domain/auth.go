package domain

import (
	"context"
	"go-fiber-postgre/dto"
)

type AuthService interface {
	Login(ctx context.Context, req dto.AuthRequest)(dto.AuthResponse, error)
}