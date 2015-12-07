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

		if len(auth) != 2 || auth[0] != "Basic" {
			http.Error(w, "bad syntax", http.StatusBadRequest)
			return
		}

		payload, _ := base64.StdEncoding.DecodeString(auth[1])
		pair := strings.SplitN(string(payload), ":", 2)

		if len(pair) != 2 || !Validate(pair[0], pair[1]) {
			http.Error(w, "authorization failed", http.StatusUnauthorized)
			return
		}
		fmt.Fprintf(w, "Before\n")
		pass(w, r)
		fmt.Fprintf(w, "After\n")
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

		if len(auth) != 2 || auth[0] != "Bearer" {
			http.Error(w, "bad syntax", http.StatusBadRequest)
			return
		}

		payload, _ := base64.StdEncoding.DecodeString(auth[1])

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

	// Check signature of
	return true
}

func generateToken(w http.ResponseWriter, r *http.Request) {

	mySigningKey := []byte(os.Getenv("SuperSecretKey"))

	if mySigningKey == nil {
		mySigningKey = []byte("SecretLike")
	}

	r.ParseForm()

	// logic part of log in
	username := r.Form["username"]
	password := r.Form["password"]
	fmt.Fprintf(w, "%s:%s\n", username, password)

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
