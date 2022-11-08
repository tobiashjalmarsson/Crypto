package main

import "example.com/crypto/enigma"
func main() {
	var to_encrypt string = "abcdefg"
	p := enigma.CreatePlugboard(to_encrypt)
	p.Encrypt()
	p.Decrypt()
}
