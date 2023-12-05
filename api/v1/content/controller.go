// di controller ini kita hanya perlu mereturn sebuah json saja.
// yg dimana pada saat pembuatan func kita harus membuat type struct untuk
// merefrensikan ke service.
package content

import (
	"hexa/business/content"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Controller struct { //menginisialisasi nya menggunakan type struct
	service content.Service
}

func NewController(service content.Service) *Controller {
	return &Controller{
		service: service,
	}
}

// func (controller *Controller) GetContentbyID(id int) error {
// 	return Controller.service.GetContentbyID(id)
// }

func (controller *Controller) GetAll(c echo.Context) error {
	contents, err := controller.service.GetContent()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, contents)
}
