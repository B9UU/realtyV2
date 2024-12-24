package data

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"realtyV2/internal/models"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
)

var ErrAlreadyExists = errors.New("Listing already exsists")

type PropertyStore struct {
	DB  *sqlx.DB
	Log zerolog.Logger
}

func NewPropertyStore(db *sqlx.DB, log zerolog.Logger) *PropertyStore {
	return &PropertyStore{
		DB:  db,
		Log: log,
	}

}

// TODO: search by location
// SELECT id
// FROM listings
// WHERE geohash && ST_MakeEnvelope(4.3793095, 51.8616672, 4.6018083, 51.9942816, 4326);

func (p *PropertyStore) GetAll() ([]models.Property, error) {
	stmt := `
SELECT 
    l.*,
    ARRAY_AGG(DISTINCT offering_type.text) FILTER (WHERE offering_type.text IS NOT NULL) as offering_type,
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

LEFT JOIN offering_types ON offering_types.listing_id = l.id
LEFT JOIN offering_type ON offering_type.id = offering_types.offering_type_id

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

    ARRAY_AGG(DISTINCT offering_type.text) FILTER (WHERE offering_type.text IS NOT NULL) as offering_type,
    ARRAY_AGG(DISTINCT mtp.text) FILTER (WHERE mtp.text IS NOT NULL) as media_types,
    ARRAY_AGG(DISTINCT amnt.text) FILTER (WHERE amnt.text IS NOT NULL) as amenities,
    ARRAY_AGG(DISTINCT acc.text) FILTER (WHERE acc.text IS NOT NULL) as accessibility,
    ARRAY_AGG(DISTINCT srnd.text) FILTER (WHERE srnd.text IS NOT NULL) as surrounding,
    ARRAY_AGG(DISTINCT parking.text) FILTER (WHERE parking.text IS NOT NULL) as parking_facility,

    COALESCE(row_to_json(par),'{}')as plot_area_range,
    COALESCE(json_agg(DISTINCT agnt) FILTER (WHERE agnt.id IS NOT NULL), '[]') AS agents,
    COALESCE(row_to_json(address),'{}') as address,
	
    ARRAY_AGG(DISTINCT thumbnail.img) FILTER (WHERE thumbnail.img IS NOT NULL) as thumbnail_id


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

LEFT JOIN offering_types ON offering_types.listing_id = l.id
LEFT JOIN offering_type ON offering_type.id = offering_types.offering_type_id

LEFT JOIN thumbnail ON thumbnail.listing_id = l.id

WHERE l.id = $1

GROUP BY l.id, par.id, address.id;

`
	properties := models.Property{ID: id}
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
	p.Log.Debug().Msgf("Checking listing: %d", listing.ID)
	var lisId int
	err := p.DB.QueryRowContext(ctx, stmt, listing.ID).Scan(&lisId)
	if err == nil {
		return ErrAlreadyExists
	}
	if err != sql.ErrNoRows {
		return err
	}
	tx, err := p.DB.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	plot_id, err := p.InsertRange(ctx, tx, listing.PlotRange, "plot")
	if err != nil {
		return err
	}

	floor_id, err := p.InsertRange(ctx, tx, listing.PlotRange, "floor")
	if err != nil {
		return err
	}
	listing.PlotId = plot_id
	listing.FloorId = floor_id

	p.Log.Debug().Msgf("Insert Listing: %d", listing.ID)
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
  floor_area_range_id,
  sell_price,
  rent_price,
  geohash
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
 	:floor_area_range_id,
	:sell_price,
	:rent_price,
	st_point(:lon, :lat)

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
	err = p.InsertAttr(ctx, tx, "offering_type", "offering_types", listing.OfferingType, listing.ID)
	if err != nil {
		return err
	}
	err = p.InsertAddress(ctx, tx, listing.Address, listing.ID)
	if err != nil {
		return err
	}
	err = p.InsertThumb(ctx, tx, listing.ID, listing.Thumb)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (p *PropertyStore) InsertRange(ctx context.Context, tx *sqlx.Tx, ARange models.AreaRange, rangeTpe string) (int, error) {
	var plot_id int
	stmt := fmt.Sprintf(`SELECT id FROM %s_area_range WHERE gte=$1 and lte=$2;`, rangeTpe)
	args := []interface{}{ARange.Gte, ARange.Lte}
	p.Log.Debug().Msg("Checking plot_area: ")
	err := tx.QueryRowContext(ctx, stmt, args...).Scan(&plot_id)
	if err != nil {
		if err != sql.ErrNoRows {
			return 0, err
		}
		stmt = fmt.Sprintf(
			`
		INSERT INTO %s_area_range (gte, lte)
		VALUES ($1, $2) RETURNING id;
		`, rangeTpe)
		p.Log.Debug().Msg("inserting plot_area:")
		err = tx.QueryRowContext(ctx, stmt, args...).Scan(&plot_id)
		if err != nil {
			return 0, err
		}
	}
	return plot_id, nil
}

func (p *PropertyStore) InsertAgents(ctx context.Context, tx *sqlx.Tx, agents models.Agents, listingID int) error {

	for _, agent := range agents {
		var agentID int
		stmt := `SELECT id from agent WHERE id=$1;`

		p.Log.Debug().Msgf("Checking agent: %d", agent.ID)
		err := tx.QueryRowContext(ctx, stmt, agent.ID).Scan(&agentID)
		if err != nil {
			if err != sql.ErrNoRows {
				return fmt.Errorf("error checking agent %w", err)
			}
			stmt = `
			INSERT INTO agent (
				id, logo_type, relative_url,
				is_primary, logo_id, name, association
			) VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING id;`
			args := []interface{}{
				agent.ID, agent.LogoType, agent.RelativeURL,
				agent.IsPrimary, agent.LogoID, agent.Name, agent.Association}

			p.Log.Debug().Msgf("inserting agent: %d", agent.ID)
			err = tx.QueryRowContext(ctx, stmt, args...).Scan(&agentID)
			if err != nil {
				return fmt.Errorf("error inserting agent %w", err)
			}
		}
		stmt = `
				INSERT INTO agents (listing_id, agent_id) VALUES($1,$2)
				`

		p.Log.Debug().Msgf("inserting agents: %d", agent.ID)
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

		p.Log.Debug().Msgf("Checking attr: %s", table)
		err := tx.QueryRowContext(ctx, stmt, attr).Scan(&attrID)
		if err != nil {
			if err != sql.ErrNoRows {
				return fmt.Errorf("error checking %s %w", table, err)
			}
			stmt = fmt.Sprintf(`INSERT INTO %s (text) VALUES($1) RETURNING id`, table)

			p.Log.Debug().Msgf("inserting to attr: %s", table)
			err = tx.QueryRowContext(ctx, stmt, attr).Scan(&attrID)
			if err != nil {
				return fmt.Errorf("error inserting %s %w", table, err)
			}
		}
		stmt = fmt.Sprintf(`INSERT INTO %s (listing_id, %s_id) VALUES($1,$2)`, tables, table)

		p.Log.Debug().Msgf("inserting to attr: %s", tables)
		_, err = tx.ExecContext(ctx, stmt, listingID, attrID)
		if err != nil {
			return fmt.Errorf("error inserting into %s %w", tables, err)
		}
	}
	return nil
}

func (p *PropertyStore) InsertAddress(ctx context.Context, tx *sqlx.Tx, address models.Address, listing_id int) error {

	stmt := `
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

	p.Log.Debug().Msg("inserting address")
	_, err := tx.ExecContext(ctx, stmt, args...)
	if err != nil {
		return fmt.Errorf("error inserting address %w", err)
	}
	return nil
}

func (p *PropertyStore) InsertThumb(ctx context.Context, tx *sqlx.Tx, listingId int, ids []int64) error {
	stmt := `
	INSERT INTO thumbnail(listing_id,img)
	VALUES($1,$2)
	`
	p.Log.Debug().Msg("inserting images")
	for _, img := range ids {
		args := []interface{}{listingId, img}
		_, err := tx.ExecContext(ctx, stmt, args...)
		if err != nil {
			p.Log.Debug().Msg(err.Error())
			continue
		}
	}
	return nil
}
func bindName() error {
	// q, args, err := sqlx.BindNamed(sqlx.BindType(p.DB.DriverName()), `SELECT listing_id from address WHERE listing_id=:id AND name=:placement_type;`, properties)
	// if err != nil {
	// 	return models.Property{}, err
	// }
	// fmt.Println(q, args)
	return nil
}
