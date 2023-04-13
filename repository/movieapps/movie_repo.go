package movieapps_repo

import (
	"context"
	"log"
	"movie-app/repository"
	"movie-app/repository/movieapps/query"
	"net/http"
)

func (x *PostgreSQLConn) GetAllMovies(
	ctx context.Context, req repository.GetAllMoviesRequest) (
	res repository.GetAllMoviesResponse, httpcode int, err error,
) {
	query := query.GetAllMovies
	args := List{
		req.Limit,
		req.Offset,
		req.Title,
		req.Description,
		req.Artists,
		req.Genres,
	}

	rows, err := x.tc.Query(query, args...)
	if err != nil {
		return res, http.StatusInternalServerError, err
	}

	defer rows.Close()

	for rows.Next() {
		data := repository.GetMovieData{}
		err := rows.Scan(&data.ID, &data.Title, &data.Description, &data.Duration, &data.Artists, &data.Genres, &data.WatchURL, &data.Viewer)
		if err != nil {
			log.Println(err)

			return res, http.StatusInternalServerError, err
		}
		res = append(res, data)
	}

	return res, httpcode, err
}

func (x *PostgreSQLConn) RegisterMovie(
	ctx context.Context, req repository.RegisterMovieRequest) (
	res repository.RegisterMovieResponse, httpcode int, err error,
) {
	var id int64

	query := query.RegisterMovie
	args := List{
		req.Title,
		req.Description,
		req.Duration,
		req.Artists,
		req.Genres,
		req.WatchURL,
		req.CreatedBy,
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

	res = repository.RegisterMovieResponse{
		ID: id,
	}

	return res, httpcode, err
}

func (x *PostgreSQLConn) UpdateMovieByID(
	ctx context.Context, req repository.UpdateMovieByIDRequest) (
	res repository.UpdateMovieByIDResponse, httpcode int, err error,
) {
	query := query.UpdateMovieByID
	args := List{
		req.ID,
		req.Title,
		req.Description,
		req.Duration,
		req.Artists,
		req.Genres,
		req.WatchURL,
		req.Viewer,
		req.UpdatedBy,
	}

	_, err = x.tc.Query(query, args...)
	if err != nil {
		return res, http.StatusInternalServerError, err
	}

	return res, httpcode, err
}

func (x *PostgreSQLConn) GetMovieByID(
	ctx context.Context, req repository.GetMovieByIDRequest) (
	res repository.GetMovieByIDResponse, httpcode int, err error,
) {
	query := query.GetMovieByID
	args := List{
		req.ID,
	}

	rows, err := x.tc.Query(query, args...)
	if err != nil {
		return res, http.StatusInternalServerError, err
	}

	defer rows.Close()

	for rows.Next() {
		data := repository.GetMovieByIDResponse{}
		err := rows.Scan(&data.ID, &data.Title, &data.Description, &data.Duration, &data.Artists, &data.Genres, &data.WatchURL, &data.Viewer)
		if err != nil {
			log.Println(err)

			return res, http.StatusInternalServerError, err
		}
		res = data
	}

	return res, httpcode, err
}
