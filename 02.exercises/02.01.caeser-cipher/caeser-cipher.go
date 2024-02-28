package main

import (
	"fmt"
	"strings"
)

func alphabetSearch(alphabet [26]string, letter string) int {
	// Linear search, checks all elements of the array
	for index, value := range alphabet {
		if value == letter {
			return index
		}
	}
	// Return -1 if the letter is not found
	// ... though this should not be encountered
	return -1
}

func encrypt(alphabet [26]string, message string, key int) string {
	var encryptedMessage string
	var encryptedLetter string
	var letterPosition int
	for i := range len(message) {
		// Find the position of the letter in the alphabet
		// and shift it right using the key
		letterPosition = alphabetSearch(alphabet, string(message[i]))
		encryptedLetter = alphabet[(letterPosition+key)%26]
		encryptedMessage = encryptedMessage + encryptedLetter
	}
	return encryptedMessage
}

func decrypt(alphabet [26]string, message string, key int) string {
	var decryptedMessage string
	var decryptedLetter string
	var letterPosition int
	for i := range len(message) {
		// Find the position of the letter in the alphabet
		// and shift it left using the key
		letterPosition = alphabetSearch(alphabet, string(message[i]))
		decryptedLetter = alphabet[(letterPosition-key)%26]
		decryptedMessage = decryptedMessage + decryptedLetter
	}
	return decryptedMessage
}

func main() {
	// Set-up an array of all the letters in the alphabet
	// for encryption and decryption
	alphabet := [26]string{
		"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M",
		"N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
	}
	var message string
	var key int

	// Indefinite loop repeats until exited
	for {
		// Encode/decode select
		var option int
		fmt.Print("\nPlease select an option:\n1. Encrypt\n2. Decrypt\n3. Exit\nSelection: ")
		fmt.Scan(&option)

		if option == 1 {
			// Take message to be encoded
			fmt.Print("\nEnter message to be encrypted: ")
			fmt.Scan(&message)
			message = strings.ToUpper(message)

			// Take key for encryption
			fmt.Print("Enter integer key to use for encryption: ")
			fmt.Scan(&key)

			// Return encrypted message
			fmt.Println("\nYour unencrypted message is", message, "and your key is", key)
			fmt.Println("Your encrypted message is", encrypt(alphabet, message, key))
		} else if option == 2 {
			// Take message
			fmt.Print("\nEnter message to be decrypted: ")
			fmt.Scan(&message)
			message = strings.ToUpper(message)

			// Take key for decryption
			fmt.Print("\nEnter integer key to use for decryption: ")
			fmt.Scan(&key)

			// Return decrypted message
			fmt.Println("\nYour encrypted message is", message, "and your key is", key)
			fmt.Println("Your decrypted message is", decrypt(alphabet, message, key))
		} else if option == 3 {
			break
		} else {
			fmt.Println("\nYour input was unrecognised")
		}
	}
}
