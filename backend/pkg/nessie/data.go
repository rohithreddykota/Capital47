// pkg/nessie/data.go

package nessie

import (
	"fmt"
	"net/http"
)

const (
	dataPath  = "/data"
	typeParam = "type"
)

// DeleteData deletes data associated with the API key and the specified type using NessieAPI.
func (api *NessieAPI) DeleteData(dataType string) error {
	url := fmt.Sprintf("%s%s?key=%s&%s=%s", api.BaseURL, dataPath, api.APIKey, typeParam, dataType)

	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return err
	}

	resp, err := api.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("failed to delete data: %s", resp.Status)
	}

	return nil
}
