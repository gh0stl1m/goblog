package repository

import (
	"encoding/json"
	"fmt"

	"github.com/boltdb/bolt"
	"github.com/gh0stl1m/goblog/accountservice/domains"
)

// GetByID allow to get data form the DB
func (bc *BoltClient) GetByID(accountID string) (domains.Account, error) {
	account := domains.Account{}

	err := bc.boltConn.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(AccountBucketName))

		accountBytes := b.Get([]byte(accountID))
		if accountBytes == nil {
			return fmt.Errorf("No account found for " + accountID)
		}

		json.Unmarshal(accountBytes, &account)

		return nil
	})

	if err != nil {
		return domains.Account{}, err
	}

	return account, nil
}
