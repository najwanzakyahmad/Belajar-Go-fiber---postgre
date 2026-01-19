package service

import (
	"context"
	"errors"
	"go-fiber-postgre/domain"
	"go-fiber-postgre/dto"
	"go-fiber-postgre/internal/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	conf           *config.Config
	userRepository domain.UserRepository
}

// Login implements [domain.AuthService].
func (a authService) Login(ctx context.Context, req dto.AuthRequest) (dto.AuthResponse, error) {
	user, err := a.userRepository.FindByEmail(ctx, req.Email)
	if err != nil {
		return dto.AuthResponse{}, err
	}

	if user.ID == "" {
		return dto.AuthResponse{}, errors.New("Authentication gagal tidak ditemukan user")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return dto.AuthResponse{}, errors.New("Authentication gagal salah password")
	}

	claim := jwt.MapClaims{
		"id" : user.ID,
		"exp" : time.Now().Add(time.Duration(a.conf.Jwt.Exp) * time.Minute).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenStr, err := token.SignedString([]byte(a.conf.Jwt.Key))
	if err != nil {
		return dto.AuthResponse{}, errors.New("Authentication gagal pembuatan token")
	}

	return dto.AuthResponse{
		Token: tokenStr,
	}, nil
}

func NewAuth(cnf *config.Config, userRepository domain.UserRepository) domain.AuthService {
	return authService{
		conf:           cnf,
		userRepository: userRepository,
	}
}
