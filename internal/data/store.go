package data

import (
	"context"
	"realtyV2/internal/models"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
)

type Store struct {
	Property PropertyInterface
}

func NewStore(db *sqlx.DB, log zerolog.Logger) *Store {
	return &Store{
		Property: NewPropertyStore(db, log),
	}
}

type PropertyInterface interface {
	GetAll() ([]models.Property, error)
	GetById(id int) (models.Property, error)
	AddOne(listing models.Property) error
	Search(ctx context.Context, b []string) ([]models.Property, error)
	// InsertAmenities(ctx context.Context, tx *sqlx.Tx, amenities []string, listingID int) error
}
