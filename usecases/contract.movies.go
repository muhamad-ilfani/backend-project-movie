package usecases

import "context"

type MoviesUseCase interface {
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
	Limit       int64  `json:"limit"`
	Offset      int64  `json:"offset"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Artists     string `json:"artists"`
	Genres      string `json:"genres"`
}

type GetMovieData struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Duration    string `json:"duration"`
	Artists     string `json:"artists"`
	Genres      string `json:"genres"`
	WatchURL    string `json:"watch_url"`
	Viewer      int64  `json:"viewer"`
}
type GetAllMoviesResponse []GetMovieData

type RegisterMovieRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Duration    string `json:"duration"`
	Artists     string `json:"artists"`
	Genres      string `json:"genres"`
	WatchURL    string `json:"watch_url"`
	CreatedBy   string `json:"-"`
}

type RegisterMovieResponse struct {
	ID int64 `json:"id"`
}

type UpdateMovieByIDRequest struct {
	ID          int64  `json:"-"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Duration    string `json:"duration"`
	Artists     string `json:"artists"`
	Genres      string `json:"genres"`
	WatchURL    string `json:"watch_url"`
	Viewer      int64  `json:"viewer"`
	UpdatedBy   string `json:"-"`
}

type UpdateMovieByIDResponse struct{}

type GetMovieByIDRequest struct {
	ID int64 `json:"-"`
}

type GetMovieByIDResponse GetMovieData
