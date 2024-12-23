package models

import "time"

// the actual request respone,
// we use custom Unmarshaller to
// only get slice of Property (Properties)
type Response struct {
	Settings struct {
		Took       int    `json:"took"`
		SearchTook int    `json:"searchTook"`
		UserID     string `json:"userId"`
	} `json:"settings"`
	SearchResult struct {
		Took     int  `json:"took"`
		TimedOut bool `json:"timed_out"`
		Shards   struct {
			Total      int `json:"total"`
			Successful int `json:"successful"`
			Skipped    int `json:"skipped"`
			Failed     int `json:"failed"`
		} `json:"_shards"`
		Hits struct {
			Total struct {
				Value    int    `json:"value"`
				Relation string `json:"relation"`
			} `json:"total"`
			MaxScore any       `json:"max_score"`
			Hits     []Listing `json:"hits"`
		} `json:"hits"`
		Status int `json:"status"`
	} `json:"search_result"`
}

type Listing struct {
	Index string `json:"_index"`
	ID    string `json:"_id"`
	Score any    `json:"_score"`
	// Source Property `json:"_source"`
	Source struct {
		PlacementType      string   `json:"placement_type"`
		ConstructionPeriod string   `json:"construction_period"`
		Amenities          []string `json:"amenities"`
		Agent              Agents   `json:"agent"`
		NumberOfBathrooms  int      `json:"number_of_bathrooms"`
		PlotAreaRange      struct {
			Gte int `json:"gte"`
			Lte int `json:"lte"`
		} `json:"plot_area_range"`
		Blikvanger struct {
			Enabled bool `json:"enabled"`
		} `json:"blikvanger"`
		Accessibility      []string `json:"accessibility"`
		RelevancySortOrder int      `json:"relevancy_sort_order"`
		EnergyLabel        string   `json:"energy_label"`
		Description        struct {
			Dutch string `json:"dutch"`
			Tags  string `json:"tags"`
		} `json:"description"`
		FloorAreaRange struct {
			Gte int `json:"gte"`
			Lte int `json:"lte"`
		} `json:"floor_area_range"`
		Availability string  `json:"availability"`
		Type         *string `json:"type"`
		Price        struct {
			SellingPrice   []int `json:"selling_price"`
			RentPriceRange struct {
				Gte int `json:"gte"`
				Lte int `json:"lte"`
			} `json:"rent_price_range"`
			SellingPriceRange struct {
				Gte int `json:"gte"`
				Lte int `json:"lte"`
			} `json:"selling_price_range"`
			SellingPriceType      string `json:"selling_price_type"`
			SellingPriceCondition string `json:"selling_price_condition"`
		} `json:"price"`
		Zoning                   string   `json:"zoning"`
		ID                       int      `json:"id"`
		AvailableMediaTypes      []string `json:"available_media_types"`
		Surrounding              []string `json:"surrounding"`
		ObjectTypeSpecifications struct {
			House struct {
				Orientation string `json:"orientation"`
				Type        string `json:"type"`
			} `json:"house"`
		} `json:"object_type_specifications"`
		NumberOfBedrooms int       `json:"number_of_bedrooms"`
		Address          Address   `json:"address"`
		TimeStamp        time.Time `json:"time_stamp"`
		ObjectType       string    `json:"object_type"`
		ParkingFacility  []string  `json:"parking_facility"`
		Garage           struct {
			Type []any `json:"type"`
		} `json:"garage"`
		FloorArea        []int     `json:"floor_area"`
		ConstructionType string    `json:"construction_type"`
		ThumbnailID      []int64   `json:"thumbnail_id"`
		OfferingType     []string  `json:"offering_type"`
		PlotArea         []int     `json:"plot_area"`
		PublishDateUtc   time.Time `json:"publish_date_utc"`
		Location         struct {
			Lon float64 `json:"lon"`
			Lat float64 `json:"lat"`
		} `json:"location"`
		ExteriorSpace struct {
			GardenOrientation []string `json:"garden_orientation"`
			GardenSize        int      `json:"garden_size"`
			Type              []string `json:"type"`
		} `json:"exterior_space"`
		PublishDate                 string `json:"publish_date"`
		ObjectDetailPageRelativeURL string `json:"object_detail_page_relative_url"`
		Status                      string `json:"status"`
		NumberOfRooms               int    `json:"number_of_rooms"`
	} `json:"_source,"`
	Sort []any `json:"sort"`
}
