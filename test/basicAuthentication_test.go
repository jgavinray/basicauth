package main

import (
	"testing"
)

func TestValidateUser(t *testing.T) {
	var tests = []struct {
		username string
		password string
		want     bool
	}{
		{"foo", "bar", true},
		{"user", "pass", false},
	}

	for _, test := range tests {
		if got := IsValidateUser(test.username, test.password); got != test.want {
			t.Error("ValidateUser(%q, %q) = %v", test.username, test.password, got)
		}
	}

}
