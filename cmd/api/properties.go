package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type dds struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Url  string `json:"url" validate:"required"`
}

func (app *Application) GetProperties(c echo.Context) error {
	dd := dds{
		Name: "laoq",
		Age:  22,
	}
	i := new(dds)
	if err := c.Bind(i); err != nil {
		app.log.Debug().Msg("unable to parse params")
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(i); err != nil {
		app.log.Debug().Msg("unable to validate")
		return err
	}
	return c.JSON(http.StatusOK, dd)

}
