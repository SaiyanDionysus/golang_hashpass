package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func HashPass(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPassHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func main() {
	password := "secret"
	hash, _ := HashPass(password)

	fmt.Println("Password: ", password)
	fmt.Println("Hash:     ", hash)

	match := CheckPassHash(password, hash)
	fmt.Println("Match: ", match)
}
