package validator

import (
    "github.com/go-playground/validator/v10"
    "github.com/labstack/echo/v4"
)

type CustomValidator struct {
    validator *validator.Validate
}

func NewValidator() echo.Validator {
    return &CustomValidator{validator: validator.New()}
}

func (cv *CustomValidator) Validate(i interface{}) error {
    return cv.validator.Struct(i)
}
