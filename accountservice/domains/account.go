package domains

import "net/http"

// Account is the base model
type Account struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// AccountUseCases it's the contract
type AccountUseCases interface {
	GetByID(accountID string) (Account, error)
}

// AccountAdapter it's the contract for the httpService
type AccountAdapter interface {
	GetByID(w http.ResponseWriter, r *http.Request)
}
