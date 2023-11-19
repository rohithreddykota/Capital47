package api

import (
	"Capital47/pkg/nessie"
	"github.com/gorilla/mux"
)

// NewRouter creates a new router for handling API requests.
func NewRouter(api *nessie.NessieAPI) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	
	router.HandleFunc("/accounts/{id}", GetAccountHandler(api)).Methods("GET")
	router.HandleFunc("/customers/{customerID}/accounts", GetAccountsByCustomerHandler(api)).Methods("GET")
	router.HandleFunc("/customers/{customerID}/accounts", CreateAccountHandler(api)).Methods("POST")
	router.HandleFunc("/accounts/{id}", UpdateAccountHandler(api)).Methods("PUT")
	router.HandleFunc("/accounts/{id}", DeleteAccountHandler(api)).Methods("DELETE")

	router.HandleFunc("/customers/{id}/bills", GetBillsByCustomerHandler(nessieAPI)).Methods("GET")
	router.HandleFunc("/bills/{id}", GetBillByIDHandler(nessieAPI)).Methods("GET")
	router.HandleFunc("/accounts/{id}/bills", CreateBillHandler(nessieAPI)).Methods("POST")
	router.HandleFunc("/bills/{id}", UpdateBillHandler(nessieAPI)).Methods("PUT")
	router.HandleFunc("/bills/{id}", DeleteBillHandler(nessieAPI)).Methods("DELETE")

	router.HandleFunc("/customers/{id}", GetCustomerByIDHandler(nessieAPI)).Methods("GET")
	router.HandleFunc("/customers", GetAllCustomersHandler(nessieAPI)).Methods("GET")
	router.HandleFunc("/customers", CreateCustomerHandler(nessieAPI)).Methods("POST")
	router.HandleFunc("/customers/{id}", UpdateCustomerHandler(nessieAPI)).Methods("PUT")

	router.HandleFunc("/data/{type}", DeleteDataHandler(nessieAPI)).Methods("DELETE")

	router.HandleFunc("/deposits/{id}", GetAllDepositsHandler(nessieAPI)).Methods("GET")
	router.HandleFunc("/deposits/{id}", GetDepositByIDHandler(nessieAPI)).Methods("GET")
	router.HandleFunc("/deposits/{id}", CreateDepositHandler(nessieAPI)).Methods("POST")
	router.HandleFunc("/deposits/{id}", UpdateDepositHandler(nessieAPI)).Methods("PUT")
	router.HandleFunc("/deposits/{id}", DeleteDepositHandler(nessieAPI)).Methods("DELETE")

	router.HandleFunc("/transfers/{accountID}", CreateTransferHandler(api)).Methods("POST")
	router.HandleFunc("/transfers/{transferID}", UpdateTransferHandler(api)).Methods("PUT")
	router.HandleFunc("/transfers/{transferID}", DeleteTransferHandler(api)).Methods("DELETE")

	router.HandleFunc("/withdrawals/{accountID}", GetAllWithdrawalsHandler(api)).Methods("GET")
	router.HandleFunc("/withdrawals/{withdrawalID}", GetWithdrawalByIDHandler(api)).Methods("GET")
	router.HandleFunc("/withdrawals/{accountID}", CreateWithdrawalHandler(api)).Methods("POST")
	router.HandleFunc("/withdrawals/{withdrawalID}", UpdateWithdrawalHandler(api)).Methods("PUT")
	router.HandleFunc("/withdrawals/{withdrawalID}", DeleteWithdrawalHandler(api)).Methods("DELETE")

	return router
}
