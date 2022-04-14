package main

import (
	"log"
	"net/http"

	"github.com/oranges0da/go-http/routes"
)

func handleRequests() {
	http.HandleFunc("/", routes.Home)
	log.Fatal(http.ListenAndServe(":4000", nil))
}

func main() {
	handleRequests()
}
