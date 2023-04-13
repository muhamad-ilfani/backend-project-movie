package movies_usecase

import (
	"context"
	"movie-app/repository"
	ur "movie-app/repository/movieapps"
	"movie-app/usecases"
)

func (x *usecase) GetAllMovies(
	ctx context.Context, req usecases.GetAllMoviesRequest) (
	res usecases.GetAllMoviesResponse, httpcode int, err error,
) {
	ctx, cancel := context.WithTimeout(ctx, x.Configuration.Timeout)
	defer cancel()

	if req.Limit <= 0 {
		req.Limit = 10
	}
	if req.Offset <= 0 {
		req.Offset = 0
	}

	tx, err := x.Postgresql.BeginTxx(ctx, nil)
	if err == nil && tx != nil {
		defer func() { err = new(repository.SQLTransaction).EndTx(tx, err) }()
	}

	moviePG := ur.NewRepository(tx)

	response, httpcode, err := moviePG.GetAllMovies(ctx, repository.GetAllMoviesRequest{
		Limit:       req.Limit,
		Offset:      req.Offset,
		Title:       req.Title,
		Description: req.Description,
		Artists:     req.Artists,
		Genres:      req.Genres,
	})
	if err != nil {
		return res, httpcode, err
	}

	for _, val := range response {
		res = append(res, usecases.GetMovieData{
			ID:          val.ID,
			Title:       val.Title,
			Description: val.Description,
			Duration:    val.Duration,
			Artists:     val.Artists,
			Genres:      val.Genres,
			WatchURL:    val.WatchURL,
			Viewer:      val.Viewer,
		})
	}

	return res, httpcode, err
}
