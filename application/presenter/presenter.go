package presenter

import (
	"dutch-treat/domain/types"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Request struct {
	Payments []Payment `json:"payment"` // TODO 長さ1以上をOKとするvalidationをつける
}

type Response struct {
	Payment []Payment `json:"payment"`
}

type Payment struct {
	Payer  types.Name   `json:"payer" validate:"required"`
	Amount types.Amount `json:"amount" validate:"required"`
}

func NewRequest(c echo.Context) (*Request, error) {
	var req Request
	if err := c.Bind(&req); err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(&req); err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	fmt.Println("Request:", req)
	return &req, nil
}
