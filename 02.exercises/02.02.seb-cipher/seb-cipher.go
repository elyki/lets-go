package main

import (
	"fmt"
	"math/rand/v2"
	"strings"
)

// Provisionally named the Seb Cipher ...
// I do not know if this already exists

func randomLetter(alphabet [26]string, numLetters int) string {
	var output string
	// Return one or more randomly selected letters
	// used for padding encrypted messages
	for i := range numLetters {
		output = output + alphabet[rand.IntN(26)]
		i++
	}
	return output
}

func encrypt(alphabet [26]string, message string, key int) string {
	// ascending indicates whether more or less
	// random characters are being added
	var ascending bool = false
	var numRandom int = key
	var encryptedMessage string
	// Iterate through every character in a message
	for i := range len(message) {
		if ascending {
			encryptedMessage = encryptedMessage + string(message[i]) + randomLetter(alphabet, numRandom)
			numRandom++
			// If the number of random characters is equal to the key,
			// the next iteration will start adding less random characters
			if numRandom == key {
				ascending = false
			}
		} else {
			encryptedMessage = encryptedMessage + string(message[i]) + randomLetter(alphabet, numRandom)
			numRandom--
			// If the number of random characters is equal to 0,
			// the next iteration will start adding more random characters
			if numRandom == 0 {
				ascending = true
			}
		}
	}
	return encryptedMessage
}

func decrypt(message string, key int) string {
	var ascending bool = false
	var numRandom int = key
	var decryptedMessage string
	// Iterate through every character in a message
	i := 0
	for i < len(message) {
		if ascending {
			// Take the current letter
			decryptedMessage = decryptedMessage + string(message[i])
			// Skip over the subsequent random characters
			i = i + numRandom + 1
			numRandom++
			if numRandom == key {
				ascending = false
			}
		} else {
			decryptedMessage = decryptedMessage + string(message[i])
			i = i + numRandom + 1
			numRandom--
			if numRandom == 0 {
				ascending = true
			}
		}
	}
	return decryptedMessage
}

func main() {
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
			fmt.Print("Enter integer key to use for decryption: ")
			fmt.Scan(&key)

			// Return decrypted message
			fmt.Println("\nYour encrypted message is", message, "and your key is", key)
			fmt.Println("Your decrypted message is", decrypt(message, key))
		} else if option == 3 {
			break
		} else {
			fmt.Println("\nYour input was unrecognised")
		}
	}
}
