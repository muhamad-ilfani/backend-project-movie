package movieapps_repo

import (
	"movie-app/repository"

	"github.com/jmoiron/sqlx"
)

type List []interface{}

type PostgreSQLConn struct {
	tc *sqlx.Tx
}

type Repository interface {
	repository.UserRepo
	repository.MovieRepo
}

func NewRepository(tc *sqlx.Tx) Repository { return &PostgreSQLConn{tc} }
