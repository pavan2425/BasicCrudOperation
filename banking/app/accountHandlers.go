package app

import (
	"banking/dto"
	"banking/service"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type accountHandler struct {
	service service.AccountService
}

func (ah *accountHandler) NewAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId := vars["customer_id"]

	var request dto.NewAccountRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeMessage(w, http.StatusBadRequest, err.Error())
	} else {
		request.CustomerID = customerId
		account, appError := ah.service.NewAccount(request)
		if appError != nil {
			writeMessage(w, appError.Code, appError.Message)
		} else {
			writeMessage(w, http.StatusCreated, account)

		}
	}
}
