package data

import (
	"time"

	"github.com/lib/pq"
)

type Property struct {
	ID                          int            `json:"id" db:"id"`
	PlacementType               string         `json:"placement_type" db:"placement_type"`
	NumberOfBathrooms           int            `json:"number_of_bathrooms" db:"number_of_bathrooms"`
	NumberOfBedrooms            int            `json:"number_of_bedrooms" db:"number_of_bedrooms"`
	NumberOfRooms               int            `json:"number_of_rooms" db:"number_of_rooms"`
	Amenities                   pq.StringArray `json:"amenities" db:"amenities"`
	RelevancySortOrder          int            `json:"relevancy_sort_order" db:"relevancy_sort_order"`
	EnergyLabel                 string         `json:"energy_label" db:"energy_label"`
	Availability                string         `json:"availability" db:"availability"`
	Type                        string         `json:"type" db:"type"`
	Zoning                      string         `json:"zoning" db:"zoning"`
	TimeStamp                   time.Time      `json:"time_stamp" db:"time_stamp"`
	ObjectType                  string         `json:"object_type" db:"object_type"`
	ConstructionType            string         `json:"construction_type" db:"construction_type"`
	PublishDateUtc              time.Time      `json:"publish_date_utc" db:"publish_date_utc"`
	PublishDate                 string         `json:"publish_date" db:"publish_date"`
	ObjectDetailPageRelativeURL string         `json:"object_detail_page_relative_url" db:"relative_url"`
}
