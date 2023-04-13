package repository

import (
	"context"
)

type UserRepo interface {
	RegisterUser(
		ctx context.Context, req RegisterUserRequest) (
		res RegisterUserResponse, httpcode int, err error,
	)
	GetUserByEmail(
		ctx context.Context, req GetUserByEmailRequest) (
		res GetUserByEmailResponse, httpcode int, err error,
	)
}

type RegisterUserRequest struct {
	ID       int64
	UserName string
	Email    string
	Password string
	IsAdmin  bool
}

type RegisterUserResponse struct {
	ID       int64
	UserName string
	Email    string
}

type GetUserByEmailRequest struct{ Email string }

type GetUserByEmailResponse struct {
	ID       int64
	Password string
	IsAdmin  bool
}
