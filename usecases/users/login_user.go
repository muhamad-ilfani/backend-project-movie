package user_usecase

import (
	"context"
	"errors"
	"movie-app/helpers"
	"movie-app/repository"
	ur "movie-app/repository/movieapps"
	"movie-app/usecases"
	"net/http"
)

func (x *usecase) LoginUser(ctx context.Context, req usecases.LoginUserRequest) (
	res usecases.LoginUserResponse, httpcode int, err error,
) {
	ctx, cancel := context.WithTimeout(ctx, x.Configuration.Timeout)
	defer cancel()

	tx, err := x.Postgresql.BeginTxx(ctx, nil)
	if err == nil && tx != nil {
		defer func() { err = new(repository.SQLTransaction).EndTx(tx, err) }()
	}

	userPG := ur.NewRepository(tx)

	reqData := repository.GetUserByEmailRequest{
		Email: req.Email,
	}

	resData, httpcode, err := userPG.GetUserByEmail(ctx, reqData)
	if err != nil {
		return res, httpcode, err
	}

	if resData.ID == 0 {
		return res, http.StatusNotFound, errors.New("invalid email/password")
	}

	if comparePass := helpers.ComparePass([]byte(resData.Password), []byte(req.Password)); !comparePass {

		return res, http.StatusBadRequest, errors.New("invalid email/password")
	}

	token := helpers.GenerateToken(uint(resData.ID), req.Email, resData.IsAdmin)

	res = usecases.LoginUserResponse{
		Email: req.Email,
		Token: token,
	}

	return res, httpcode, err
}
