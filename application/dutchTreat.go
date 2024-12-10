package application

import (
	"go-capabilities-showcase/application/presenter"
	"go-capabilities-showcase/domain"
	"net/http"

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
	req, err := presenter.NewRequest(c)
	if err != nil {
		return err
	}

	res, err := h.dutchTreatService.DutchTreat(req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, *res)
}
