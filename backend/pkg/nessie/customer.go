// pkg/nessie/customer.go

package nessie

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Customer represents a Capital One customer.
type Customer struct {
	ID        string  `json:"_id"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Address   Address `json:"address"`
}

// GetCustomerByID retrieves a customer by ID from Nessie API using NessieAPI.
func (api *NessieAPI) GetCustomerByID(customerID string) (*Customer, error) {
	url := fmt.Sprintf("%s/customers/%s?key=%s", api.BaseURL, customerID, api.APIKey)

	resp, err := api.HTTPClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch customer: %s", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var customer Customer
	err = json.Unmarshal(body, &customer)
	if err != nil {
		return nil, err
	}

	return &customer, nil
}

// GetAllCustomers retrieves all customers from Nessie API using NessieAPI.
func (api *NessieAPI) GetAllCustomers() ([]Customer, error) {
	url := fmt.Sprintf("%s/customers?key=%s", api.BaseURL, api.APIKey)

	resp, err := api.HTTPClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch customers: %s", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var customers []Customer
	err = json.Unmarshal(body, &customers)
	if err != nil {
		return nil, err
	}

	return customers, nil
}

// CreateCustomer creates a new customer in the Nessie API using NessieAPI.
func (api *NessieAPI) CreateCustomer(firstName, lastName, streetNumber, streetName, city, state, zip string) error {
	url := fmt.Sprintf("%s/customers?key=%s", api.BaseURL, api.APIKey)

	address := Address{
		StreetNumber: streetNumber,
		StreetName:   streetName,
		City:         city,
		State:        state,
		Zip:          zip,
	}

	newCustomer := Customer{
		FirstName: firstName,
		LastName:  lastName,
		Address:   address,
	}

	customerJSON, err := json.Marshal(newCustomer)
	if err != nil {
		return err
	}

	resp, err := api.HTTPClient.Post(url, "application/json", bytes.NewBuffer(customerJSON))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("failed to create customer: %s", resp.Status)
	}

	return nil
}

// UpdateCustomer updates an existing customer in the Nessie API using NessieAPI.
func (api *NessieAPI) UpdateCustomer(customerID, streetNumber, streetName, city, state, zip string) error {
	url := fmt.Sprintf("%s/customers/%s?key=%s", api.BaseURL, customerID, api.APIKey)

	address := Address{
		StreetNumber: streetNumber,
		StreetName:   streetName,
		City:         city,
		State:        state,
		Zip:          zip,
	}

	updatedCustomer := Customer{
		Address: address,
	}

	customerJSON, err := json.Marshal(updatedCustomer)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(customerJSON))
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
		return fmt.Errorf("failed to update customer: %s", resp.Status)
	}

	return nil
}
