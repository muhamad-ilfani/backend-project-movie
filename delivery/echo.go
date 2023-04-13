package delivery

import (
	"context"
	"movie-app/usecases"

	"github.com/labstack/echo"
)

const (
	TokenIsRequired              = "Token must be provided"
	SuccessMsg                   = "Success"
	WelcomeMsg                   = "welcome"
	FailedToUnmarshall           = "Failed to Unmarshall"
	FailedToRegisterUser         = "Failed to Register User"
	FailedToGetData              = "Failed to Get Data"
	FailedToUpdateDataActivities = "Failed to Update Data Activitiy"
	FailedToDeleteDataActivities = "Failed to Delete Data Activitiy"
	DeleteMsg                    = "Activity with ID %v Not Found"
)

type echoObject struct {
	*echo.Echo
	UseCase
}

type UseCase struct {
	usecases.UserUseCase
	usecases.MoviesUseCase
}

func NewEchoHandler(ctx context.Context, c *echo.Echo, uc UseCase) {
	obj := &echoObject{c, uc}
	obj.initRoute(ctx)

	obj.Logger.Fatal(obj.Start(":3030"))
}
