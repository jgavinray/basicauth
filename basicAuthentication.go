package main

import (
	"encoding/base64"
	"fmt"
	"github.com/dagopherboy/jwt-go"
	"net/http"
	"os"
	"strings"
	"time"
)

func BasicAuth(pass http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		auth := strings.SplitN(r.Header["Authorization"][0], " ", 2)

		// Check to see if the Authorization Header has two parts.  A "Basic" followed by something
		if len(auth) != 2 || auth[0] != "Basic" {
			http.Error(w, "bad syntax", http.StatusBadRequest)
			return
		}

		// Get the second element of the auth array, and attempt to base64 decode it
		payload, _ := base64.StdEncoding.DecodeString(auth[1])
		// Split the username:password string into two array elements
		pair := strings.SplitN(string(payload), ":", 2)

		// Check to make sure we have two elements in the pair array, or that we can confirm the credentials are valid
		if len(pair) != 2 || !Validate(pair[0], pair[1]) {
			http.Error(w, "authorization failed", http.StatusUnauthorized)
			return
		}

		// If both of the above checks pass, forward the original request to the appropriate function.
		pass(w, r)

	}
}

func Validate(username, password string) bool {
	// Make come database call to validate username and password
	if username == "foo" && password == "bar" {
		return true
	}
	return false
}

func TokenAuth(pass http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		auth := strings.SplitN(r.Header["Authorization"][0], " ", 2)

		// Check to see if the Authorization Header has two parts.  A "Bearer" followed by something
		if len(auth) != 2 || auth[0] != "Bearer" {
			http.Error(w, "bad syntax", http.StatusBadRequest)
			return
		}

		// Get the second element of the auth array, and attempt to base64 decode it
		payload, _ := base64.StdEncoding.DecodeString(auth[1])

		// Attempt to Validate what was sent in the second array element of the Authorization header
		if !ValidateToken(payload) {
			http.Error(w, "authorization failed", http.StatusUnauthorized)
			return
		}
		fmt.Fprintf(w, "Before\n")
		pass(w, r)
		fmt.Fprintf(w, "After\n")
	}
}

func ValidateToken(tkn []byte) bool {

	tokenToBeValidated := string(tkn)
	token, err := jwt.Parse(tokenToBeValidated, func(token *jwt.Token) (interface{}, error) {

		// Check What Algorithm Signed the Token
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// Get the key used to sign the token in generateToken
		return GetSigningKey(), nil
	})

	if err == nil && token.Valid {
		return true
	} else {
		return false
	}
}

func GetSigningKey() []byte {
	// Read File, Database Call, Environment Variable, whereever we want to store the key
	mySigningKey := []byte(os.Getenv("SuperSecretKey"))

	if mySigningKey == nil {
		mySigningKey = []byte("SecretLike")
	}

	return mySigningKey
}

func generateToken(w http.ResponseWriter, r *http.Request) {

	mySigningKey := GetSigningKey()

	token := jwt.New(jwt.SigningMethodHS512)
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
