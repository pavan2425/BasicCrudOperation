package app

import (
	"banking/service"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// func greet(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "hello world")
// }

// func GetCustomer(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	fmt.Fprint(w, vars["customer_id"])
// }

// func createCustomer(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "customercreated")
// }

type customerHandler struct {
	service service.CustomerService
}

func (ch *customerHandler) GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")
	customers, err := ch.service.GetAllCustomer(status)
	if err != nil {
		writeMessage(w, err.Code, err.AsMessage())
	} else {
		writeMessage(w, http.StatusOK, customers)
	}
}

func (ch *customerHandler) GetCustomer(w http.ResponseWriter, r *http.Request) {
	// customers := []Customer{
	// 	{Name: "ashish", City: "mys", Zipcode: "55555"},
	// 	{Name: "shish", City: "mysa", Zipcode: "55554"},
	// }
	vars := mux.Vars(r)
	id := vars["customer_id"]
	customer, err := ch.service.GetCustomerById(id)
	if err != nil {
		fmt.Println(err.AsMessage())
		writeMessage(w, err.Code, err.AsMessage())
	} else {
		fmt.Println(err.AsMessage())
		writeMessage(w, http.StatusOK, customer)
	}

}

func writeMessage(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(data)
}
