package usecases

import (
	"fmt"

	"github.com/gh0stl1m/goblog/accountservice/domains"
	repository "github.com/gh0stl1m/goblog/accountservice/repository/boltdb"
)

// DBClient contains the connection to the DB
var DBClient repository.BoltClient

// GetByID Retrieves an account based on the ID
func GetByID(accountID string) (domains.Account, error) {
	if accountID == "" {
		fmt.Errorf("Account ID can not be empty")
	}

	return DBClient.GetByID(accountID)
}
