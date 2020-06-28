package delivery

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gh0stl1m/goblog/accountservice/usecases"
	"github.com/gorilla/mux"
)

// GetByID is the handler for get account by id
func GetByID(w http.ResponseWriter, r *http.Request) {
	var accountID string = mux.Vars(r)["accountId"]

	account, err := usecases.GetByID(accountID)

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
