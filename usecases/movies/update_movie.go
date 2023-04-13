package movies_usecase

import (
	"context"
	"movie-app/repository"
	ur "movie-app/repository/movieapps"
	"movie-app/usecases"
)

func (x *usecase) UpdateMovieByID(
	ctx context.Context, req usecases.UpdateMovieByIDRequest) (
	res usecases.UpdateMovieByIDResponse, httpcode int, err error,
) {
	ctx, cancel := context.WithTimeout(ctx, x.Configuration.Timeout)
	defer cancel()

	tx, err := x.Postgresql.BeginTxx(ctx, nil)
	if err == nil && tx != nil {
		defer func() { err = new(repository.SQLTransaction).EndTx(tx, err) }()
	}

	moviePG := ur.NewRepository(tx)

	_, httpcode, err = moviePG.UpdateMovieByID(ctx,
		repository.UpdateMovieByIDRequest{
			ID:          req.ID,
			Title:       req.Title,
			Description: req.Description,
			Duration:    req.Duration,
			Artists:     req.Artists,
			Genres:      req.Genres,
			WatchURL:    req.WatchURL,
			Viewer:      req.Viewer,
			UpdatedBy:   req.UpdatedBy,
		})
	if err != nil {
		return res, httpcode, err
	}

	return res, httpcode, err
}
