package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type dds struct {
	Name string   `json:"name"`
	Age  int      `json:"age"`
	Url  []string `query:"url" validate:"required,min=1,dive,required,alpha"`
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
	dd.Url = i.Url
	properties, err := app.scraper.Properties("dd")
	if err != nil {
		app.log.Debug().Msgf("Unable to scrape, Error: %s", err.Error())
		return err
	}
	for _, prop := range properties {
		err = app.store.Property.AddOne(prop)
		if err != nil {
			app.log.Error().Caller().Msgf("unable to add one: %v", err.Error())
		}
	}
	// app.log.Debug().Msg("getting data")
	// properties, err := app.store.Property.GetAll()
	// if err != nil {
	// 	app.log.Fatal().Msg(err.Error())
	// 	return err
	// }
	return c.JSON(http.StatusOK, properties)
}
func (app *Application) GetPropertyById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	property, err := app.store.Property.GetById(id)
	if err != nil {
		app.log.Error().Msg(err.Error())
		return err
	}
	return c.JSON(http.StatusOK, property)
}
