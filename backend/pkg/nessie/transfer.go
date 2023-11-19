// pkg/nessie/transfer.go

package nessie

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const (
	transfersPath = "/transfers"
)

// Transfer represents a transfer entity.
type Transfer struct {
	ID              string    `json:"_id,omitempty"`
	Type            string    `json:"type,omitempty"`
	TransactionDate time.Time `json:"transaction_date,omitempty"`
	Status          string    `json:"status,omitempty"`
	Medium          string    `json:"medium,omitempty"`
	PayerID         string    `json:"payer_id,omitempty"`
	PayeeID         string    `json:"payee_id,omitempty"`
	Description     string    `json:"description,omitempty"`
}

// GetAllTransfers retrieves all transfers associated with a specific account.
func (api *NessieAPI) GetAllTransfers(accountID, transferType string) ([]Transfer, error) {
	url := fmt.Sprintf("%s/accounts/%s/transfers?key=%s&type=%s", api.BaseURL, accountID, api.APIKey, transferType)

	resp, err := api.HTTPClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get transfers: %s", resp.Status)
	}

	var transfers []Transfer
	if err := json.NewDecoder(resp.Body).Decode(&transfers); err != nil {
		return nil, err
	}

	return transfers, nil
}

// GetTransferByID retrieves a transfer by its ID.
func (api *NessieAPI) GetTransferByID(transferID string) (*Transfer, error) {
	url := fmt.Sprintf("%s/transfers/%s?key=%s", api.BaseURL, transferID, api.APIKey)

	resp, err := api.HTTPClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get transfer: %s", resp.Status)
	}

	var transfer Transfer
	if err := json.NewDecoder(resp.Body).Decode(&transfer); err != nil {
		return nil, err
	}

	return &transfer, nil
}

// CreateTransfer creates a new transfer.
func (api *NessieAPI) CreateTransfer(accountID string, transfer *Transfer) (*Transfer, error) {
	url := fmt.Sprintf("%s/accounts/%s/transfers?key=%s", api.BaseURL, accountID, api.APIKey)

	body, err := json.Marshal(transfer)
	if err != nil {
		return nil, err
	}

	resp, err := api.HTTPClient.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("failed to create transfer: %s", resp.Status)
	}

	var createdTransfer Transfer
	if err := json.NewDecoder(resp.Body).Decode(&createdTransfer); err != nil {
		return nil, err
	}

	return &createdTransfer, nil
}

// UpdateTransfer updates an existing transfer.
func (api *NessieAPI) UpdateTransfer(transferID string, updatedTransfer *Transfer) error {
	url := fmt.Sprintf("%s/transfers/%s?key=%s", api.BaseURL, transferID, api.APIKey)

	body, err := json.Marshal(updatedTransfer)
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
		return fmt.Errorf("failed to update transfer: %s", resp.Status)
	}

	return nil
}

// DeleteTransfer deletes an existing transfer.
func (api *NessieAPI) DeleteTransfer(transferID string) error {
	url := fmt.Sprintf("%s/transfers/%s?key=%s", api.BaseURL, transferID, api.APIKey)

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
		return fmt.Errorf("failed to delete transfer: %s", resp.Status)
	}

	return nil
}
