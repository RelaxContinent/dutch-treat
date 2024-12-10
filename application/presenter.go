package application

import (
	"github.com/labstack/echo/v4"
)

type Request struct {
	Member  []Member  `json:"member"`
	Payment []Payment `json:"payment"`
}

type Response struct {
	Payment []Payment `json:"payment"`
}

type Member struct {
	Name string `json:"name"`
}

type Payment struct {
	Payer  string `json:"payer"`
	Amount int    `json:"amount"`
}

func NewRequest(c echo.Context) Request {
	var req Request
	c.Bind(&req)
	return req
}
