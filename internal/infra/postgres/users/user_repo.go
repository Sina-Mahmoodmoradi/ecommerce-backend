package users

import (
	"context"
	"errors"

	"github.com/Sina-Mahmoodmoradi/ecommerce-backend/internal/features/users/entity"
	"github.com/Sina-Mahmoodmoradi/ecommerce-backend/internal/features/users/repository"
	"github.com/Sina-Mahmoodmoradi/ecommerce-backend/internal/infra/postgres/models"
	"github.com/Sina-Mahmoodmoradi/ecommerce-backend/internal/pkg/apperrors"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type GormUserRepository struct {
	db *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) repository.UserRepository {
	return &GormUserRepository{db: db}
}


func (r *GormUserRepository) Create(ctx context.Context, u *entity.User) error {
	m := models.FromEntity(u)


	err := r.db.WithContext(ctx).Create(m).Error
	if err != nil {
		if isUniqueViolation(err) {
			return apperrors.New(apperrors.CodeEmailExists, "Email already exists", err)
		}
		return apperrors.New(apperrors.CodeInternal, "Failed to create user", err)
	}
	return nil
}


func isUniqueViolation(err error) bool {
    var pgErr *pq.Error
    if errors.As(err, &pgErr) {
        return pgErr.Code == "23505"
    }
    return false
}

