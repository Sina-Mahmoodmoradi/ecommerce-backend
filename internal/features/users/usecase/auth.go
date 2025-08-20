package usecase

import (
	"context"
	"time"

	"github.com/Sina-Mahmoodmoradi/ecommerce-backend/internal/features/users/entity"
	"github.com/Sina-Mahmoodmoradi/ecommerce-backend/internal/features/users/repository"
)




type TokenManager interface {
	CreateAccessToken(u *entity.User, now time.Time) (string, error)
	CreateRefreshToken(u *entity.User, now time.Time) (string, string, error) // token, tokenID
	RevokeRefreshToken(ctx context.Context, tokenID string) error
	ValidateRefreshToken(ctx context.Context, token string) (*entity.User, string, error) // user, tokenID
}

type TokenPair struct {
	AccessToken  string
	RefreshToken string
}

type AuthService struct {
	repo  repository.UserRepository
	tm    TokenManager
	nowFn func() time.Time
	hasher  PasswordHasher
}


func NewAuthService(repo repository.UserRepository, tm TokenManager, hasher PasswordHasher) *AuthService {
	return &AuthService{
		repo:  repo,
		tm:    tm,
		nowFn: time.Now,
		hasher: hasher,
	}
}

type RegisterInput struct {
	Email    string
	Password string
	FullName string
}


type PasswordHasher interface {
    Hash(password string) (string, error)
    Compare(hash string, password string) error
}


func (s *AuthService) Register(ctx context.Context, in RegisterInput) (*entity.User, TokenPair, error) {

	hash, err := s.hasher.Hash(in.Password)
	if err != nil {
		return nil, TokenPair{}, err
	}
	u := &entity.User{
		Email:        in.Email,
		PasswordHash: string(hash),
		FullName:     in.FullName,
		Role:         entity.RoleCustomer,
	}
	if err := s.repo.Create(ctx, u); err != nil {
		return nil, TokenPair{}, err
	}
	now := s.nowFn()
	access, err := s.tm.CreateAccessToken(u, now)
	if err != nil {
		return nil, TokenPair{}, err
	}
	refresh, _, err := s.tm.CreateRefreshToken(u, now)
	if err != nil {
		return nil, TokenPair{}, err
	}
	return u, TokenPair{AccessToken: access, RefreshToken: refresh}, nil
}