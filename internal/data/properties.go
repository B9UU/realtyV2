package data

import (
	"context"
	"database/sql"
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
    ARRAY_AGG(DISTINCT mtp.text) FILTER (WHERE mtp.text IS NOT NULL) as media_types,
    ARRAY_AGG(DISTINCT amnt.text) FILTER (WHERE amnt.text IS NOT NULL) as amenities,
    ARRAY_AGG(DISTINCT acc.text) FILTER (WHERE acc.text IS NOT NULL) as accessibility,
    ARRAY_AGG(DISTINCT srnd.text) FILTER (WHERE srnd.text IS NOT NULL) as surrounding,
    ARRAY_AGG(DISTINCT parking.text) FILTER (WHERE parking.text IS NOT NULL) as parking_facility,

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
LEFT JOIN parkings ON parkings.listing_id = l.id
LEFT JOIN parking ON parkings.parking_id = parking.id


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
    ARRAY_AGG(DISTINCT parking.text) FILTER (WHERE parking.text IS NOT NULL) as parking_facility,

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
LEFT JOIN parkings ON parkings.listing_id = l.id
LEFT JOIN parking ON parkings.parking_id = parking.id

WHERE l.id = $1

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

func (p *PropertyStore) AddOne(listing models.Property) error {
	ctx := context.TODO()
	tx, err := p.DB.Beginx()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
			return
		}
		err = tx.Commit()
	}()

	var plot_id int
	stmt :=
		`
		INSERT INTO plot_area_range (gte, lte)
		VALUES ($1, $2) RETURNING id;
	`
	err = tx.QueryRowContext(ctx, stmt,
		listing.PlotRange.Gte, listing.PlotRange.Lte).Scan(&plot_id)
	if err != nil {
		return err
	}

	fmt.Println("added plot")
	listing.PlogId = plot_id

	_, err = tx.NamedExecContext(ctx, `
INSERT INTO listings (
	id,
  placement_type,
  number_of_bathrooms,
  number_of_bedrooms,
  number_of_rooms,
  relevancy_sort_order,
  energy_label,
  availability,
  type,
  zoning,
  time_stamp,
  object_type,
  construction_type,
  publish_date_utc,
  publish_date,
  relative_url,
  plot_area_range_id
) 
VALUES(
	:id,
	:placement_type,
 	:number_of_bathrooms,
 	:number_of_bedrooms,
 	:number_of_rooms,
 	:relevancy_sort_order,
 	:energy_label,
 	:availability,
 	:type,
 	:zoning,
 	:time_stamp,
 	:object_type,
 	:construction_type,
 	:publish_date_utc,
 	:publish_date,
 	:relative_url,
 	:plot_area_range_id
)
`, listing)
	if err != nil {
		return fmt.Errorf("Failed at inserting to Property: %v", err.Error())
	}
	fmt.Println("added listing")
	err = p.InsertAmenities(ctx, tx, listing.Amenities, listing.ID)
	if err != nil {
		return err
	}
	return nil

}
func (p *PropertyStore) InsertAmenities(ctx context.Context, tx *sqlx.Tx, amenities []string, listingID int) error {
	for _, amenity := range amenities {
		var amenityID int
		stmt := `SELECT id from amenity WHERE text=$1;`
		err := tx.QueryRowContext(ctx, stmt, amenity).Scan(&amenityID)
		if err != nil {
			if err == sql.ErrNoRows {
				stmt = `INSERT INTO amenity (text) VALUES($1) RETURNING id`
				err := tx.QueryRowContext(ctx, stmt, amenity).Scan(&amenityID)
				if err != nil {
					return fmt.Errorf("error inserting amenity %w", err)
				}
			} else {
				return fmt.Errorf("error checking amenity %w", err)
			}
		}
		stmt = `
		INSERT INTO amenities (listing_id, amenity_id) VALUES($1,$2)
		`
		_, err = tx.ExecContext(ctx, stmt, listingID, amenityID)
		if err != nil {
			return fmt.Errorf("error inserting into amenities %w", err)
		}

	}
	return nil
}
