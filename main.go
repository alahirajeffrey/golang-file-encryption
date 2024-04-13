package main

import (
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

func printHelp() {}

func handleEncryption(){}

func handleDecryption(){}

func getPassword(){}

func validateFile(){}

func validatePassword(){}