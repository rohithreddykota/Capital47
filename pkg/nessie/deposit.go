// pkg/nessie/deposit.go

package nessie

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	depositsPath = "/deposits"
)

// Deposit represents a deposit entity.
type Deposit struct {
	ID              string `json:"_id,omitempty"`
	Type            string `json:"type,omitempty"`
	TransactionDate string `json:"transaction_date,omitempty"`
	Status          string `json:"status,omitempty"`
	PayeeID         string `json:"payee_id,omitempty"`
	Medium          string `json:"medium,omitempty"`
	Description     string `json:"description,omitempty"`
}

// GetAllDeposits retrieves all deposits associated with the specified account ID.
func (api *NessieAPI) GetAllDeposits(accountID string) ([]Deposit, error) {
	url := fmt.Sprintf("%s/accounts/%s%s?key=%s", api.BaseURL, accountID, depositsPath, api.APIKey)

	resp, err := api.HTTPClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get deposits: %s", resp.Status)
	}

	var deposits []Deposit
	if err := json.NewDecoder(resp.Body).Decode(&deposits); err != nil {
		return nil, err
	}

	return deposits, nil
}

// GetDepositByID retrieves a deposit by its ID.
func (api *NessieAPI) GetDepositByID(depositID string) (*Deposit, error) {
	url := fmt.Sprintf("%s%s/%s?key=%s", api.BaseURL, depositsPath, depositID, api.APIKey)

	resp, err := api.HTTPClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get deposit: %s", resp.Status)
	}

	var deposit Deposit
	if err := json.NewDecoder(resp.Body).Decode(&deposit); err != nil {
		return nil, err
	}

	return &deposit, nil
}

// CreateDeposit creates a new deposit for the specified account.
func (api *NessieAPI) CreateDeposit(accountID string, deposit *Deposit) error {
	url := fmt.Sprintf("%s/accounts/%s%s?key=%s", api.BaseURL, accountID, depositsPath, api.APIKey)

	body, err := json.Marshal(deposit)
	if err != nil {
		return err
	}

	resp, err := api.HTTPClient.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("failed to create deposit: %s", resp.Status)
	}

	return nil
}

// UpdateDeposit updates an existing deposit.
func (api *NessieAPI) UpdateDeposit(depositID string, updatedDeposit *Deposit) error {
	url := fmt.Sprintf("%s%s/%s?key=%s", api.BaseURL, depositsPath, depositID, api.APIKey)

	body, err := json.Marshal(updatedDeposit)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := api.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusAccepted {
		return fmt.Errorf("failed to update deposit: %s", resp.Status)
	}

	return nil
}

// DeleteDeposit deletes an existing deposit.
func (api *NessieAPI) DeleteDeposit(depositID string) error {
	url := fmt.Sprintf("%s%s/%s?key=%s", api.BaseURL, depositsPath, depositID, api.APIKey)

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
		return fmt.Errorf("failed to delete deposit: %s", resp.Status)
	}

	return nil
}
