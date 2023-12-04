package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"hexa/config"
	"hexa/util"

	// "github.com/cloudinary/cloudinary-go/config"
	"github.com/labstack/echo/v4"
)

func main() {
	config := config.GetConfig()

	dbCon := util.NewConnectionDatabase(config)
	defer dbCon.CloseConnection()

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome!")
	})

	go func() {
		address := fmt.Sprintf(":%d", config.App.Port)
		if err := e.Start(address); err != nil {
			log.Fatal(err)
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
}
