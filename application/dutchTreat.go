package application

import (
	"go-capabilities-showcase/domain"

	"github.com/labstack/echo/v4"
)

type DutchTreatHandler interface {
	DutchTreat(c echo.Context) error
}

type dutchTreatHandler struct {
	dutchTreatService domain.DutchTreatService
}

func NewDutchTreatHandler(dutchTreatService domain.DutchTreatService) DutchTreatHandler {
	return &dutchTreatHandler{dutchTreatService: dutchTreatService}
}

func (h dutchTreatHandler) DutchTreat(c echo.Context) error {
	h.dutchTreatService.DutchTreat()
	return nil
}
