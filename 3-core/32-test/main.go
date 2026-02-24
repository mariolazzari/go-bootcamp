package main

import (
	"fmt"
	"regexp"
)

var pl = fmt.Println

func main() {
	// ----- AUTOMATED TESTING -----
	// Automated tests make sure your program still
	// works while you change the code
	// Create app2 directory with testemail.go
	// cd app2
	// Create testemail_test.go
	// go mod init app2
	// Run Tests : go test -v
}

func IsEmail(s string) (string, error) {
	// Used a raw string here so I didn't have
	// to double backslashes
	r, _ := regexp.Compile(`[\w._%+-]{1,20}@[\w.-]{2,20}.[A-Za-z]{2,3}`)

	if r.MatchString(s) {
		return "Valid Email", nil
	} else {
		return "", fmt.Errorf("not a valid email")
	}
}
