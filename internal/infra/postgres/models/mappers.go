package models

import (
	"github.com/Sina-Mahmoodmoradi/ecommerce-backend/internal/features/users/entity"
)

func ToEntity(m *UserModel) *entity.User {
	if m == nil {
		return nil
	}
	return &entity.User{
		ID:           m.ID,
		Email:        m.Email,
		PasswordHash: m.PasswordHash,
		FullName:     m.FullName,
		Role:         entity.Role(m.Role),
		CreatedAt:    m.CreatedAt,
		UpdatedAt:    m.UpdatedAt,
		DeletedAt:    m.DeletedAt,
	}

}

func FromEntity(e *entity.User) *UserModel {
	if e == nil {
		return nil
	}
	return &UserModel{
		ID:           e.ID,
		Email:        e.Email,
		PasswordHash: e.PasswordHash,
		FullName:     e.FullName,
		Role:         string(e.Role),
		CreatedAt:    e.CreatedAt,
		UpdatedAt:    e.UpdatedAt,
		DeletedAt:    e.DeletedAt,
	}
}