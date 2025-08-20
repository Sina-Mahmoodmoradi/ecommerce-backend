package usecase

import "errors"

var (
	ErrEmailExists   = errors.New("email already exists")
	ErrBadCreds      = errors.New("invalid email or password")
	ErrUserNotFound  = errors.New("user not found")
)