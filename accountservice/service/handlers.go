package service

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gh0stl1m/goblog/accountservice/dbclient"
	"github.com/gorilla/mux"
)

// DBClient contains the connection to the DB
var DBClient dbclient.BoltClient

// GetAccountByID Retrieves an account based on the ID
func GetAccountByID(w http.ResponseWriter, r *http.Request) {
	var accountID = mux.Vars(r)["accountId"]

	account, err := DBClient.QueryAccount(accountID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	data, _ := json.Marshal(account)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
