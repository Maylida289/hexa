package api

import (
	"hexa/api/v1/content"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	ContentController *content.Controller
}

func RegistrationPath(e *echo.Echo, controller Controller) {
	contentV1 := e.Group("/v1/content")
	contentV1.GET("", controller.ContentController.GetAll)
}
