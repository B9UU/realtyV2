package scraper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"realtyV2/internal/models"

	"github.com/rs/zerolog"
)

type Scraper struct {
	Log  zerolog.Logger
	Size int
}

func (s Scraper) Properties(q string, page int) ([]models.Property, error) {

	s.Log.Debug().Msg(q)
	url := "https://listing-search-wonen-arc.funda.io/listings-wonen-searcher-alias-prod/_reactivesearch"

	settings := map[string]interface{}{
		"settings": map[string]interface{}{
			"recordAnalytics":     false,
			"enableQueryRules":    true,
			"emptyQuery":          true,
			"suggestionAnalytics": true,
			"queryParams": map[string]interface{}{
				"preference": "_local",
			},
		},
		"query": []map[string]interface{}{
			{
				"id":        "search_result",
				"type":      "search",
				"dataField": []string{"availability"},
				"execute":   true,
				"react": map[string]interface{}{
					"and": []string{
						"selected_area", "offering_type", "sort", "price", "floor_area", "plot_area", "bedrooms",
						"rooms", "exterior_space_garden_size", "garage_capacity", "publication_date", "object_type",
						"availability", "accessibility", "construction_type", "construction_period", "surrounding",
						"garage_type", "exterior_space_type", "exterior_space_garden_orientation", "energy_label", "zoning",
						"amenities", "type", "nvm_open_house_day", "open_house", "free_text_search", "agent_id", "map_results",
						"object_type", "object_type_house_orientation", "object_type_house", "object_type_apartment_orientation",
						"object_type_apartment", "object_type_parking", "object_type_parking_capacity", "search_result__internal",
					},
				},
				"size": s.Size,
				"from": page * s.Size,
				"defaultQuery": map[string]interface{}{
					"track_total_hits": true,
					"timeout":          "1s",
					"sort": []map[string]interface{}{
						{"placement_type": "asc"},
						{"relevancy_sort_order": "desc"},
						{"id.number": "desc"},
					},
				},
			},
			{
				"id":        "selected_area",
				"type":      "term",
				"dataField": []string{"reactive_component_field"},
				"execute":   false,
				"customQuery": map[string]interface{}{
					"id":     "location-query-v2",
					"params": map[string]interface{}{"location": []string{q}},
				},
			},
			{
				"id":        "offering_type",
				"type":      "term",
				"dataField": []string{"offering_type"},
				"execute":   false,
				"defaultQuery": map[string]interface{}{
					"timeout": "500ms",
				},
				"value": "buy",
			},
			{
				"id":        "sort",
				"type":      "term",
				"dataField": []string{"sort"},
				"execute":   false,
				"defaultQuery": map[string]interface{}{
					"timeout": "500ms",
				},
			},
			{
				"id":        "price",
				"type":      "range",
				"dataField": []string{"price"},
				"execute":   false,
				"defaultQuery": map[string]interface{}{
					"timeout": "500ms",
				},
			},
			{
				"id":        "floor_area",
				"type":      "range",
				"dataField": []string{"floor_area"},
				"execute":   false,
				"defaultQuery": map[string]interface{}{
					"timeout": "500ms",
				},
			},
			{
				"id":        "plot_area",
				"type":      "range",
				"dataField": []string{"plot_area"},
				"execute":   false,
				"defaultQuery": map[string]interface{}{
					"timeout": "500ms",
				},
			},
			{
				"id":        "bedrooms",
				"type":      "range",
				"dataField": []string{"bedrooms"},
				"execute":   false,
				"defaultQuery": map[string]interface{}{
					"timeout": "500ms",
				},
			},
			{
				"id":        "rooms",
				"type":      "range",
				"dataField": []string{"rooms"},
				"execute":   false,
				"defaultQuery": map[string]interface{}{
					"timeout": "500ms",
				},
			},
			{
				"id":        "exterior_space_garden_size",
				"type":      "range",
				"dataField": []string{"exterior_space_garden_size"},
				"execute":   false,
				"defaultQuery": map[string]interface{}{
					"timeout": "500ms",
				},
			},
			{
				"id":        "garage_capacity",
				"type":      "range",
				"dataField": []string{"garage_capacity"},
				"execute":   false,
				"defaultQuery": map[string]interface{}{
					"timeout": "500ms",
				},
			},
			{
				"id":        "publication_date",
				"type":      "range",
				"dataField": []string{"publication_date"},
				"execute":   false,
				"defaultQuery": map[string]interface{}{
					"timeout": "500ms",
				},
			},
			{
				"id":        "object_type",
				"type":      "term",
				"dataField": []string{"object_type"},
				"execute":   false,
				"defaultQuery": map[string]interface{}{
					"timeout": "500ms",
				},
			},
			{
				"id":        "availability",
				"type":      "term",
				"dataField": []string{"availability"},
				"execute":   false,
				"defaultQuery": map[string]interface{}{
					"timeout": "500ms",
				},
			},
			{
				"id":        "accessibility",
				"type":      "term",
				"dataField": []string{"accessibility"},
				"execute":   false,
				"defaultQuery": map[string]interface{}{
					"timeout": "500ms",
				},
			},
			{
				"id":        "construction_type",
				"type":      "term",
				"dataField": []string{"construction_type"},
				"execute":   false,
				"defaultQuery": map[string]interface{}{
					"timeout": "500ms",
				},
			},
			{
				"id":        "construction_period",
				"type":      "term",
				"dataField": []string{"construction_period"},
				"execute":   false,
				"defaultQuery": map[string]interface{}{
					"timeout": "500ms",
				},
			},
			{
				"id":        "surrounding",
				"type":      "term",
				"dataField": []string{"surrounding"},
				"execute":   false,
				"defaultQuery": map[string]interface{}{
					"timeout": "500ms",
				},
			},
			{
				"id":        "garage_type",
				"type":      "term",
				"dataField": []string{"garage_type"},
				"execute":   false,
				"defaultQuery": map[string]interface{}{
					"timeout": "500ms",
				},
			},
			{
				"id":        "exterior_space_type",
				"type":      "term",
				"dataField": []string{"exterior_space_type"},
				"execute":   false,
				"defaultQuery": map[string]interface{}{
					"timeout": "500ms",
				},
			},
			{
				"id":        "exterior_space_garden_orientation",
				"type":      "term",
				"dataField": []string{"exterior_space_garden_orientation"},
				"execute":   false,
				"defaultQuery": map[string]interface{}{
					"timeout": "500ms",
				},
			},
			{
				"id":        "energy_label",
				"type":      "term",
				"dataField": []string{"energy_label"},
				"execute":   false,
				"defaultQuery": map[string]interface{}{
					"timeout": "500ms",
				},
			},
			{
				"id":        "zoning",
				"type":      "term",
				"dataField": []string{"zoning"},
				"execute":   false,
				"defaultQuery": map[string]interface{}{
					"timeout": "500ms",
				},
			},
			{
				"id":        "amenities",
				"type":      "term",
				"dataField": []string{"amenities"},
				"execute":   false,
				"defaultQuery": map[string]interface{}{
					"timeout": "500ms",
				},
			},
			{
				"id":        "type",
				"type":      "term",
				"dataField": []string{"type"},
				"execute":   false,
				"defaultQuery": map[string]interface{}{
					"timeout": "500ms",
				},
			},
			{
				"id":        "nvm_open_house_day",
				"type":      "term",
				"dataField": []string{"nvm_open_house_day"},
				"execute":   false,
				"defaultQuery": map[string]interface{}{
					"timeout": "500ms",
				},
			},
			{
				"id":        "open_house",
				"type":      "term",
				"dataField": []string{"open_house"},
				"execute":   false,
				"defaultQuery": map[string]interface{}{
					"timeout": "500ms",
				},
			},
			{
				"id":        "free_text_search",
				"type":      "search",
				"dataField": []string{"free_text_search"},
				"execute":   false,
				"defaultQuery": map[string]interface{}{
					"timeout": "500ms",
				},
			},
			{
				"id":        "agent_id",
				"type":      "term",
				"dataField": []string{"agent_id"},
				"execute":   false,
				"defaultQuery": map[string]interface{}{
					"timeout": "500ms",
				},
			},
			{
				"id":        "map_results",
				"type":      "term",
				"dataField": []string{"map_results"},
				"execute":   false,
				"defaultQuery": map[string]interface{}{
					"timeout": "500ms",
				},
			},
		},
	}
	payload, err := json.Marshal(settings)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	req.Header.Add("authority", "listing-search-wonen-arc.funda.io")
	req.Header.Add("accept", "application/json")
	req.Header.Add("accept-language", "en-US,en;q=0.6")

	// TODO: can be generate with basic auth, username:password available in the html in window.__nuxt__ script
	req.Header.Add("authorization", "Basic ZjVhMjQyZGIxZmUwOjM5ZDYxMjI3LWQ1YTgtNDIxMi04NDY4LWU1NWQ0MjhjMmM2Zg==")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("origin", "https://www.funda.nl")
	req.Header.Add("pragma", "no-cache")
	req.Header.Add("referer", "https://www.funda.nl/")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-site", "cross-site")
	req.Header.Add("sec-gpc", "1")
	req.Header.Add("user-agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36")
	req.Header.Add("x-search-client", "ReactiveSearch Vue")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Unable to send request")
	}
	df := models.Properties{}
	err = json.NewDecoder(res.Body).Decode(&df)
	if err != nil {
		return nil, err
	}
	return df, nil

}
