package main

import (
	"net/http"
	"realtyV2/internal/data"
	"realtyV2/internal/models"
	"realtyV2/internal/validator"
	"strconv"

	"github.com/labstack/echo/v4"
)

type queries struct {
	Page  int    `query:"page"`
	Query string `query:"q"`
}

func ValidateInput(v *validator.Validator, q *queries) {
	v.Check(q.Page > 0, "page", "must be greater than zero")
	v.Check(q.Page <= 10_000_000, "page", "must be between 1 - 10 million")

	v.Check(q.Query != "", "query", "must be provided")
	v.Check(len(q.Query) >= 3, "query", "must be longer than 3")

}
func (app *Application) GetProperties(c echo.Context) error {
	var v = validator.New()
	q := new(queries)
	if err := c.Bind(q); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid input")
	}
	if q.Page < 1 {
		q.Page = 1
	}
	if ValidateInput(v, q); !v.Valid() {
		return app.failedValidationRespone(c, v.Errors)

	}
	properties, err := app.scraper.Properties(q.Query, q.Page)
	if err != nil {
		app.log.Debug().Msgf("Unable to scrape, Error: %s", err.Error())
		return err
	}
	props := []models.Prop{}
	for _, prop := range properties {
		app.log.Debug().Msgf("Add %d", prop.ID)
		props = append(props,
			models.Prop{
				Id: prop.ID, ObjectType: prop.ObjectType,
				OfferingType: prop.OfferingType,
				Type:         prop.Type, Address: prop.Address,
				RentPrice: prop.RentPrice, SellPrince: prop.SellPrice,
			})
		err = app.store.Property.AddOne(prop)
		if err != nil {
			if err == data.ErrAlreadyExists {
				app.log.Debug().Msg("Already exists")
				continue
			}
			app.log.Error().Caller().Msgf("unable to add one: %v", err.Error())
		}
	}
	return c.JSON(http.StatusOK, props)
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
