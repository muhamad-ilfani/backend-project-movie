package user_usecase

import (
	"context"
	"movie-app/helpers"
	"movie-app/repository"
	"movie-app/usecases"

	ur "movie-app/repository/movieapps"
)

func (x *usecase) RegisterUser(
	ctx context.Context, req usecases.RegisterUserRequest) (
	res usecases.RegisterUserResponse, httpcode int, err error,
) {
	ctx, cancel := context.WithTimeout(ctx, x.Configuration.Timeout)
	defer cancel()

	tx, err := x.Postgresql.BeginTxx(ctx, nil)
	if err == nil && tx != nil {
		defer func() { err = new(repository.SQLTransaction).EndTx(tx, err) }()
	}

	hashPass := helpers.HashPass(req.Password)

	userPG := ur.NewRepository(tx)

	request := repository.RegisterUserRequest{
		UserName: req.UserName,
		Email:    req.Email,
		Password: string(hashPass),
		IsAdmin:  req.IsAdmin,
	}

	response, httpcode, err := userPG.RegisterUser(ctx, request)
	if err != nil {
		return res, httpcode, err
	}

	res = usecases.RegisterUserResponse{
		ID:       response.ID,
		UserName: response.UserName,
		Email:    response.Email,
	}

	return res, httpcode, err
}
