package repository


import (
	"context"

	"github.com/Sina-Mahmoodmoradi/ecommerce-backend/internal/features/users/entity"
)

type UserRepository interface {
	Create(ctx context.Context, user *entity.User) error
	// FindByEmail(ctx context.Context, email string) (*entity.User, error)
	// GetByID(ctx context.Context, id string) (*entity.User, error)
}