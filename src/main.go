package main

import (
	"fmt"

	"example.com/crypto/enigma"
)
func main() {
	var to_encrypt string = "abcdefg"
	p := enigma.CreatePlugboard("")
	encrypted_message := p.Encrypt(to_encrypt)
	decrypted_message := p.Decrypt(encrypted_message)
	fmt.Println(decrypted_message)
}
