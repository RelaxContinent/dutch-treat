package main

import (
	"go-capabilities-showcase/application"
	"go-capabilities-showcase/domain"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	s := domain.NewDutchTreatService()
	h := application.NewDutchTreatHandler(s)

	e.POST("/", h.DutchTreat)

	e.Logger.Fatal(e.Start(":8080"))
}
