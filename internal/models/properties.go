package models

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/lib/pq"
)

type Amenities struct {
	ListingID int `db:"listing_id"`
	AmenityID int `db:"amenity_id"`
}
type Amenity struct {
	ID   int    `db:"id"`
	Text string `db:"text"`
}

// to unmarshall JSON_AGG into Agent slice
type Agents []Agent

// property wrapper for custom Unmarshaller
type Properties []Property

type Property struct {
	ID                          int            `json:"id,omitempty" db:"id"`
	PlacementType               string         `json:"placement_type,omitempty" db:"placement_type"`
	NumberOfBathrooms           int            `json:"number_of_bathrooms,omitempty" db:"number_of_bathrooms"`
	NumberOfBedrooms            int            `json:"number_of_bedrooms,omitempty" db:"number_of_bedrooms"`
	NumberOfRooms               int            `json:"number_of_rooms,omitempty" db:"number_of_rooms"`
	Amenities                   pq.StringArray `json:"amenities,omitempty" db:"amenities"`
	RelevancySortOrder          int            `json:"relevancy_sort_order,omitempty" db:"relevancy_sort_order"`
	EnergyLabel                 string         `json:"energy_label,omitempty" db:"energy_label"`
	Availability                string         `json:"availability,omitempty" db:"availability"`
	Type                        string         `json:"type,omitempty" db:"type"`
	Zoning                      string         `json:"zoning,omitempty" db:"zoning"`
	TimeStamp                   time.Time      `json:"time_stamp,omitempty" db:"time_stamp"`
	ObjectType                  string         `json:"object_type,omitempty" db:"object_type"`
	ConstructionType            string         `json:"construction_type,omitempty" db:"construction_type"`
	PublishDateUtc              time.Time      `json:"publish_date_utc,omitempty" db:"publish_date_utc"`
	PublishDate                 string         `json:"publish_date,omitempty" db:"publish_date"`
	ObjectDetailPageRelativeURL string         `json:"object_detail_page_relative_url,omitempty" db:"relative_url"`
	PlotRange                   AreaRange      `json:"plot_area_range,omitempty" db:"plot_area_range"`
	FloorRange                  AreaRange      `json:"floor_area_range,omitempty" db:"floor_area_range"`
	Agents                      Agents         `json:"agents,omitempty" db:"agents"`
	PlotId                      int            `json:"plot_area_range_id,omitempty" db:"plot_area_range_id"`
	FloorId                     int            `json:"floor_area_range_id,omitempty" db:"floor_area_range_id"`
	Accessibility               pq.StringArray `json:"accessibility,omitempty" db:"accessibility"`
	MediaTypes                  pq.StringArray `json:"media_types,omitempty" db:"media_types"`
	Surrounding                 pq.StringArray `json:"surrounding,omitempty" db:"surrounding"`
	Address                     Address        `json:"address,omitempty" db:"address"`
	Parking                     pq.StringArray `json:"parking_facility,omitempty" db:"parking_facility"`
	SellPrice                   int            `json:"sell_price,omitempty" db:"sell_price"`
	RentPrice                   int            `json:"rent_price,omitempty" db:"rent_price"`
	OfferingType                pq.StringArray `json:"offering_type" db:"offering_type"`
	Thumb                       pq.Int64Array  `json:"thumbnail_id" db:"thumbnail_id"`
	Lon                         float32        `json:"lon" db:"lon"`
	Lat                         float32        `json:"lat" db:"lat"`
}
type AreaRange struct {
	Gte int `json:"gte" db:"gte"`
	Lte int `json:"lte" db:"lte"`
}

func (a *AreaRange) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}

type Agent struct {
	ID          int64  `db:"id" json:"id,omitempty"`
	LogoType    string `db:"logo_type" json:"logo_type,omitempty"`
	RelativeURL string `db:"relative_url" json:"relative_url,omitempty"`
	IsPrimary   bool   `db:"is_primary" json:"is_primary,omitempty"`
	LogoID      int    `db:"logo_id" json:"logo_id,omitempty"`
	Name        string `db:"name" json:"name,omitempty"`
	Association string `db:"association" json:"association,omitempty"`
}

func (a *Agents) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}

type Address struct {
	Country           string `json:"country,omitempty" db:"country"`
	Province          string `json:"province,omitempty" db:"province"`
	Wijk              string `json:"wijk,omitempty" db:"wijk"`
	City              string `json:"city,omitempty" db:"city"`
	Neighbourhood     string `json:"neighbourhood,omitempty" db:"neighbourhood"`
	HouseNumberSuffix string `json:"house_number_suffix" db:"house_number_suffix"`
	Municipality      string `json:"municipality,omitempty" db:"municipality"`
	IsBagAddress      bool   `json:"is_bag_address,omitempty" db:"is_bag_address"`
	HouseNumber       string `json:"house_number,omitempty" db:"house_number"`
	PostalCode        string `json:"postal_code,omitempty" db:"postal_code"`
	StreetName        string `json:"street_name,omitempty" db:"street_name"`
}

func (a *Address) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}
func (p *Properties) UnmarshalJSON(data []byte) error {
	temp := Response{}
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	for _, v := range temp.SearchResult.Hits.Hits {
		newP := Property{
			ID:                          v.Source.ID,
			PlacementType:               v.Source.PlacementType,
			NumberOfBathrooms:           v.Source.NumberOfBathrooms,
			NumberOfBedrooms:            v.Source.NumberOfBedrooms,
			NumberOfRooms:               v.Source.NumberOfRooms,
			Amenities:                   v.Source.Amenities,
			RelevancySortOrder:          v.Source.RelevancySortOrder,
			EnergyLabel:                 v.Source.EnergyLabel,
			Availability:                v.Source.Availability,
			Type:                        *v.Source.Type,
			Zoning:                      v.Source.Zoning,
			TimeStamp:                   v.Source.TimeStamp,
			ObjectType:                  v.Source.ObjectType,
			ConstructionType:            v.Source.ConstructionType,
			PublishDateUtc:              v.Source.PublishDateUtc,
			PublishDate:                 v.Source.PublishDate,
			ObjectDetailPageRelativeURL: v.Source.ObjectDetailPageRelativeURL,
			PlotRange:                   AreaRange(v.Source.PlotAreaRange),
			FloorRange:                  AreaRange(v.Source.FloorAreaRange),
			Agents:                      v.Source.Agent,
			Accessibility:               v.Source.Accessibility,
			MediaTypes:                  v.Source.AvailableMediaTypes,
			Surrounding:                 v.Source.Surrounding,
			Address:                     v.Source.Address,
			Parking:                     v.Source.ParkingFacility,
			SellPrice:                   v.Source.Price.SellingPriceRange.Lte,
			RentPrice:                   v.Source.Price.RentPriceRange.Lte,
			OfferingType:                v.Source.OfferingType,
			Thumb:                       v.Source.ThumbnailID,
		}
		newP.Lat = float32(v.Source.Location.Lat)
		newP.Lon = float32(v.Source.Location.Lon)
		*p = append(*p, newP)
	}
	return nil
}

type Prop struct {
	Id           int            `json:"id,omitempty"`
	ObjectType   string         `json:"object_type,omitempty"`
	OfferingType pq.StringArray `json:"offering_type,omitempty"`
	Type         string         `json:"type,omitempty"`
	Address      Address        `json:"address,omitempty"`
	RentPrice    int            `json:"rent_price,omitempty"`
	SellPrince   int            `json:"sell_prince,omitempty"`
}

// func (a Agents) Value() (driver.Value, error) {
// 	return json.Marshal(a)
// }
