package user_usecase

import (
	"movie-app/usecases"
	"time"

	"github.com/jmoiron/sqlx"
)

func New(c Configuration, d Depencency) usecases.UserUseCase {
	return &usecase{c, d}
}

type Configuration struct {
	Timeout time.Duration
}

type Depencency struct {
	Postgresql *sqlx.DB
}

type usecase struct {
	Configuration
	Depencency
}
