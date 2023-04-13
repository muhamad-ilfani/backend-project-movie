package movies_usecase

import (
	"context"
	"movie-app/repository"
	ur "movie-app/repository/movieapps"
	"movie-app/usecases"
)

func (x *usecase) GetMovieByID(
	ctx context.Context, req usecases.GetMovieByIDRequest) (
	res usecases.GetMovieByIDResponse, httpcode int, err error,
) {
	ctx, cancel := context.WithTimeout(ctx, x.Configuration.Timeout)
	defer cancel()

	tx, err := x.Postgresql.BeginTxx(ctx, nil)
	if err == nil && tx != nil {
		defer func() { err = new(repository.SQLTransaction).EndTx(tx, err) }()
	}

	moviePG := ur.NewRepository(tx)

	response, httpcode, err := moviePG.GetMovieByID(ctx, repository.GetMovieByIDRequest{ID: req.ID})
	if err != nil {
		return res, httpcode, err
	}

	viewer := response.Viewer + 1

	_, httpcode, err = moviePG.UpdateMovieByID(ctx,
		repository.UpdateMovieByIDRequest{
			ID:     req.ID,
			Viewer: viewer,
		})

	res = usecases.GetMovieByIDResponse{
		ID:          response.ID,
		Title:       response.Title,
		Description: response.Description,
		Duration:    response.Duration,
		Artists:     response.Artists,
		Genres:      response.Genres,
		WatchURL:    response.WatchURL,
		Viewer:      viewer,
	}

	return res, httpcode, err
}
