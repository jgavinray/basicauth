package main

import (
    "encoding/base64"
    "fmt"
    "net/http"
    "strings"    
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