package main

import (
	"fmt"

	"github.com/gh0stl1m/goblog/accountservice/dbclient"
	"github.com/gh0stl1m/goblog/accountservice/service"
)

var appName string = "accountservice"

func main() {
	fmt.Printf("Starting %v\n", appName)
	initializeBoltClient()
	service.StartWebServer("6767")

}

func initializeBoltClient() {
	service.DBClient = dbclient.BoltClient{}
	service.DBClient.OpenBoltDB()
	service.DBClient.Seed()
}
