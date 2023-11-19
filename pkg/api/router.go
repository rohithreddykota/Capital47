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

	router.HandleFunc("/customers/{id}/bills", GetBillsByCustomerHandler(api)).Methods("GET")
	router.HandleFunc("/bills/{id}", GetBillByIDHandler(api)).Methods("GET")
	router.HandleFunc("/accounts/{id}/bills", CreateBillHandler(api)).Methods("POST")
	router.HandleFunc("/bills/{id}", UpdateBillHandler(api)).Methods("PUT")
	router.HandleFunc("/bills/{id}", DeleteBillHandler(api)).Methods("DELETE")

	router.HandleFunc("/customers/{id}", GetCustomerByIDHandler(api)).Methods("GET")
	router.HandleFunc("/customers", GetAllCustomersHandler(api)).Methods("GET")
	router.HandleFunc("/customers", CreateCustomerHandler(api)).Methods("POST")
	router.HandleFunc("/customers/{id}", UpdateCustomerHandler(api)).Methods("PUT")

	router.HandleFunc("/data/{type}", DeleteDataHandler(api)).Methods("DELETE")

	router.HandleFunc("/deposits/{id}", GetAllDepositsHandler(api)).Methods("GET")
	router.HandleFunc("/deposits/{id}", GetDepositByIDHandler(api)).Methods("GET")
	router.HandleFunc("/deposits/{id}", CreateDepositHandler(api)).Methods("POST")
	router.HandleFunc("/deposits/{id}", UpdateDepositHandler(api)).Methods("PUT")
	router.HandleFunc("/deposits/{id}", DeleteDepositHandler(api)).Methods("DELETE")

	router.HandleFunc("/accounts/{id}/transfers", GetAllTransfersHandler(api)).Methods("GET")
	router.HandleFunc("/transfers/{id}", GetTransferByIDHandler(api)).Methods("GET")
	router.HandleFunc("/accounts/{id}/transfers", CreateTransferHandler(api)).Methods("POST")
	router.HandleFunc("/transfers/{id}", UpdateTransferHandler(api)).Methods("PUT")
	router.HandleFunc("/transfers/{id}", DeleteTransferHandler(api)).Methods("DELETE")

	router.HandleFunc("/withdrawals/{id}", GetAllWithdrawalsHandler(api)).Methods("GET")
	router.HandleFunc("/withdrawals/{id}", GetWithdrawalByIDHandler(api)).Methods("GET")
	router.HandleFunc("/withdrawals/{id}", CreateWithdrawalHandler(api)).Methods("POST")
	router.HandleFunc("/withdrawals/{id}", UpdateWithdrawalHandler(api)).Methods("PUT")
	router.HandleFunc("/withdrawals/{id}", DeleteWithdrawalHandler(api)).Methods("DELETE")

	return router
}
