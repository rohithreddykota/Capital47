// pkg/nessie/withdrawal.go

package nessie

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const (
	withdrawalsPath = "/withdrawals"
)

// Withdrawal represents a withdrawal entity.
type Withdrawal struct {
	ID              string    `json:"_id,omitempty"`
	Type            string    `json:"type,omitempty"`
	TransactionDate time.Time `json:"transaction_date,omitempty"`
	Status          string    `json:"status,omitempty"`
	PayerID         string    `json:"payer_id,omitempty"`
	Medium          string    `json:"medium,omitempty"`
	Amount          float64   `json:"amount,omitempty"`
	Description     string    `json:"description,omitempty"`
}

// GetAllWithdrawals retrieves all withdrawals associated with a specific account.
func (api *NessieAPI) GetAllWithdrawals(accountID string) ([]Withdrawal, error) {
	url := fmt.Sprintf("%s/accounts/%s/withdrawals?key=%s", api.BaseURL, accountID, api.APIKey)

	resp, err := api.HTTPClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get withdrawals: %s", resp.Status)
	}

	var withdrawals []Withdrawal
	if err := json.NewDecoder(resp.Body).Decode(&withdrawals); err != nil {
		return nil, err
	}

	return withdrawals, nil
}

// GetWithdrawalByID retrieves a withdrawal by its ID.
func (api *NessieAPI) GetWithdrawalByID(withdrawalID string) (*Withdrawal, error) {
	url := fmt.Sprintf("%s/withdrawals/%s?key=%s", api.BaseURL, withdrawalID, api.APIKey)

	resp, err := api.HTTPClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get withdrawal: %s", resp.Status)
	}

	var withdrawal Withdrawal
	if err := json.NewDecoder(resp.Body).Decode(&withdrawal); err != nil {
		return nil, err
	}

	return &withdrawal, nil
}

// CreateWithdrawal creates a new withdrawal.
func (api *NessieAPI) CreateWithdrawal(accountID string, withdrawal *Withdrawal) (*Withdrawal, error) {
	url := fmt.Sprintf("%s/accounts/%s/withdrawals?key=%s", api.BaseURL, accountID, api.APIKey)

	body, err := json.Marshal(withdrawal)
	if err != nil {
		return nil, err
	}

	resp, err := api.HTTPClient.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("failed to create withdrawal: %s", resp.Status)
	}

	var createdWithdrawal Withdrawal
	if err := json.NewDecoder(resp.Body).Decode(&createdWithdrawal); err != nil {
		return nil, err
	}

	return &createdWithdrawal, nil
}

// UpdateWithdrawal updates an existing withdrawal.
func (api *NessieAPI) UpdateWithdrawal(withdrawalID string, updatedWithdrawal *Withdrawal) error {
	url := fmt.Sprintf("%s/withdrawals/%s?key=%s", api.BaseURL, withdrawalID, api.APIKey)

	body, err := json.Marshal(updatedWithdrawal)
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
		return fmt.Errorf("failed to update withdrawal: %s", resp.Status)
	}

	return nil
}

// DeleteWithdrawal deletes an existing withdrawal.
func (api *NessieAPI) DeleteWithdrawal(withdrawalID string) error {
	url := fmt.Sprintf("%s/withdrawals/%s?key=%s", api.BaseURL, withdrawalID, api.APIKey)

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
		return fmt.Errorf("failed to delete withdrawal: %s", resp.Status)
	}

	return nil
}
