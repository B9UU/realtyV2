package data

import "github.com/jmoiron/sqlx"

type Store struct {
	*PropertyStore
}

func NewStore(dbUrl string) (*Store, error) {
	db, err := sqlx.Open("postgres", dbUrl)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &Store{
		NewPropertyStore(db),
	}, err
}
