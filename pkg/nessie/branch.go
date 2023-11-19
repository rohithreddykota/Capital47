// pkg/nessie/branch.go

package nessie

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Branch represents a Capital One branch.
type Branch struct {
	ID          string   `json:"_id"`
	Name        string   `json:"name"`
	Hours       []string `json:"hours"`
	PhoneNumber string   `json:"phone_number"`
	Address     Address  `json:"address"`
}

// Address represents the address of a Capital One branch.
type Address struct {
	StreetNumber string `json:"street_number"`
	StreetName   string `json:"street_name"`
	City         string `json:"city"`
	State        string `json:"state"`
	Zip          string `json:"zip"`
}

// GetBranches retrieves all branches from Nessie API.
func (api *NessieAPI) GetBranches() ([]Branch, error) {
	url := fmt.Sprintf("%s/branches?key=%s", api.BaseURL, api.APIKey)
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
		return nil, fmt.Errorf("failed to get branches, status code: %d", resp.StatusCode)
	}

	var branches []Branch
	err = json.NewDecoder(resp.Body).Decode(&branches)
	if err != nil {
		return nil, err
	}

	return branches, nil
}

// GetBranchByID retrieves a branch by ID from Nessie API.
func (api *NessieAPI) GetBranchByID(branchID string) (*Branch, error) {
	url := fmt.Sprintf("%s/branches/%s?key=%s", api.BaseURL, branchID, api.APIKey)
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
		return nil, fmt.Errorf("failed to get branch, status code: %d", resp.StatusCode)
	}

	var branch Branch
	err = json.NewDecoder(resp.Body).Decode(&branch)
	if err != nil {
		return nil, err
	}

	return &branch, nil
}
