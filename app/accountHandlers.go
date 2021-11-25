package app

import (
	"encoding/json"
	"fmt"
	"github.com/AliSayyah/banking/dto"
	"github.com/AliSayyah/banking/service"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type AccountHandler struct {
	service service.AccountService
}

func (h AccountHandler) NewAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerID := vars["id"]
	var request dto.AccountRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		customerID, _ := strconv.Atoi(customerID)
		request.CustomerID = customerID
		fmt.Println(request.CustomerID)
		account, appError := h.service.NewAccount(request)
		if appError != nil {
			writeResponse(w, appError.Code, appError.Message)
		} else {
			writeResponse(w, http.StatusCreated, account)
		}
	}
}
