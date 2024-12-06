package data

import (
	"fmt"

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
	GetAll() ([]Property, error)
}
