package delivery

import (
	"context"
	"movie-app/middleware"

	"github.com/labstack/echo"
)

func (x *echoObject) initRoute(ctx context.Context) {
	x.Echo = echo.New()

	x.Echo.GET("/", welcome(ctx))
	x.Echo.POST("/register-admin", RegisterAdmin(ctx, x.UserUseCase))
	x.Echo.POST("/register-user", RegisterUser(ctx, x.UserUseCase))
	x.Echo.POST("/login", LoginUser(ctx, x.UseCase))

	user := x.Echo.Group("/user")
	user.Use(middleware.Authentication)
	user.POST("/", GetAllMovies(ctx, x.MoviesUseCase))
	user.GET("/:id", GetMovieByID(ctx, x.MoviesUseCase))

	admin := x.Echo.Group("/admin")
	admin.Use(middleware.Authentication, middleware.AdminAuthorization)
	admin.POST("/register-movie", RegisterMovie(ctx, x.MoviesUseCase))
	admin.PATCH("/update-movie/:id", UpdateMovieByID(ctx, x.MoviesUseCase))
}
