package data

import (
	"fmt"

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

func (p *PropertyStore) GetAll() ([]Property, error) {
	properties := []Property{}
	fmt.Println("selecting")
	err := p.DB.Select(&properties, `
        SELECT 
          l.id, l.placement_type,
          l.number_of_bathrooms, l.number_of_bedrooms,
          l.number_of_rooms, l.relevancy_sort_order,
          l.energy_label, l.availability,
          l.type, l.zoning, l.time_stamp,
          l.object_type, l.construction_type,
          l.publish_date_utc, l.publish_date,
          l.relative_url, 
    ARRAY_REMOVE(ARRAY_AGG(a.amenity),NULL) AS amenities
        FROM 
          "listing" l
        LEFT JOIN 
          "amenities" am ON l.id = am.listing_id
        LEFT JOIN 
          "amenity" a ON am.amenity_id = a.id
        GROUP BY 
          l.id;
        `)
	if err != nil {
		return nil, err
	}
	return properties, nil
}
