package data

import (
	"fmt"
	"realtyV2/internal/models"

	"github.com/jmoiron/sqlx"
)

type PropertyStore struct {
	DB *sqlx.DB
}

func NewPropertyStore(db *sqlx.DB) *PropertyStore {
	return &PropertyStore{
		DB: db,
	}

}

func (p *PropertyStore) GetAll() ([]models.Property, error) {
	stmt := `
	SELECT 
		l.id,
        ARRAY_REMOVE(ARRAY_AGG(a.amenity),NULL) AS amenities,
        json_agg(to_json(ag)) AS agents
        FROM 
          "listings" l
        LEFT JOIN 
          amenities am ON l.id = am.listing_id
        LEFT JOIN 
          amenity a ON am.amenity_id = a.id
        LEFT JOIN 
            agents ag_rel ON l.id = ag_rel.listing_id
        LEFT JOIN 
            agent ag ON ag_rel.agent_id = ag.id
        GROUP BY 
          l.id;
        `
	properties := []models.Property{}
	fmt.Println("selecting")
	err := p.DB.Select(&properties, stmt)
	if err != nil {
		fmt.Println("here")
		return nil, err
	}
	// row := p.DB.QueryRow(stmt)
	// err := row.Scan(&properties.ID, &properties.Amenities, &properties.Agents)
	// if err != nil {
	// 	return nil, err
	// }
	return properties, nil
}
