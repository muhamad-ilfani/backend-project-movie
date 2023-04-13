package delivery

import (
	"context"
	"movie-app/usecases"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func RegisterAdmin(ctx context.Context, uc usecases.UserUseCase) echo.HandlerFunc {
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

		form.IsAdmin = true

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

func RegisterMovie(ctx context.Context, uc usecases.MoviesUseCase) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx = c.Request().WithContext(ctx).Context()

		form := usecases.RegisterMovieRequest{}

		err := c.Bind(&form)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": FailedToUnmarshall,
				"error":   err.Error(),
			})
		}

		userData := c.Get("userData").(jwt.MapClaims)
		form.CreatedBy = userData["email"].(string)

		res, httpcode, err := uc.RegisterMovie(ctx, form)
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

func UpdateMovieByID(ctx context.Context, uc usecases.MoviesUseCase) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx = c.Request().WithContext(ctx).Context()
		form := usecases.UpdateMovieByIDRequest{}

		GetId, _ := strconv.Atoi(c.Param("id"))

		err := c.Bind(&form)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": FailedToUnmarshall,
				"error":   err.Error(),
			})
		}

		form.ID = int64(GetId)

		userData := c.Get("userData").(jwt.MapClaims)
		form.UpdatedBy = userData["email"].(string)

		res, httpcode, err := uc.UpdateMovieByID(ctx, form)
		if err != nil {
			return c.JSON(httpcode, map[string]interface{}{
				"message": "failed to update Movie",
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
