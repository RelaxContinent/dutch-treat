package domain

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type Validator struct {
	validator *validator.Validate
}

func NewValidator() echo.Validator {
	return &Validator{validator: validator.New()}
}

func (v Validator) Validate(i interface{}) error {
	return v.validator.Struct(i)
}
