package service

import "net/http"

// Route define a structure to create routes
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes defines the type Routes which is just an array (slice) of Route structs.
type Routes []Route

// Initialize our routes
var routes Routes = Routes{
	Route{
		"GetAccount",
		"GET",
		"/accounts/{accountId}",
		GetAccountByID,
	},
}
