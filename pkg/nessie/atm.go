// pkg/nessie/atm.go

package nessie

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// ATM represents the ATM model from Nessie API.
type ATM struct {
	ID            string   `json:"_id"`
	Name          string   `json:"name"`
	LanguageList  []string `json:"language_list"`
	Geocode       Geocode  `json:"geocode"`
	Hours         []string `json:"hours"`
	Accessibility bool     `json:"accessibility"`
	AmountLeft    int      `json:"amount_left"`
}

// Geocode represents the latitude and longitude for an ATM.
type Geocode struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

// GetATMs retrieves all ATMs within the specified search area from Nessie API.
func (api *NessieAPI) GetATMs(lat, lng float64, rad int) ([]ATM, error) {
	url := fmt.Sprintf("%s/atms?lat=%f&lng=%f&rad=%d?key=%s", api.BaseURL, lat, lng, rad, api.APIKey)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Api-Key", api.APIKey)

	resp, err := api.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get ATMs, status code: %d", resp.StatusCode)
	}

	var atms []ATM
	err = json.NewDecoder(resp.Body).Decode(&atms)
	if err != nil {
		return nil, err
	}

	return atms, nil
}

// GetATMByID retrieves an ATM by ID from Nessie API.
func (api *NessieAPI) GetATMByID(atmID string) (*ATM, error) {
	url := fmt.Sprintf("%s/atms/%s?key=%s", api.BaseURL, atmID, api.APIKey)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Api-Key", api.APIKey)

	resp, err := api.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get ATM, status code: %d", resp.StatusCode)
	}

	var atm ATM
	err = json.NewDecoder(resp.Body).Decode(&atm)
	if err != nil {
		return nil, err
	}

	return &atm, nil
}
