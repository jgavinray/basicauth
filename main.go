package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/authentication", BasicAuth(generateToken))
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
