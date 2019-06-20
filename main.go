package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/authentication", BasicAuth(generateToken))
	http.HandleFunc("/hello", TokenAuth(resource))
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func resource(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!\n")
}

