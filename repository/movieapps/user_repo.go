package movieapps_repo

import (
	"context"
	"log"
	"movie-app/repository"
	"movie-app/repository/movieapps/query"
	"net/http"
)

func (x *PostgreSQLConn) RegisterUser(
	ctx context.Context, req repository.RegisterUserRequest) (
	res repository.RegisterUserResponse, httpcode int, err error,
) {
	var id int64

	query := query.RegisterUser
	args := List{
		req.UserName,
		req.Email,
		req.Password,
		req.IsAdmin,
	}

	rows, err := x.tc.Query(query, args...)
	if err != nil {
		return res, http.StatusInternalServerError, err
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id)
		if err != nil {
			log.Println(err)

			return res, http.StatusInternalServerError, err
		}
	}

	res = repository.RegisterUserResponse{
		ID:       id,
		UserName: req.UserName,
		Email:    req.Email,
	}

	return res, httpcode, err
}

func (x *PostgreSQLConn) GetUserByEmail(
	ctx context.Context, req repository.GetUserByEmailRequest) (
	res repository.GetUserByEmailResponse, httpcode int, err error,
) {
	query := query.GetUserByEmail
	args := List{
		req.Email,
	}

	rows, err := x.tc.Query(query, args...)
	if err != nil {
		return res, http.StatusInternalServerError, err
	}

	defer rows.Close()

	for rows.Next() {
		data := repository.GetUserByEmailResponse{}
		err := rows.Scan(&data.ID, &data.Password, &data.IsAdmin)
		if err != nil {
			log.Println(err)

			return res, http.StatusInternalServerError, err
		}
		res = data
	}

	return res, httpcode, err
}
