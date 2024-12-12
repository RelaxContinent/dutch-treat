package presenter

import (
	"dutch-treat/domain/types"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Request struct {
	Members  []Member  `json:"member"`
	Payments []Payment `json:"payment"`
}

type Response struct {
	Payment Payment `json:"payment"`
}

type Member struct {
	Name types.Name `json:"name" validate:"required"`
}

type Payment struct {
	Payer  types.Name   `json:"payer" validate:"required"`
	Amount types.Amount `json:"amount" validate:"required"`
}

func NewRequest(c echo.Context) (*Request, error) {
	req := new(Request)
	if err := c.Bind(req); err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(req); err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return req, nil
}
