package main

import (
	"log"
	"net/http"

	"github.com/webservice-golang-api/routes"
)

func main() {
	routes.Routes()
	log.Fatal(http.ListenAndServe(":8001", nil))
}
