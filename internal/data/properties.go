package data

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"realtyV2/internal/models"

	"github.com/jmoiron/sqlx"
)

var AlreadyExists = errors.New("Listing already exsists")

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
	err := p.DB.Select(&properties, stmt)
	if err != nil {
		return nil, err
	}
	return properties, nil
}
func (p *PropertyStore) GetById(id int) (models.Property, error) {

	stmt := `
SELECT 
    l.*,
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

WHERE l.id = $1

GROUP BY l.id, par.id, address.id;

`
	properties := models.Property{}
	err := p.DB.Get(&properties, stmt, id)
	if err != nil {
		return models.Property{}, err
	}
	return properties, nil
}

func (p *PropertyStore) AddOne(listing models.Property) error {
	ctx := context.TODO()
	stmt := `
	SELECT id FROM listings WHERE id = $1;
	`
	var lisId int
	err := p.DB.QueryRowContext(ctx, stmt, listing.ID).Scan(&lisId)
	if err == nil {
		return AlreadyExists
	}
	if err != sql.ErrNoRows {
		return err
	}
	tx, err := p.DB.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	plot_id, err := p.InsertPlotRange(ctx, tx, listing.PlotRange)
	if err != nil {
		return err
	}

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
  plot_area_range_id,
  price
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
 	:plot_area_range_id,
	:price


)
`, listing)
	if err != nil {
		return fmt.Errorf("Failed at inserting to Property: %v", err.Error())
	}
	err = p.InsertAttr(ctx, tx, "amenity", "amenities", listing.Amenities, listing.ID)
	if err != nil {
		return err
	}

	err = p.InsertAgents(ctx, tx, listing.Agents, listing.ID)
	if err != nil {
		return err
	}
	err = p.InsertAttr(ctx, tx, "accessibility", "accessibilities", listing.Accessibility, listing.ID)
	if err != nil {
		return err
	}
	err = p.InsertAttr(ctx, tx, "media_type", "media_types", listing.MediaTypes, listing.ID)
	if err != nil {
		return err
	}
	err = p.InsertAttr(ctx, tx, "surrounding", "surroundings", listing.Surrounding, listing.ID)
	if err != nil {
		return err
	}
	err = p.InsertAttr(ctx, tx, "parking", "parkings", listing.Parking, listing.ID)
	if err != nil {
		return err
	}
	err = p.InsertAddress(ctx, tx, listing.Address, listing.ID)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil

}

//	func (p *PropertyStore) InsertAmenities(ctx context.Context, tx *sqlx.Tx, amenities []string, listingID int) error {
//		for _, amenity := range amenities {
//			var amenityID int
//			stmt := `SELECT id from amenity WHERE text=$1;`
//			err := tx.QueryRowContext(ctx, stmt, amenity).Scan(&amenityID)
//			if err != nil {
//				if err == sql.ErrNoRows {
//					stmt = `INSERT INTO amenity (text) VALUES($1) RETURNING id`
//					err := tx.QueryRowContext(ctx, stmt, amenity).Scan(&amenityID)
//					if err != nil {
//						return fmt.Errorf("error inserting amenity %w", err)
//					}
//				} else {
//					return fmt.Errorf("error checking amenity %w", err)
//				}
//			}
//			stmt = `
//			INSERT INTO amenities (listing_id, amenity_id) VALUES($1,$2)
//			`
//			_, err = tx.ExecContext(ctx, stmt, listingID, amenityID)
//			if err != nil {
//				return fmt.Errorf("error inserting into amenities %w", err)
//			}
//
//		}
//
// return nil
// }
func (p *PropertyStore) InsertPlotRange(ctx context.Context, tx *sqlx.Tx, plot models.PlotAreaRange) (int, error) {
	var plot_id int
	stmt := `SELECT id FROM plot_area_range WHERE gte=$1 and lte=$2;`
	args := []interface{}{plot.Gte, plot.Lte}
	err := tx.QueryRowContext(ctx, stmt, args...).Scan(&plot_id)
	if err != nil && err != sql.ErrNoRows {
		return 0, err
	}
	stmt =
		`
		INSERT INTO plot_area_range (gte, lte)
		VALUES ($1, $2) RETURNING id;
	`
	err = tx.QueryRowContext(ctx, stmt, args...).Scan(&plot_id)
	if err != nil {
		return 0, err
	}
	return plot_id, nil
}

func (p *PropertyStore) InsertAgents(ctx context.Context, tx *sqlx.Tx, agents models.Agents, listingID int) error {
	for _, agent := range agents {
		var agentID int
		stmt := `SELECT id from agent WHERE id=$1;`
		err := tx.QueryRowContext(ctx, stmt, agent.ID).Scan(&agentID)
		if err != nil {
			if err == sql.ErrNoRows {
				stmt = `
					INSERT INTO agent (
						id, logo_type, relative_url,
						is_primary, logo_id, name, association
					) VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING id;`
				args := []interface{}{
					agent.ID, agent.LogoType, agent.RelativeURL,
					agent.IsPrimary, agent.LogoID, agent.Name, agent.Association}

				err = tx.QueryRowContext(ctx, stmt, args...).Scan(&agentID)
				if err != nil {
					return fmt.Errorf("error inserting agent %w", err)
				}
			} else {
				return fmt.Errorf("error checking agent %w", err)
			}
		}
		stmt = `
		INSERT INTO agents (listing_id, agent_id) VALUES($1,$2)
		`
		_, err = tx.ExecContext(ctx, stmt, listingID, agentID)
		if err != nil {
			return fmt.Errorf("error inserting into agents %w", err)
		}
	}
	return nil
}
func (p *PropertyStore) InsertAttr(ctx context.Context, tx *sqlx.Tx, table, tables string, attrs []string, listingID int) error {
	for _, attr := range attrs {
		var attrID int
		stmt := fmt.Sprintf(`SELECT id from %s WHERE text=$1;`, table)
		err := tx.QueryRowContext(ctx, stmt, attr).Scan(&attrID)
		if err != nil {
			if err == sql.ErrNoRows {
				stmt = fmt.Sprintf(`INSERT INTO %s (text) VALUES($1) RETURNING id`, table)
				err := tx.QueryRowContext(ctx, stmt, attr).Scan(&attrID)
				if err != nil {
					return fmt.Errorf("error inserting %s %w", table, err)
				}
			} else {
				return fmt.Errorf("error checking %s %w", table, err)
			}
		}
		stmt = fmt.Sprintf(`INSERT INTO %s (listing_id, %s_id) VALUES($1,$2)`, tables, table)
		_, err = tx.ExecContext(ctx, stmt, listingID, attrID)
		if err != nil {
			return fmt.Errorf("error inserting into %s %w", tables, err)
		}

	}
	return nil
}

func (p *PropertyStore) InsertAddress(ctx context.Context, tx *sqlx.Tx, address models.Address, listing_id int) error {

	var listingID int
	stmt := `SELECT listing_id from address WHERE listing_id=$1;`
	err := tx.QueryRowContext(ctx, stmt, listing_id).Scan(&listingID)
	if err != nil && err != sql.ErrNoRows {
		return fmt.Errorf("error checking address %w", err)
	}
	stmt = `
		INSERT INTO address (
		listing_id, country, province, wijk, city,
		neighbourhood, house_number_suffix, municipality,
		is_bag_address, house_number,postal_code, street_name
		) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12)
		`
	args := []interface{}{
		listing_id, address.Country, address.Province,
		address.Wijk, address.City, address.Neighbourhood,
		address.HouseNumberSuffix, address.Municipality,
		address.IsBagAddress, address.HouseNumber,
		address.PostalCode, address.StreetName}

	_, err = tx.ExecContext(ctx, stmt, args...)
	if err != nil {
		return fmt.Errorf("error inserting address %w", err)
	}
	return nil
}
