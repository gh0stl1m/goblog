package repository

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/boltdb/bolt"
	"github.com/gh0stl1m/goblog/accountservice/domains"
)

// AccountBucketName is the bucket name where the data will be stored
const AccountBucketName string = "AccountBucket"

// BoltClient stores a pointer of the db connection
type BoltClient struct {
	boltConn *bolt.DB
}

// IBoltClient expose the contract of the Bolt DB
type IBoltClient interface {
	OpenConn()
	Seed()
}

// OpenConn open the connection to the DB
func (bc *BoltClient) OpenConn() {
	var err error
	bc.boltConn, err = bolt.Open("accounts.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
}

// Seed start seeding accounts
func (bc *BoltClient) Seed() {
	bc.initializeBucket()
	bc.seedAccounts()
}

// Creates an "AccountBucket" in the BoltDB. It will overwrite any existing bucket of the same name.
func (bc *BoltClient) initializeBucket() {
	bc.boltConn.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte(AccountBucketName))
		if err != nil {
			return fmt.Errorf("create bucket failed: %s", err)
		}
		return nil
	})
}

// Seed (n) make-believe account objects into the AcountBucket bucket.
func (bc *BoltClient) seedAccounts() {
	total := 5
	for i := 0; i < total; i++ {

		// Generate a key 10000 or larger
		key := strconv.Itoa(10000 + i)

		// Create an instance of our Account struct
		acc := domains.Account{
			ID:   key,
			Name: "Person_" + strconv.Itoa(i),
		}

		// Serialize the struct to JSON
		jsonBytes, _ := json.Marshal(acc)

		// Write the data to the AccountBucket
		bc.boltConn.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte(AccountBucketName))
			err := b.Put([]byte(key), jsonBytes)
			return err
		})
	}
	fmt.Printf("Seeded %v fake accounts...\n", total)
}
