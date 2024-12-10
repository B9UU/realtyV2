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
			ARRAY_AGG(DISTINCT mtp.text) FILTER (WHERE mtp.text IS NOT NULL) as types,
			ARRAY_AGG(DISTINCT amnt.text) FILTER (WHERE amnt.text IS NOT NULL) as amenities,
			ARRAY_AGG(DISTINCT acc.text) FILTER (WHERE acc.text IS NOT NULL) as accessibility,
			ARRAY_AGG(DISTINCT srnd.text) FILTER (WHERE srnd.text IS NOT NULL) as surrounding,

			COALESCE(row_to_json(par),'{}')as plot_area_range,
			COALESCE(json_agg(DISTINCT agnt) FILTER (WHERE agnt.id IS NOT NULL), '[]') AS agents,
			COALESCE(row_to_json(address),'{}') as address

		FROM listings l
		LEFT JOIN amenities amnts ON amnts.listing_id = l.id
		LEFT JOIN amenity amnt ON amnts.amenity_id = amnt.id

		LEFT JOIN media_types mtps ON mtps.listing_id = l.id
		LEFT JOIN media_type mtp ON mtps.media_type_id = mtp.id

		LEFT JOIN agents agnts ON agnts.listing_id = l.id
		LEFT JOIN agent agnt ON agnts.agent_id = agnt.id

		LEFT JOIN accessibilities accs ON accs.listing_id =l.id
		LEFT JOIN accessibility acc ON acc.id = accs.accessibility_id

		LEFT JOIN surroundings srnds ON srnds.listing_id = l.id
		LEFT JOIN surrounding srnd ON srnds.surrounding_id = srnd.id

		LEFT JOIN address ON address.listing_id = l.id

		LEFT JOIN plot_area_range par ON par.id = l.plot_area_range_id

		GROUP BY l.id, par.id, address.id;

		`
	properties := []models.Property{}
	fmt.Println("selecting")
	err := p.DB.Select(&properties, stmt)
	if err != nil {
		fmt.Println("here")
		return nil, err
	}
	return properties, nil
}
func (p *PropertyStore) GetById(id int) (models.Property, error) {

	stmt := `
		SELECT 
			l.id,
			ARRAY_AGG(DISTINCT mtp.text) FILTER (WHERE mtp.text IS NOT NULL) as types,
			ARRAY_AGG(DISTINCT amnt.text) FILTER (WHERE amnt.text IS NOT NULL) as amenities,
			ARRAY_AGG(DISTINCT acc.text) FILTER (WHERE acc.text IS NOT NULL) as accessibility,
			ARRAY_AGG(DISTINCT srnd.text) FILTER (WHERE srnd.text IS NOT NULL) as surrounding,

			COALESCE(row_to_json(par),'{}')as plot_area_range,
			COALESCE(json_agg(DISTINCT agnt) FILTER (WHERE agnt.id IS NOT NULL), '[]') AS agents,
			COALESCE(row_to_json(address),'{}') as address

		FROM listings l
		LEFT JOIN amenities amnts ON amnts.listing_id = l.id
		LEFT JOIN amenity amnt ON amnts.amenity_id = amnt.id

		LEFT JOIN media_types mtps ON mtps.listing_id = l.id
		LEFT JOIN media_type mtp ON mtps.media_type_id = mtp.id

		LEFT JOIN agents agnts ON agnts.listing_id = l.id
		LEFT JOIN agent agnt ON agnts.agent_id = agnt.id

		LEFT JOIN accessibilities accs ON accs.listing_id =l.id
		LEFT JOIN accessibility acc ON acc.id = accs.accessibility_id

		LEFT JOIN surroundings srnds ON srnds.listing_id = l.id
		LEFT JOIN surrounding srnd ON srnds.surrounding_id = srnd.id

		LEFT JOIN address ON address.listing_id = l.id

		LEFT JOIN plot_area_range par ON par.id = l.plot_area_range_id
		where l.id = $1

		GROUP BY l.id, par.id, address.id;

		`
	properties := models.Property{}
	fmt.Println("selecting")
	err := p.DB.Get(&properties, stmt, 4)
	if err != nil {
		fmt.Println("here")
		return models.Property{}, err
	}
	return properties, nil
}
