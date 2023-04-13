package query

import _ "embed"

var (
	//go:embed user/register_user.sql
	RegisterUser string
	//go:embed user/get_user_by_email.sql
	GetUserByEmail string

	//go:embed movie/get_all_movies_filter.sql
	GetAllMovies string
	//go:embed movie/get_movie_by_id.sql
	GetMovieByID string
	//go:embed movie/register_movie.sql
	RegisterMovie string
	//go:embed movie/update_movie_by_id.sql
	UpdateMovieByID string
)
