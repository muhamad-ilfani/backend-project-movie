package app

import (
	"context"
	"movie-app/delivery"
	"movie-app/usecases"
	movie_usecase "movie-app/usecases/movies"
	user_usecase "movie-app/usecases/users"
	"time"
)

func (x *App) initService(ctx context.Context) {
	timeout := 55 * time.Second

	userusecase := user_usecase.New(
		user_usecase.Configuration{
			Timeout: timeout,
		},
		user_usecase.Depencency{
			Postgresql: x.DB,
		},
	)

	movieusecase := movie_usecase.New(
		movie_usecase.Configuration{
			Timeout: timeout,
		},
		movie_usecase.Depencency{
			Postgresql: x.DB,
		},
	)

	delivery.NewEchoHandler(ctx, x.Echo, struct {
		usecases.UserUseCase
		usecases.MoviesUseCase
	}{
		userusecase,
		movieusecase,
	})
}
