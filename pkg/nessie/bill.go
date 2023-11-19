// pkg/nessie/bill.go

package nessie

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// Bill represents the Bill model from Nessie API.
type Bill struct {
	ID                  string `json:"_id"`
	Status              string `json:"status"`
	Payee               string `json:"payee"`
	Nickname            string `json:"nickname"`
	CreationDate        string `json:"creation_date"`
	PaymentDate         string `json:"payment_date"`
	RecurringDate       int    `json:"recurring_date"`
	UpcomingPaymentDate string `json:"upcoming_payment_date"`
	AccountID           string `json:"account_id"`
}

// BillData represents the data needed to create or update a bill.
type BillData struct {
	Status        string `json:"status"`
	Payee         string `json:"payee"`
	Nickname      string `json:"nickname"`
	PaymentDate   string `json:"payment_date"`
	RecurringDate int    `json:"recurring_date"`
}

// GetBillsByAccount retrieves all bills for a specific account from Nessie API.
func (api *NessieAPI) GetBillsByAccount(accountID string) ([]Bill, error) {
	url := fmt.Sprintf("%s/accounts/%s/bills?key=%s", api.BaseURL, accountID)
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
		return nil, fmt.Errorf("failed to get bills for account, status code: %d", resp.StatusCode)
	}

	var bills []Bill
	err = json.NewDecoder(resp.Body).Decode(&bills)
	if err != nil {
		return nil, err
	}

	return bills, nil
}

// GetBillByID retrieves a bill by ID from Nessie API.
func (api *NessieAPI) GetBillByID(billID string) (*Bill, error) {
	url := fmt.Sprintf("%s/bills/%s?key=%s", api.BaseURL, billID, api.APIKey)
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
		return nil, fmt.Errorf("failed to get bill, status code: %d", resp.StatusCode)
	}

	var bill Bill
	err = json.NewDecoder(resp.Body).Decode(&bill)
	if err != nil {
		return nil, err
	}

	return &bill, nil
}

// GetBillsByCustomer retrieves bills associated with a specific customer from Nessie API.
func (api *NessieAPI) GetBillsByCustomer(customerID string) ([]Bill, error) {
	url := fmt.Sprintf("%s/customers/%s/bills?key=%s", api.BaseURL, customerID, api.APIKey)
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
		return nil, fmt.Errorf("failed to get bills for customer, status code: %d", resp.StatusCode)
	}

	var bills []Bill
	err = json.NewDecoder(resp.Body).Decode(&bills)
	if err != nil {
		return nil, err
	}

	return bills, nil
}

// CreateBill creates a bill for a specific account using Nessie API.
func (api *NessieAPI) CreateBill(accountID string, billData BillData) (*Bill, error) {
	url := fmt.Sprintf("%s/accounts/%s/bills?key=%s", api.BaseURL, accountID, api.APIKey)
	reqBody, err := json.Marshal(billData)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(reqBody))
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
		return nil, fmt.Errorf("failed to create bill, status code: %d", resp.StatusCode)
	}

	var createdBill Bill
	err = json.NewDecoder(resp.Body).Decode(&createdBill)
	if err != nil {
		return nil, err
	}

	return &createdBill, nil
}

// UpdateBill updates a specific bill by ID using Nessie API.
func (api *NessieAPI) UpdateBill(billID string, updatedData BillData) error {
	url := fmt.Sprintf("%s/bills/%s?key=%s", api.BaseURL, billID, api.APIKey)
	reqBody, err := json.Marshal(updatedData)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(reqBody))
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
		return fmt.Errorf("failed to update bill, status code: %d", resp.StatusCode)
	}

	return nil
}

// DeleteBill deletes a specific bill by ID using Nessie API.
func (api *NessieAPI) DeleteBill(billID string) error {
	url := fmt.Sprintf("%s/bills/%s?key=%s", api.BaseURL, billID, api.APIKey)
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
		return fmt.Errorf("failed to delete bill, status code: %d", resp.StatusCode)
	}

	return nil
}
