package enigma

import (
	"fmt"
	"log"
	"strings"
)

type Rotor struct {
	reflector bool
	shift int
}

func CreateRotor(shift int, reflector bool) *Rotor {
	r := Rotor{reflector: reflector}
	r.shift = shift
	return &r
}

func (r *Rotor) Encrypt(message string) string {
	// Create string builder to create encrypted string
	var sb strings.Builder

	for _, char := range message {

		char_index := strings.IndexRune(characters, char)
		if char_index == -1 {
			log.Fatal("Encryption failed due to invalid character")
		}
		
		// Calculate new character index from [characters]
		// and add to the string
		new_index := mod(char_index + r.shift, len(characters))
		sb.WriteByte(characters[new_index])
	}

	return sb.String()
}

func (r *Rotor) Decrypt(message string) string {
	var sb strings.Builder

	for _, char := range message {

		char_index := strings.IndexRune(characters, char)
		if char_index == -1 {
			log.Fatal("Decryption failed due to invalid character")
		}
		
		// Calculate new character index from [characters]
		// and add to the string
		new_index := mod(char_index - r.shift, len(characters))
		sb.WriteByte(characters[new_index])
	}

	return sb.String()
}

/*
Utilities
*/

// Mod similar to the one in python
func mod(a, b int) int {
	return (a % b + b) % b
}

func (r *Rotor) DisplayRotor() {
	fmt.Println("--------------------------------")
	fmt.Println("Current contents of Rotor :")
	fmt.Println(" - Shift is: ", r.shift)
	fmt.Printf(" - Reflector is: %t \n", r.reflector)
	fmt.Println("--------------------------------")
}

func (r *Rotor) SetShift(new_shift int) {
	r.shift = new_shift
}

func (r *Rotor) IncrementShift() {
	r.shift = r.shift + 1
}
