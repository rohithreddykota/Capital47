// pkg/nessie/account.go

package nessie

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// Account represents the account model from Nessie API.
type Account struct {
	ID            string `json:"_id"`
	Type          string `json:"type"`
	Nickname      string `json:"nickname"`
	Rewards       int    `json:"rewards"`
	Balance       int    `json:"balance"`
	AccountNumber string `json:"account_number"`
	CustomerID    string `json:"customer_id"`
}

// AccountData represents the data needed to create an account.
type AccountData struct {
	Type          string `json:"type"`
	Nickname      string `json:"nickname"`
	Rewards       int    `json:"rewards"`
	Balance       int    `json:"balance"`
	AccountNumber string `json:"account_number"`
}

// GetAccount retrieves an account by ID from Nessie API.
func (api *NessieAPI) GetAccount(accountID string) (*Account, error) {
	url := fmt.Sprintf("%s/accounts/%s?key=%s", api.BaseURL, accountID, api.APIKey)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := api.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get account, status code: %d", resp.StatusCode)
	}

	var account Account
	err = json.NewDecoder(resp.Body).Decode(&account)
	if err != nil {
		return nil, err
	}

	return &account, nil
}

// GetAccountsByCustomer retrieves accounts associated with a specific customer ID from Nessie API.
func (api *NessieAPI) GetAccountsByCustomer(customerID string) ([]Account, error) {
	url := fmt.Sprintf("%s/customers/%s/accounts?key=%s", api.BaseURL, customerID, api.APIKey)
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
		return nil, fmt.Errorf("failed to get accounts, status code: %d", resp.StatusCode)
	}

	var accounts []Account
	err = json.NewDecoder(resp.Body).Decode(&accounts)
	if err != nil {
		return nil, err
	}

	return accounts, nil
}

// CreateAccount creates an account for a specific customer ID using Nessie API.
func (api *NessieAPI) CreateAccount(customerID string, accountData AccountData) (*Account, error) {
	url := fmt.Sprintf("%s/customers/%s/accounts?key=%s", api.BaseURL, customerID, api.APIKey)
	body, err := json.Marshal(accountData)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
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

	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("failed to create account, status code: %d", resp.StatusCode)
	}

	var newAccount Account
	err = json.NewDecoder(resp.Body).Decode(&newAccount)
	if err != nil {
		return nil, err
	}

	return &newAccount, nil
}

// UpdateAccount updates an account by ID using Nessie API.
func (api *NessieAPI) UpdateAccount(accountID string, updatedData AccountData) error {
	url := fmt.Sprintf("%s/accounts/%s?key=%s", api.BaseURL, accountID, api.APIKey)
	body, err := json.Marshal(updatedData)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Api-Key", api.APIKey)

	resp, err := api.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusAccepted {
		return fmt.Errorf("failed to update account, status code: %d", resp.StatusCode)
	}

	return nil
}

// DeleteAccount deletes an account by ID using Nessie API.
func (api *NessieAPI) DeleteAccount(accountID string) error {
	url := fmt.Sprintf("%s/accounts/%s?key=%s", api.BaseURL, accountID, api.APIKey)
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Api-Key", api.APIKey)

	resp, err := api.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("failed to delete account, status code: %d", resp.StatusCode)
	}

	return nil
}
