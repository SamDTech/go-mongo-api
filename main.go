package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/samdtech/go-mongo-api/router"
)

func main(){
	router := router.Router()

	fmt.Println("Netflix API")
	fmt.Println("Connected to MongoDB!")
	fmt.Println("Listening on port 4000...")

	log.Fatal(http.ListenAndServe(":4000", router))
}