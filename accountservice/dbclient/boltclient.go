package dbclient

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/boltdb/bolt"
	"github.com/gh0stl1m/goblog/accountservice/model"
)

const ACOUNT_BUCKET_NAME string = "AccountBucket"

// IBoltClient expose the contract of the Bolt DB
type IBoltClient interface {
	OpenBoltDB()
	QueryAccount(accountID string) (model.Account, error)
	Seed()
}

// BoltClient stores a pointer of the db connection
type BoltClient struct {
	boltDB *bolt.DB
}

// OpenBoltDB open the connection to the DB
func (bc *BoltClient) OpenBoltDB() {
	var err error
	bc.boltDB, err = bolt.Open("accounts.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
}

// Seed start seeding accounts
func (bc *BoltClient) Seed() {
	bc.initializeBucket()
	bc.seedAccounts()
}

// QueryAccount allow to get data form the DB
func (bc *BoltClient) QueryAccount(accountID string) (model.Account, error) {
	account := model.Account{}

	err := bc.boltDB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(ACOUNT_BUCKET_NAME))

		accountBytes := b.Get([]byte(accountID))
		if accountBytes == nil {
			return fmt.Errorf("No account found for " + accountID)
		}

		json.Unmarshal(accountBytes, &account)

		return nil
	})

	if err != nil {
		return model.Account{}, err
	}

	return account, nil
}

// Creates an "AccountBucket" in the BoltDB. It will overwrite any existing bucket of the same name.
func (bc *BoltClient) initializeBucket() {
	bc.boltDB.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte(ACOUNT_BUCKET_NAME))
		if err != nil {
			return fmt.Errorf("create bucket failed: %s", err)
		}
		return nil
	})
}

// Seed (n) make-believe account objects into the AcountBucket bucket.
func (bc *BoltClient) seedAccounts() {

	total := 100
	for i := 0; i < total; i++ {

		// Generate a key 10000 or larger
		key := strconv.Itoa(10000 + i)

		// Create an instance of our Account struct
		acc := model.Account{
			ID:   key,
			Name: "Person_" + strconv.Itoa(i),
		}

		// Serialize the struct to JSON
		jsonBytes, _ := json.Marshal(acc)

		// Write the data to the AccountBucket
		bc.boltDB.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte(ACOUNT_BUCKET_NAME))
			err := b.Put([]byte(key), jsonBytes)
			return err
		})
	}
	fmt.Printf("Seeded %v fake accounts...\n", total)
}
