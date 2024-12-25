package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (app *Application) failedValidationRespone(c echo.Context, errors map[string]string) error {
	return c.JSON(http.StatusBadRequest, map[string]interface{}{"errors": errors})
}
