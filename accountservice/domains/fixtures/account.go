package fixtures

import (
	"math/rand"

	"github.com/gh0stl1m/goblog/accountservice/domains"
)

// GenerateID create random ids
func GenerateID() int {
	return rand.Intn(10000)
}

// GenerateAccount creates accounts
func GenerateAccount(accountID string) domains.Account {
	return domains.Account{
		ID:   accountID,
		Name: "Test_" + accountID,
	}
}
