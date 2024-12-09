package models

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/lib/pq"
)

// to unmarshall JSON_AGG into Agent slice
type Agents []Agent

// property wrapper for custom Unmarshaller
type Properties []Property

type Agent struct {
	ID          int64  `db:"id" json:"id"`
	LogoType    string `db:"logo_type" json:"logo_type"`
	RelativeURL string `db:"relative_url" json:"relative_url"`
	IsPrimary   bool   `db:"is_primary" json:"is_primary"`
	LogoID      int    `db:"logo_id" json:"logo_id"`
	Name        string `db:"name" json:"name"`
	Association string `db:"association" json:"association"`
}

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
	Agents                      Agents         `json:"agent,omitempty" db:"agents"`
	PlotRange                   PlotAreaRange  `json:"plot_area_range" db:"plot_area_range"`
	PlogId                      int            `json:"plot_area_range_id" db:"plot_area_range_id"`
}
type PlotAreaRange struct {
	Gte int `json:"gte" db:"gte"`
	Lte int `json:"lte" db:"lte"`
}

func (a *PlotAreaRange) Scan(value interface{}) error {
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
		*p = append(*p, v.Source)
	}
	return nil
}

// func (a Agents) Value() (driver.Value, error) {
// 	return json.Marshal(a)
// }

func (a *Agents) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}
