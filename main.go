package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		printHelp()
		os.Exit(0)
	}

	function := os.Args[1]

	switch function{
	case "help":
		printHelp()
	case "encrypt":
		handleEncryption()
	case "decrypt":
		handleDecryption()
	default:
		fmt.Println("Run encrypt to encrypt a file or decrypt to decrypt a file")
		os.Exit(1)
	}
	
}

func printHelp() {
	fmt.Println("File encryption")
	fmt.Println("Simple file encrypter for day to day use")
	fmt.Println("")
	fmt.Println("Useage:")
	fmt.Println("")
	fmt.Println("\tgo run . encrypt /path/to/file")
	fmt.Println("")
	fmt.Println("Commands:")
	fmt.Println("")
	fmt.Println("\t encrypt\tEncrypts a file when given a password")
	fmt.Println("\t decrypt\tDecrypts a file using a given password")
	fmt.Println("\t help\tDisplay help text")
	fmt.Println("")
}

func handleEncryption(){
	if len(os.Args) < 3 {
		println("missing a required argument. For more info, run go run . help")
		os.Exit(0)
	}

	file := os.Args[2]
	if !validateFile(file){
		panic("File not found")
	}

	password := getPassword()()
}

func handleDecryption(){}

func getPassword() []byte{
	fmt.Print("Enter password:")
	password, _ := term.ReadPassword(0)
	fmt.Print("Confirm password:")
	confimrPassword, _ := term.ReadPassword(0)

	if !validatePassword(password, confimrPassword){
		fmt.Println("\nPasswords do not match. Please try again")
		return getPassword()
	}

	return password
}

func validateFile(file string) bool {
	if _, err := os.Stat(file); 
	os.IsNotExist(err){return false}
	return true
}

func validatePassword(password []byte, confimrPassword []byte) bool{
	if !bytes.Equal(password, confimrPassword){
		return false
	}

	return true
}