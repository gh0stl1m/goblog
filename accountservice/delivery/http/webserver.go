package delivery

import (
	"log"
	"net/http"
)

// StartWebServer initialize the http server
func StartWebServer(port string) {
	// Initialize router
	r := NewRouter()
	http.Handle("/", r)

	log.Println("Starting HTTP server at port " + port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Println("An error occured starting HTTP server at port " + port)
		log.Println("Error: " + err.Error())
	}
}
