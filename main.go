package main

import (
	"fmt"
	"log"
	"os"
	"github.com/dagopherboy/jwt-go"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/authentication", authenticationHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func authenticationHandler(w http.ResponseWriter, r *http.Request) {
	
	mySigningKey := []byte(os.Getenv("SuperSecretKey"))
	
	if  mySigningKey == nil {
		mySigningKey = []byte("SecretLike")
	}
	
	token := jwt.New(jwt.SigningMethodHS256)
	// Set some claims
	token.Claims["foo"] = "bar"
	token.Claims["nbf"] = time.Now()
	token.Claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		fmt.Fprintf(w, "%s\n", err)		
		return
	} 	
	fmt.Fprintf(w, "%s\n", tokenString)		
}