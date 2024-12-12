package main

import (
	"fmt"
	"net/http"
	"realtyV2/internal/data"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type queries struct {
	Page  int    `query:"page" validate:"numeric"`
	Query string `query:"q" validate:"required"`
}

var validate = validator.New()

func (app *Application) GetProperties(c echo.Context) error {
	q := new(queries)

	if err := c.Bind(q); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid input")
	}

	if err := validate.Struct(q); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	fmt.Println(q)
	properties, err := app.scraper.Properties(q.Query, q.Page)
	if err != nil {
		app.log.Debug().Msgf("Unable to scrape, Error: %s", err.Error())
		return err
	}
	for _, prop := range properties {
		err = app.store.Property.AddOne(prop)
		if err != nil {
			if err == data.AlreadyExists {
				app.log.Info().Msg("Already exists")
				continue
			}
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
