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
    ARRAY_REMOVE(ARRAY_AGG(DISTINCT a.text), NULL) AS amenities,
    ARRAY_REMOVE(ARRAY_AGG(DISTINCT accessibility.text), NULL) AS accessibility,
    json_agg(DISTINCT ag) AS agents,
    row_to_json(par) AS plot_area_range
FROM listings l
LEFT JOIN amenities am ON l.id = am.listing_id
LEFT JOIN amenity a ON am.amenity_id = a.id
LEFT JOIN accessibilities acc ON l.id = acc.listing_id
LEFT JOIN accessibility ON acc.accessibility_id = accessibility.id
JOIN plot_area_range par ON l.plot_area_range_id = par.id
LEFT JOIN LATERAL (
    SELECT DISTINCT ag.*
    FROM agents
    JOIN agent ag ON agents.agent_id = ag.id
    WHERE agents.listing_id = l.id
) ag ON true
GROUP BY l.id, par.id;
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
func (p *PropertyStore) GetById(id int) (models.Property, error) {

	stmt := `
SELECT 
    l.*,
    ARRAY_REMOVE(ARRAY_AGG(DISTINCT a.text), NULL) AS amenities,
    ARRAY_REMOVE(ARRAY_AGG(DISTINCT accessibility.text), NULL) AS accessibility,
    json_agg(DISTINCT ag) AS agents,
    row_to_json(par) AS plot_area_range
FROM listings l
LEFT JOIN amenities am ON l.id = am.listing_id
LEFT JOIN amenity a ON am.amenity_id = a.id
LEFT JOIN accessibilities acc ON l.id = acc.listing_id
LEFT JOIN accessibility ON acc.accessibility_id = accessibility.id
JOIN plot_area_range par ON l.plot_area_range_id = par.id
LEFT JOIN LATERAL (
    SELECT DISTINCT ag.*
    FROM agents
    JOIN agent ag ON agents.agent_id = ag.id
    WHERE agents.listing_id = l.id
) ag ON true
GROUP BY l.id, par.id;
`
	properties := models.Property{}
	fmt.Println("selecting")
	err := p.DB.Get(&properties, stmt, 1)
	if err != nil {
		fmt.Println("here")
		return models.Property{}, err
	}
	return properties, nil
}
