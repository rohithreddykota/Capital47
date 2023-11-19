package api

import (
	"Capital47/pkg/nessie"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

var nessieAPI *nessie.NessieAPI

// SetNessieAPI sets the initialized NessieAPI.
func SetNessieAPI(api *nessie.NessieAPI) {
	nessieAPI = api
}

// GetAccountHandler retrieves an account by ID.
func GetAccountHandler(api *nessie.NessieAPI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		accountID := vars["id"]

		account, err := api.GetAccount(accountID)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
			return
		}

		// Serialize account to JSON and send the response
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(account)
	}
}

// GetAccountsByCustomerHandler retrieves accounts associated with a specific customer ID.
func GetAccountsByCustomerHandler(api *nessie.NessieAPI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		customerID := vars["customerID"]

		accounts, err := api.GetAccountsByCustomer(customerID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Convert the accounts to JSON and write it to the response
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(accounts)
	}
}

// CreateAccountHandler creates a new account for a customer.
func CreateAccountHandler(api *nessie.NessieAPI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extract customer ID from the request parameters
		vars := mux.Vars(r)
		customerID := vars["customerID"]

		// Parse the request body to get the account data
		var accountData nessie.AccountData
		err := json.NewDecoder(r.Body).Decode(&accountData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Use the NessieAPI to create a new account
		createdAccount, err := api.CreateAccount(customerID, accountData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Convert created account to JSON and send the response
		responseJSON, err := json.Marshal(createdAccount)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(responseJSON)
	}
}

// UpdateAccountHandler updates an existing account.
func UpdateAccountHandler(api *nessie.NessieAPI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extract account ID from the request parameters
		vars := mux.Vars(r)
		accountID := vars["accountID"]

		// Parse the request body to get the updated account data
		var updatedData nessie.AccountData
		err := json.NewDecoder(r.Body).Decode(&updatedData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Use the NessieAPI to update the account
		err = api.UpdateAccount(accountID, updatedData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Send a success response
		w.WriteHeader(http.StatusNoContent)
	}
}

// DeleteAccountHandler handles DELETE requests to delete an account by ID.
func DeleteAccountHandler(api *nessie.NessieAPI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extract account ID from the request parameters
		vars := mux.Vars(r)
		accountID := vars["id"]

		// Use the NessieAPI to delete the account
		err := api.DeleteAccount(accountID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Send a success response
		w.WriteHeader(http.StatusNoContent)
	}
}

// Import necessary packages

/*// GetATM handles GET requests to retrieve an ATM by ID.
func GetATM(w http.ResponseWriter, r *http.Request) {
	// Extract ATM ID from the request parameters
	vars := mux.Vars(r)
	atmID := vars["id"]

	// Use the Nessie package to get the ATM details
	atm, err := nessieAPI.GetATM(atmID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Convert the ATM details to JSON and send the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(atm)
}

// GetAllATMs handles GET requests to retrieve all ATMs.
func GetAllATMs(w http.ResponseWriter, r *http.Request) {
	// Use the Nessie package to get all ATMs
	atms, err := nessieAPI.GetAllATMs()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Convert the ATM list to JSON and send the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(atms)
}

// CreateATM handles POST requests to create a new ATM.
func CreateATM(w http.ResponseWriter, r *http.Request) {
	// Extract ATM details from the request body
	var newATM nessieAPI.ATM
	err := json.NewDecoder(r.Body).Decode(&newATM)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Use the Nessie package to create the new ATM
	createdATM, err := nessieAPI.CreateATM(newATM)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Convert the created ATM details to JSON and send the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdATM)
}

// UpdateATM handles PUT requests to update an existing ATM.
func UpdateATM(w http.ResponseWriter, r *http.Request) {
	// Extract ATM ID from the request parameters
	vars := mux.Vars(r)
	atmID := vars["id"]

	// Extract updated ATM details from the request body
	var updatedATM nessieAPI.ATM
	err := json.NewDecoder(r.Body).Decode(&updatedATM)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Use the Nessie package to update the ATM
	err = nessieAPI.UpdateATM(atmID, updatedATM)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send a success response
	w.WriteHeader(http.StatusNoContent)
}

// DeleteATM handles DELETE requests to delete an ATM by ID.
func DeleteATM(w http.ResponseWriter, r *http.Request) {
	// Extract ATM ID from the request parameters
	vars := mux.Vars(r)
	atmID := vars["id"]

	// Use the Nessie package to delete the ATM
	err := nessieAPI.DeleteATM(atmID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send a success response
	w.WriteHeader(http.StatusNoContent)
}

// GetATMHandler handles requests to get ATM details by ID.
func GetATMHandler(w http.ResponseWriter, r *http.Request) {
	// Extract ATM ID from the request parameters
	vars := mux.Vars(r)
	atmID := vars["id"]

	// Use the NessieAPI to get ATM details
	atm, err := nessieAPI.GetATMByID(atmID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Convert ATM details to JSON and send the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(atm)
}

*/
// Similarly, implement handlers for Bill endpoints

// GetBillsByCustomerHandler handles requests to get bills associated with a specific customer.
// GetBillsByCustomerHandler retrieves bills associated with a specific customer ID.
func GetBillsByCustomerHandler(api *nessie.NessieAPI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extract customer ID from the request parameters
		vars := mux.Vars(r)
		customerID := vars["id"]

		// Use the NessieAPI to get bills for the customer
		bills, err := api.GetBillsByCustomer(customerID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Convert bills to JSON and send the response
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(bills)
	}
}

// GetBillByIDHandler handles requests to get a bill by ID.
func GetBillByIDHandler(api *nessie.NessieAPI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extract bill ID from the request parameters
		vars := mux.Vars(r)
		billID := vars["id"]

		// Use the NessieAPI to get the bill details
		bill, err := api.GetBillByID(billID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Convert bill details to JSON and send the response
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(bill)
	}
}

// CreateBillHandler handles requests to create a new bill.
func CreateBillHandler(api *nessie.NessieAPI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extract account ID from the request parameters
		vars := mux.Vars(r)
		accountID := vars["id"]

		// Decode the request body to get the bill data
		var billData nessie.BillData
		err := json.NewDecoder(r.Body).Decode(&billData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Use the NessieAPI to create a new bill
		createdBill, err := api.CreateBill(accountID, billData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Convert the created bill details to JSON and send the response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(createdBill)
	}
}

// UpdateBillHandler handles requests to update an existing bill.
func UpdateBillHandler(api *nessie.NessieAPI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extract bill ID from the request parameters
		vars := mux.Vars(r)
		billID := vars["id"]

		// Decode the request body to get the updated bill data
		var updatedData nessie.BillData
		err := json.NewDecoder(r.Body).Decode(&updatedData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Use the NessieAPI to update the bill
		err = api.UpdateBill(billID, updatedData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Send a success response
		w.WriteHeader(http.StatusNoContent)
	}
}

// DeleteBillHandler handles requests to delete a bill by ID.
func DeleteBillHandler(api *nessie.NessieAPI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extract bill ID from the request parameters
		vars := mux.Vars(r)
		billID := vars["id"]

		// Use the NessieAPI to delete the bill
		err := api.DeleteBill(billID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Send a success response
		w.WriteHeader(http.StatusNoContent)
	}
}

/*// Import necessary packages

// GetBranch handles GET requests to retrieve a branch by ID.
func GetBranch(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	branchID := vars["id"]

	branch, err := nessieAPI.GetBranch(branchID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(branch)
}

// GetAllBranches handles GET requests to retrieve all branches.
func GetAllBranches(w http.ResponseWriter, r *http.Request) {
	branches, err := nessieAPI.GetAllBranches()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(branches)
}

// CreateBranch handles POST requests to create a new branch.
func CreateBranch(w http.ResponseWriter, r *http.Request) {
	var newBranch nessieAPI.Branch
	err := json.NewDecoder(r.Body).Decode(&newBranch)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdBranch, err := nessieAPI.CreateBranch(newBranch)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdBranch)
}

// UpdateBranch handles PUT requests to update an existing branch.
func UpdateBranch(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	branchID := vars["id"]

	var updatedBranch nessieAPI.Branch
	err := json.NewDecoder(r.Body).Decode(&updatedBranch)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = nessieAPI.UpdateBranch(branchID, updatedBranch)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// DeleteBranch handles DELETE requests to delete a branch by ID.
func DeleteBranch(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	branchID := vars["id"]

	err := nessieAPI.DeleteBranch(branchID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// Similarly, implement handlers for Customer Data endpoints
*/
// handlers.go in the api package

// GetCustomerByIDHandler returns a handler function to handle requests to get a customer by ID.
func GetCustomerByIDHandler(nessieAPI *nessie.NessieAPI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extract customer ID from the request parameters
		vars := mux.Vars(r)
		customerID := vars["id"]

		// Use the NessieAPI to get customer details
		customer, err := nessieAPI.GetCustomerByID(customerID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Convert customer details to JSON and send the response
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customer)
	}
}

// GetAllCustomersHandler returns a handler function to handle requests to get all customers.
func GetAllCustomersHandler(nessieAPI *nessie.NessieAPI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Use the NessieAPI to get all customers
		customers, err := nessieAPI.GetAllCustomers()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Convert customers to JSON and send the response
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}

// CreateCustomerHandler returns a handler function to handle requests to create a new customer.
func CreateCustomerHandler(nessieAPI *nessie.NessieAPI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Decode the request body to get the customer data
		var newCustomer nessie.Customer
		err := json.NewDecoder(r.Body).Decode(&newCustomer)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Use the NessieAPI to create a new customer
		err = nessieAPI.CreateCustomer(newCustomer.FirstName, newCustomer.LastName,
			newCustomer.Address.StreetNumber, newCustomer.Address.StreetName, newCustomer.Address.City,
			newCustomer.Address.State, newCustomer.Address.Zip)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Send a success response
		w.WriteHeader(http.StatusCreated)
	}
}

// UpdateCustomerHandler returns a handler function to handle requests to update an existing customer.
func UpdateCustomerHandler(nessieAPI *nessie.NessieAPI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extract customer ID from the request parameters
		vars := mux.Vars(r)
		customerID := vars["id"]

		// Decode the request body to get the updated customer data
		var updatedCustomer nessie.Customer
		err := json.NewDecoder(r.Body).Decode(&updatedCustomer)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Use the NessieAPI to update the customer
		err = nessieAPI.UpdateCustomer(customerID, updatedCustomer.Address.StreetNumber,
			updatedCustomer.Address.StreetName, updatedCustomer.Address.City, updatedCustomer.Address.State,
			updatedCustomer.Address.Zip)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Send a success response
		w.WriteHeader(http.StatusNoContent)
	}
}

// DeleteDataHandler returns a handler function to handle the deletion of data associated with the API key and the specified type.
func DeleteDataHandler(nessieAPI *nessie.NessieAPI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extract data type from the request parameters
		vars := mux.Vars(r)
		dataType := vars["type"]

		// Use the NessieAPI to delete data
		err := nessieAPI.DeleteData(dataType)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Send success response
		w.WriteHeader(http.StatusNoContent)
	}
}

// GetAllDepositsHandler retrieves all deposits associated with the specified account ID.
func GetAllDepositsHandler(api *nessie.NessieAPI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extract account ID from the request parameters
		vars := mux.Vars(r)
		accountID := vars["id"]

		// Use the NessieAPI to get all deposits
		deposits, err := api.GetAllDeposits(accountID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Convert the deposits to JSON and send the response
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(deposits)
	}
}

// GetDepositByIDHandler retrieves a deposit by its ID.
func GetDepositByIDHandler(api *nessie.NessieAPI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extract deposit ID from the request parameters
		vars := mux.Vars(r)
		depositID := vars["id"]

		// Use the NessieAPI to get the deposit by ID
		deposit, err := api.GetDepositByID(depositID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Convert the deposit to JSON and send the response
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(deposit)
	}
}

// CreateDepositHandler creates a new deposit for the specified account.
func CreateDepositHandler(api *nessie.NessieAPI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extract account ID from the request parameters
		vars := mux.Vars(r)
		accountID := vars["id"]

		// Decode the request body into a Deposit struct
		var deposit nessie.Deposit
		err := json.NewDecoder(r.Body).Decode(&deposit)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Use the NessieAPI to create the deposit
		err = api.CreateDeposit(accountID, &deposit)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Respond with a success message
		w.WriteHeader(http.StatusCreated)
	}
}

// UpdateDepositHandler updates an existing deposit.
func UpdateDepositHandler(api *nessie.NessieAPI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extract deposit ID from the request parameters
		vars := mux.Vars(r)
		depositID := vars["id"]

		// Decode the request body into an updated Deposit struct
		var updatedDeposit nessie.Deposit
		err := json.NewDecoder(r.Body).Decode(&updatedDeposit)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Use the NessieAPI to update the deposit
		err = api.UpdateDeposit(depositID, &updatedDeposit)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Respond with a success message
		w.WriteHeader(http.StatusAccepted)
	}
}

// DeleteDepositHandler deletes an existing deposit.
func DeleteDepositHandler(api *nessie.NessieAPI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extract deposit ID from the request parameters
		vars := mux.Vars(r)
		depositID := vars["id"]

		// Use the NessieAPI to delete the deposit
		err := api.DeleteDeposit(depositID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Respond with a success message
		w.WriteHeader(http.StatusNoContent)
	}
}

// Import necessary packages

// GetAllLoansHandler retrieves all loans associated with the specified account ID.
func GetAllLoansHandler(w http.ResponseWriter, r *http.Request) {
	// Extract account ID from the request parameters
	vars := mux.Vars(r)
	accountID := vars["accountID"]

	// Use the NessieAPI to get all loans
	loans, err := nessieAPI.GetAllLoans(accountID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Convert loans to JSON and send the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(loans)
}

// GetLoanByIDHandler retrieves a loan by its ID.
func GetLoanByIDHandler(w http.ResponseWriter, r *http.Request) {
	// Extract loan ID from the request parameters
	vars := mux.Vars(r)
	loanID := vars["loanID"]

	// Use the NessieAPI to get the loan by ID
	loan, err := nessieAPI.GetLoanByID(loanID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Convert loan to JSON and send the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(loan)
}

// CreateLoanHandler creates a new loan for the specified account.
func CreateLoanHandler(w http.ResponseWriter, r *http.Request) {
	// Extract account ID from the request parameters
	vars := mux.Vars(r)
	accountID := vars["accountID"]

	// Decode the loan data from the request body
	var newLoan nessie.Loan
	err := json.NewDecoder(r.Body).Decode(&newLoan)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Use the NessieAPI to create a new loan
	createdLoan, err := nessieAPI.CreateLoan(accountID, &newLoan)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Convert created loan to JSON and send the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdLoan)
}

// UpdateLoanHandler updates an existing loan.
func UpdateLoanHandler(w http.ResponseWriter, r *http.Request) {
	// Extract loan ID from the request parameters
	vars := mux.Vars(r)
	loanID := vars["loanID"]

	// Decode the updated loan data from the request body
	var updatedLoan nessie.Loan
	err := json.NewDecoder(r.Body).Decode(&updatedLoan)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Use the NessieAPI to update the loan
	err = nessieAPI.UpdateLoan(loanID, &updatedLoan)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send a success response
	w.WriteHeader(http.StatusNoContent)
}

// DeleteLoanHandler deletes an existing loan.
func DeleteLoanHandler(w http.ResponseWriter, r *http.Request) {
	// Extract loan ID from the request parameters
	vars := mux.Vars(r)
	loanID := vars["loanID"]

	// Use the NessieAPI to delete the loan
	err := nessieAPI.DeleteLoan(loanID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send a success response
	w.WriteHeader(http.StatusNoContent)
}

/*// GetMerchant handles GET requests to retrieve a merchant by ID.
func GetMerchant(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	merchantID := vars["id"]

	merchant, err := nessieAPI.GetMerchant(merchantID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(merchant)
}

// GetAllMerchants handles GET requests to retrieve all merchants.
func GetAllMerchants(w http.ResponseWriter, r *http.Request) {
	merchants, err := nessieAPI.GetAllMerchants()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(merchants)
}

// CreateMerchant handles POST requests to create a new merchant.
func CreateMerchant(w http.ResponseWriter, r *http.Request) {
	var newMerchant nessieAPI.Merchant
	err := json.NewDecoder(r.Body).Decode(&newMerchant)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdMerchant, err := nessieAPI.CreateMerchant(newMerchant)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdMerchant)
}

// UpdateMerchant handles PUT requests to update an existing merchant.
func UpdateMerchant(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	merchantID := vars["id"]

	var updatedMerchant nessieAPI.Merchant
	err := json.NewDecoder(r.Body).Decode(&updatedMerchant)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = nessieAPI.UpdateMerchant(merchantID, updatedMerchant)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// DeleteMerchant handles DELETE requests to delete a merchant by ID.
func DeleteMerchant(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	merchantID := vars["id"]

	err := nessieAPI.DeleteMerchant(merchantID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
*/
/*// Import necessary packages

// GetPurchase handles GET requests to retrieve a purchase by ID.
func GetPurchase(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	purchaseID := vars["id"]

	purchase, err := nessieAPI.GetPurchase(purchaseID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(purchase)
}

// GetAllPurchases handles GET requests to retrieve all purchases.
func GetAllPurchases(w http.ResponseWriter, r *http.Request) {
	purchases, err := nessieAPI.GetAllPurchases()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(purchases)
}

// CreatePurchase handles POST requests to create a new purchase.
func CreatePurchase(w http.ResponseWriter, r *http.Request) {
	var newPurchase nessieAPI.Purchase
	err := json.NewDecoder(r.Body).Decode(&newPurchase)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdPurchase, err := nessieAPI.CreatePurchase(newPurchase)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdPurchase)
}

// UpdatePurchase handles PUT requests to update an existing purchase.
func UpdatePurchase(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	purchaseID := vars["id"]

	var updatedPurchase nessieAPI.Purchase
	err := json.NewDecoder(r.Body).Decode(&updatedPurchase)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = nessieAPI.UpdatePurchase(purchaseID, updatedPurchase)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// DeletePurchase handles DELETE requests to delete a purchase by ID.
func DeletePurchase(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	purchaseID := vars["id"]

	err := nessieAPI.DeletePurchase(purchaseID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
*/

// CreateTransferHandler creates a new transfer.
func CreateTransferHandler(api *nessie.NessieAPI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extract account ID from the request parameters
		vars := mux.Vars(r)
		accountID := vars["accountID"]

		// Decode the request body into a Transfer struct
		var newTransfer nessie.Transfer
		err := json.NewDecoder(r.Body).Decode(&newTransfer)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Use the NessieAPI to create the transfer
		createdTransfer, err := api.CreateTransfer(accountID, &newTransfer)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Convert the created transfer to JSON and send the response
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(createdTransfer)
	}
}

// UpdateTransferHandler updates an existing transfer.
func UpdateTransferHandler(api *nessie.NessieAPI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extract transfer ID from the request parameters
		vars := mux.Vars(r)
		transferID := vars["transferID"]

		// Decode the request body into a Transfer struct
		var updatedTransfer nessie.Transfer
		err := json.NewDecoder(r.Body).Decode(&updatedTransfer)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Use the NessieAPI to update the transfer
		err = api.UpdateTransfer(transferID, &updatedTransfer)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Send a success response
		w.WriteHeader(http.StatusAccepted)
	}
}

// DeleteTransferHandler deletes an existing transfer.
func DeleteTransferHandler(api *nessie.NessieAPI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extract transfer ID from the request parameters
		vars := mux.Vars(r)
		transferID := vars["transferID"]

		// Use the NessieAPI to delete the transfer
		err := api.DeleteTransfer(transferID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Send a success response
		w.WriteHeader(http.StatusNoContent)
	}
}

// GetAllWithdrawalsHandler retrieves all withdrawals associated with a specific account.
func GetAllWithdrawalsHandler(api *nessie.NessieAPI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extract account ID from the request parameters
		vars := mux.Vars(r)
		accountID := vars["accountID"]

		// Use the NessieAPI to get all withdrawals
		withdrawals, err := api.GetAllWithdrawals(accountID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Convert withdrawals to JSON and send the response
		responseJSON, err := json.Marshal(withdrawals)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(responseJSON)
	}
}

// GetWithdrawalByIDHandler retrieves a withdrawal by its ID.
func GetWithdrawalByIDHandler(api *nessie.NessieAPI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extract withdrawal ID from the request parameters
		vars := mux.Vars(r)
		withdrawalID := vars["withdrawalID"]

		// Use the NessieAPI to get the withdrawal by ID
		withdrawal, err := api.GetWithdrawalByID(withdrawalID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Convert withdrawal to JSON and send the response
		responseJSON, err := json.Marshal(withdrawal)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(responseJSON)
	}
}

// CreateWithdrawalHandler creates a new withdrawal.
func CreateWithdrawalHandler(api *nessie.NessieAPI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extract account ID from the request parameters
		vars := mux.Vars(r)
		accountID := vars["accountID"]

		// Parse the request body to get the withdrawal data
		var withdrawal nessie.Withdrawal
		err := json.NewDecoder(r.Body).Decode(&withdrawal)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Use the NessieAPI to create a new withdrawal
		createdWithdrawal, err := api.CreateWithdrawal(accountID, &withdrawal)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Convert created withdrawal to JSON and send the response
		responseJSON, err := json.Marshal(createdWithdrawal)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(responseJSON)
	}
}

// UpdateWithdrawalHandler updates an existing withdrawal.
func UpdateWithdrawalHandler(api *nessie.NessieAPI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extract withdrawal ID from the request parameters
		vars := mux.Vars(r)
		withdrawalID := vars["withdrawalID"]

		// Parse the request body to get the updated withdrawal data
		var updatedWithdrawal nessie.Withdrawal
		err := json.NewDecoder(r.Body).Decode(&updatedWithdrawal)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Use the NessieAPI to update the withdrawal
		err = api.UpdateWithdrawal(withdrawalID, &updatedWithdrawal)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Send a success response
		w.WriteHeader(http.StatusNoContent)
	}
}

// DeleteWithdrawalHandler deletes an existing withdrawal.
func DeleteWithdrawalHandler(api *nessie.NessieAPI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extract withdrawal ID from the request parameters
		vars := mux.Vars(r)
		withdrawalID := vars["withdrawalID"]

		// Use the NessieAPI to delete the withdrawal
		err := api.DeleteWithdrawal(withdrawalID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Send a success response
		w.WriteHeader(http.StatusNoContent)
	}
}
