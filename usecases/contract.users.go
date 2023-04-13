package usecases

import "context"

type UserUseCase interface {
	RegisterUser(
		ctx context.Context, req RegisterUserRequest) (
		res RegisterUserResponse, httpcode int, err error,
	)
	LoginUser(
		ctx context.Context, req LoginUserRequest) (
		res LoginUserResponse, httpcode int, err error,
	)
}

type RegisterUserRequest struct {
	ID       int64  `json:"-"`
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	IsAdmin  bool   `json:"-"`
}

type RegisterUserResponse struct {
	ID       int64  `json:"id"`
	UserName string `json:"user_name"`
	Email    string `json:"email"`
}

type LoginUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginUserResponse struct {
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}
