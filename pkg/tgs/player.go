// Author: Omar Crosby
// File: player.go

package tgs

import (
	"encoding/json"
	"io"
	"net/http"
	"path"
)

// baseURL = "https://public.totalglobalsports.com/api/player"

// State represents a state in the Total Global Sports public API.
type State struct {
	ID         int    `json:"stateID"`
	RegionID   int    `json:"regionID"`
	CountryID  int    `json:"countryID"`
	Code       string `json:"stateCode"`
	Name       string `json:"stateName"`
	Image      string `json:"stateImage"`
	TimeZoneID string `json:"timeZoneID"`
}

// getJSON fetches JSON data from a URL and decodes it into the target interface.
func getJSON(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(r.Body)

	return json.NewDecoder(r.Body).Decode(target)
}

// GetAllStates fetches a list of all states from the Total Global Sports public API.
// It returns a slice of State structs populated with the data retrieved,
// or an error if the fetch operation fails.
func GetAllStates(baseUrl string) ([]State, error) {
	var states []State

	url := path.Join(baseUrl, "states")

	if err := getJSON(url, &states); err != nil {
		return nil, err
	}

	return states, nil
}
