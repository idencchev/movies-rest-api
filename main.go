package main

import (
	"fmt"
	"log"
	"movies-rest-api/router"
	"net/http"
)

func main() {
	r := router.Router()
	fmt.Println("Starting server on the port 8000...")
	log.Fatal(http.ListenAndServe(":8000", r))
}
