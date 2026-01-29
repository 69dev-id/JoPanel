package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	bytes, _ := bcrypt.GenerateFromPassword([]byte("password123"), 14)
	fmt.Printf("%s", string(bytes))
}
