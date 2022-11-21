package models

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type (
    PostValidator struct {
        Validator *validator.Validate
    }
)

func (v *PostValidator) Validate(i interface{}) error {
    if err := v.Validator.Struct(i); err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }
    return nil
}

func NewValidator() *PostValidator {
    v := new(PostValidator)

    v.Validator = validator.New()
    return v
}
