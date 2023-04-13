package movies_usecase

import (
	"context"
	"movie-app/repository"
	ur "movie-app/repository/movieapps"
	"movie-app/usecases"
)

func (x *usecase) RegisterMovie(
	ctx context.Context, req usecases.RegisterMovieRequest) (
	res usecases.RegisterMovieResponse, httpcode int, err error,
) {
	ctx, cancel := context.WithTimeout(ctx, x.Configuration.Timeout)
	defer cancel()

	tx, err := x.Postgresql.BeginTxx(ctx, nil)
	if err == nil && tx != nil {
		defer func() { err = new(repository.SQLTransaction).EndTx(tx, err) }()
	}

	moviePG := ur.NewRepository(tx)

	request := repository.RegisterMovieRequest{
		Title:       req.Title,
		Description: req.Description,
		Duration:    req.Duration,
		Artists:     req.Artists,
		Genres:      req.Genres,
		WatchURL:    req.WatchURL,
		CreatedBy:   req.CreatedBy,
	}

	response, httpcode, err := moviePG.RegisterMovie(ctx, request)
	if err != nil {
		return res, httpcode, err
	}

	res = usecases.RegisterMovieResponse{
		ID: response.ID,
	}

	return res, httpcode, err
}
