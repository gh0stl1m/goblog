package main

import (
	"fmt"

	delivery "github.com/gh0stl1m/goblog/accountservice/delivery/http"
	repository "github.com/gh0stl1m/goblog/accountservice/repository/boltdb"
	"github.com/gh0stl1m/goblog/accountservice/usecases"
)

var appName string = "accountservice"

func main() {
	fmt.Printf("Starting %v\n", appName)
	initializeBoltClient()
	delivery.StartWebServer("6767")

}

func initializeBoltClient() {
	usecases.DBClient = repository.BoltClient{}
	usecases.DBClient.OpenConn()
	usecases.DBClient.Seed()
}
