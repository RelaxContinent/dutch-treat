package main

import (
	"dutch-treat/application"
	"dutch-treat/domain"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.Validator = domain.NewValidator()

	s := domain.NewDutchTreatService()
	h := application.NewDutchTreatHandler(s)

	e.GET("/", h.Test)
	e.POST("/test", h.DutchTreat)

	e.Logger.Fatal(e.Start(":8080"))
}
