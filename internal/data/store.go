package data

import (
	"context"
	"fmt"
	"realtyV2/internal/models"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
)

type Store struct {
	Property PropertyInterface
}

func NewStore(dbUrl string, log zerolog.Logger) (*Store, error) {
	db, err := sqlx.Open("postgres", dbUrl)
	if err != nil {
		fmt.Println("error: ", err)
		return nil, err
	}
	fmt.Println(db.DriverName())
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &Store{
		Property: NewPropertyStore(db, log),
	}, err
}

type PropertyInterface interface {
	GetAll() ([]models.Property, error)
	GetById(id int) (models.Property, error)
	AddOne(listing models.Property) error
	// InsertAmenities(ctx context.Context, tx *sqlx.Tx, amenities []string, listingID int) error

	Search(ctx context.Context, b []string) ([]models.Prop, error)
}
