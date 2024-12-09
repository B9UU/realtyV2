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
    l.*,
    ARRAY_REMOVE(ARRAY_AGG(DISTINCT a.amenity), NULL) AS amenities,
    (
        SELECT json_agg(to_json(ag_subquery))
        FROM (
            SELECT DISTINCT ag.*
            FROM agents ag_rel
            JOIN agent ag ON ag_rel.agent_id = ag.id
            WHERE ag_rel.listing_id = l.id
        ) AS ag_subquery
    ) AS agents,
    row_to_json(par) AS "plot_area_range"
FROM 
    listings l
LEFT JOIN 
    amenities am ON l.id = am.listing_id
LEFT JOIN 
    amenity a ON am.amenity_id = a.id
JOIN 
    plot_area_range par ON l.plot_area_range_id = par.id
GROUP BY 
    l.id, par.id;
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
func (p *PropertyStore) Add(property *models.Property) error {
	return nil

}
