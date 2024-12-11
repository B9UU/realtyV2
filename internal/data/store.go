package data

import (
	"fmt"
	"realtyV2/internal/models"

	"github.com/jmoiron/sqlx"
)

type Store struct {
	Property PropertyInterface
}

func NewStore(dbUrl string) (*Store, error) {
	db, err := sqlx.Open("postgres", dbUrl)
	if err != nil {
		fmt.Println("error: ", err)
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &Store{
		Property: NewPropertyStore(db),
	}, err
}

type PropertyInterface interface {
	GetAll() ([]models.Property, error)
	GetById(id int) (models.Property, error)
	AddOne(listing models.Property) error
	// InsertAmenities(ctx context.Context, tx *sqlx.Tx, amenities []string, listingID int) error
}
