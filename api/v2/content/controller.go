package content

import (
	"hexa/api/v2/content/response"
	contentBusiness "hexa/business/content"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Controller struct { //menginisialisasi nya menggunakan type struct
	service contentBusiness.Service
}

func NewController(service contentBusiness.Service) *Controller {
	return &Controller{
		service: service,
	}
}

func (controller *Controller) GetAll(c echo.Context) error {
	contents, err := controller.service.GetContent()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed bro"})
	}

	response := response.NewGetContents(contents)

	return c.JSON(http.StatusOK, response)
}
