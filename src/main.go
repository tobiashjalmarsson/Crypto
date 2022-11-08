package main

import "example.com/crypto/enigma"
func main() {
	var to_encrypt string = "cccccccc"
	r := enigma.CreateRotor(to_encrypt, 2)
	r.Encrypt()
	r.Decrypt()
	r.DisplayRotor()
}
