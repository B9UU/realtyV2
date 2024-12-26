package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type BoundBox struct {
	PlaceID     int      `json:"place_id"`
	Licence     string   `json:"licence"`
	Lat         string   `json:"lat"`
	Lon         string   `json:"lon"`
	Class       string   `json:"class"`
	Type        string   `json:"type"`
	PlaceRank   int      `json:"place_rank"`
	Importance  float64  `json:"importance"`
	Addresstype string   `json:"addresstype"`
	Name        string   `json:"name"`
	DisplayName string   `json:"display_name"`
	Boundingbox []string `json:"boundingbox"`
}

func getBoundBox(q string) ([]BoundBox, error) {
	url := fmt.Sprintf("https://nominatim.openstreetmap.org/search?q=%s&format=json", q)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed request")
	}
	bd := []BoundBox{}
	err = json.NewDecoder(res.Body).Decode(&bd)
	if err != nil {
		return nil, err
	}
	return bd, nil
}
