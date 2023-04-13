package repository

import "context"

type MovieRepo interface {
	GetAllMovies(
		ctx context.Context, req GetAllMoviesRequest) (
		res GetAllMoviesResponse, httpcode int, err error,
	)
	RegisterMovie(
		ctx context.Context, req RegisterMovieRequest) (
		res RegisterMovieResponse, httpcode int, err error,
	)
	UpdateMovieByID(
		ctx context.Context, req UpdateMovieByIDRequest) (
		res UpdateMovieByIDResponse, httpcode int, err error,
	)
	GetMovieByID(
		ctx context.Context, req GetMovieByIDRequest) (
		res GetMovieByIDResponse, httpcode int, err error,
	)
}

type GetAllMoviesRequest struct {
	Limit       int64
	Offset      int64
	Title       string
	Description string
	Artists     string
	Genres      string
}

type GetMovieData struct {
	ID          int64
	Title       string
	Description string
	Duration    string
	Artists     string
	Genres      string
	WatchURL    string
	Viewer      int64
}
type GetAllMoviesResponse []GetMovieData

type RegisterMovieRequest struct {
	Title       string
	Description string
	Duration    string
	Artists     string
	Genres      string
	WatchURL    string
	CreatedBy   string
}

type RegisterMovieResponse struct {
	ID int64
}

type UpdateMovieByIDRequest struct {
	ID          int64
	Title       string
	Description string
	Duration    string
	Artists     string
	Genres      string
	WatchURL    string
	Viewer      int64
	UpdatedBy   string
}

type UpdateMovieByIDResponse struct{}

type GetMovieByIDRequest struct{ ID int64 }
type GetMovieByIDResponse GetMovieData
