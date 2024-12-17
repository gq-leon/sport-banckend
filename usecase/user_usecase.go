package usecase

import (
	"context"
	"time"

	"github.com/gq-leon/sport-backend/domain"
	"github.com/gq-leon/sport-backend/internal/tokenutil"
)

type userUseCase struct {
	contextTimeout time.Duration
	userRepository domain.UserRepository
}

func NewUserUseCase(repository domain.UserRepository, timeout time.Duration) domain.UserUseCase {
	return &userUseCase{
		contextTimeout: timeout,
		userRepository: repository,
	}
}

func (uu *userUseCase) Create(ctx context.Context, user *domain.User) error {
	ctx, cancel := context.WithTimeout(ctx, uu.contextTimeout)
	defer cancel()
	return uu.userRepository.Create(ctx, user)
}

func (uu *userUseCase) GetUserByEmail(ctx context.Context, email string) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, uu.contextTimeout)
	defer cancel()
	return uu.userRepository.GetByEmail(ctx, email)
}

func (uu *userUseCase) CreateAccessToken(user *domain.User, secret string, expiry int) (accessToken string, err error) {
	return tokenutil.CreateAccessToken(user, secret, expiry)
}

func (uu *userUseCase) CreateRefreshToken(user *domain.User, secret string, expiry int) (refreshToken string, err error) {
	return tokenutil.CreateRefreshToken(user, secret, expiry)
}
