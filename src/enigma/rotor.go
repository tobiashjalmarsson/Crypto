package enigma

import (
	"fmt"
	"log"
	"strings"
)

type Rotor struct {
	unencrypted string
	encrypted string
	decrypted string
	shift int
}

const characters string = "abcdefghijklmnopqrstvwyuzx "

func CreateRotor(to_encrypt string, shift int) Rotor {
	var r Rotor;
	r.unencrypted = strings.ToLower(to_encrypt)
	r.shift = shift
	return r
}

func (r Rotor) Encrypt() {
	fmt.Println("Starting rotor")

	// Create string builder to create encrypted string
	var sb strings.Builder

	for _, char := range r.unencrypted {

		char_index := strings.IndexRune(characters, char)
		if char_index == -1 {
			log.Fatal("Encryption failed due to invalid character")
		}
		
		// Calculate new character index from [characters]
		// and add to the string
		new_index := mod(char_index + (r.shift + 1), len(characters))
		sb.WriteByte(characters[new_index])
	}

	r.encrypted = sb.String()
	fmt.Println("New string is: ", r.encrypted)
}

func (r Rotor) Decrypt() {
	var sb strings.Builder

	fmt.Println("Size of string is: ", len(characters))

	for _, char := range r.encrypted {

		char_index := strings.IndexRune(characters, char)
		if char_index == -1 {
			log.Fatal("Decryption failed due to invalid character")
		}
		
		// Calculate new character index from [characters]
		// and add to the string
		new_index := mod(char_index - r.shift, len(characters))
		fmt.Printf("char_inde: %d, new_index: %d \n", char_index, new_index)
		sb.WriteByte(characters[new_index])
	}

	decrypted_string := sb.String()

	fmt.Printf("Decryption finished, message is: %s \n", decrypted_string)
	r.decrypted = decrypted_string

	//if decrypted_string != r.unencrypted {
	//	log.Fatal("Decrypted message does not match original message")
	//}
}

/*
Utilities
*/

// Mod similar to the one in python
func mod(a, b int) int {
	return (a % b + b) % b
}

func (r Rotor) DisplayRotor() {
	fmt.Printf("Unencrypted string is: %s \n", r.unencrypted)
	fmt.Printf("Encrypted string is: %s \n", r.encrypted)
	fmt.Printf("Decrypted string is: %s \n", r.decrypted)
	fmt.Println("Shift is: ", r.shift)
}

func (r Rotor) SetShift(new_shift int) {
	r.shift = new_shift
}

func (r Rotor) IncrementShift() {
	r.shift = r.shift + 1
}
