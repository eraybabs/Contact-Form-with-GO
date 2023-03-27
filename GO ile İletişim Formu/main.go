package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/submit", submitHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func submitHandler(w http.ResponseWriter, r *http.Request) {

}
