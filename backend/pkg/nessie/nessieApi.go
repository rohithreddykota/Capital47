package nessie

import "net/http"

var nessieAPI *NessieAPI

// NessieAPI represents the Nessie API client.
type NessieAPI struct {
	BaseURL    string
	APIKey     string
	HTTPClient *http.Client
}

// NewNessieAPI creates a new instance of the NessieAPI.
func NewNessieAPI(baseURL, apiKey string) *NessieAPI {
	return &NessieAPI{
		BaseURL:    baseURL,
		APIKey:     apiKey,
		HTTPClient: &http.Client{},
	}
}

// SetNessieAPI sets the initialized NessieAPI.
func SetNessieAPI(api *NessieAPI) {
	nessieAPI = api
}

// GetNessieAPI returns the initialized NessieAPI.
func GetNessieAPI() *NessieAPI {
	return nessieAPI
}
