package delivery

import (
	"context"
	"movie-app/usecases"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func RegisterUser(ctx context.Context, uc usecases.UserUseCase) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx = c.Request().WithContext(ctx).Context()

		form := usecases.RegisterUserRequest{}

		err := c.Bind(&form)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": FailedToUnmarshall,
				"error":   err.Error(),
			})
		}

		res, httpcode, err := uc.RegisterUser(ctx, form)
		if err != nil {
			return c.JSON(httpcode, map[string]interface{}{
				"message": FailedToRegisterUser,
				"error":   err.Error(),
			})
		}

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"status":  SuccessMsg,
			"message": SuccessMsg,
			"data":    res,
		})
	}
}

func LoginUser(ctx context.Context, uc usecases.UserUseCase) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx = c.Request().WithContext(ctx).Context()

		form := usecases.LoginUserRequest{}

		err := c.Bind(&form)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": FailedToUnmarshall,
				"error":   err.Error(),
			})
		}

		res, httpcode, err := uc.LoginUser(ctx, form)
		if err != nil {
			return c.JSON(httpcode, map[string]interface{}{
				"message": "failed to login",
				"error":   err.Error(),
			})
		}

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"status":  SuccessMsg,
			"message": SuccessMsg,
			"data":    res,
		})
	}
}

func GetAllMovies(ctx context.Context, uc usecases.MoviesUseCase) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx = c.Request().WithContext(ctx).Context()

		form := usecases.GetAllMoviesRequest{}

		err := c.Bind(&form)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": FailedToUnmarshall,
				"error":   err.Error(),
			})
		}

		res, httpcode, err := uc.GetAllMovies(ctx, form)
		if err != nil {
			return c.JSON(httpcode, map[string]interface{}{
				"message": FailedToGetData,
				"error":   err.Error(),
			})
		}

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"status":  SuccessMsg,
			"message": SuccessMsg,
			"data":    res,
		})
	}
}

func GetMovieByID(ctx context.Context, uc usecases.MoviesUseCase) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx = c.Request().WithContext(ctx).Context()

		GetId, _ := strconv.Atoi(c.Param("id"))

		form := usecases.GetMovieByIDRequest{
			ID: int64(GetId),
		}

		res, httpcode, err := uc.GetMovieByID(ctx, form)
		if err != nil {
			return c.JSON(httpcode, map[string]interface{}{
				"message": FailedToGetData,
				"error":   err.Error(),
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"status":  SuccessMsg,
			"message": SuccessMsg,
			"data":    res,
		})
	}
}
