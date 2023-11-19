// pkg/nessie/loan.go

package nessie

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	loansPath = "/loans"
)

// Loan represents a loan entity.
type Loan struct {
	ID             string `json:"_id,omitempty"`
	Type           string `json:"type,omitempty"`
	CreationDate   string `json:"creation_date,omitempty"`
	Status         string `json:"status,omitempty"`
	CreditScore    int    `json:"credit_score,omitempty"`
	MonthlyPayment int    `json:"monthly_payment,omitempty"`
	Amount         int    `json:"amount,omitempty"`
	Description    string `json:"description,omitempty"`
}

// GetAllLoans retrieves all loans associated with the specified account ID.
func (api *NessieAPI) GetAllLoans(accountID string) ([]Loan, error) {
	url := fmt.Sprintf("%s/accounts/%s%s?key=%s", api.BaseURL, accountID, loansPath, api.APIKey)

	resp, err := api.HTTPClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get loans: %s", resp.Status)
	}

	var loans []Loan
	if err := json.NewDecoder(resp.Body).Decode(&loans); err != nil {
		return nil, err
	}

	return loans, nil
}

// GetLoanByID retrieves a loan by its ID.
func (api *NessieAPI) GetLoanByID(loanID string) (*Loan, error) {
	url := fmt.Sprintf("%s%s/%s?key=%s", api.BaseURL, loansPath, loanID, api.APIKey)

	resp, err := api.HTTPClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get loan: %s", resp.Status)
	}

	var loan Loan
	if err := json.NewDecoder(resp.Body).Decode(&loan); err != nil {
		return nil, err
	}

	return &loan, nil
}

// CreateLoan creates a new loan for the specified account.
func (api *NessieAPI) CreateLoan(accountID string, loan *Loan) (*Loan, error) {
	url := fmt.Sprintf("%s/accounts/%s%s?key=%s", api.BaseURL, accountID, loansPath, api.APIKey)

	body, err := json.Marshal(loan)
	if err != nil {
		return nil, err
	}

	resp, err := api.HTTPClient.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("failed to create loan: %s", resp.Status)
	}

	var createdLoan Loan
	if err := json.NewDecoder(resp.Body).Decode(&createdLoan); err != nil {
		return nil, err
	}

	return &createdLoan, nil
}

// UpdateLoan updates an existing loan.
func (api *NessieAPI) UpdateLoan(loanID string, updatedLoan *Loan) error {
	url := fmt.Sprintf("%s%s/%s?key=%s", api.BaseURL, loansPath, loanID, api.APIKey)

	body, err := json.Marshal(updatedLoan)
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
		return fmt.Errorf("failed to update loan: %s", resp.Status)
	}

	return nil
}

// DeleteLoan deletes an existing loan.
func (api *NessieAPI) DeleteLoan(loanID string) error {
	url := fmt.Sprintf("%s%s/%s?key=%s", api.BaseURL, loansPath, loanID, api.APIKey)

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
		return fmt.Errorf("failed to delete loan: %s", resp.Status)
	}

	return nil
}
