package enigma

import (
	"testing"
)

/*
Rotor tests beginning
*/
func TestRotorEncrypt(t *testing.T) {
	original_string := "abc"
	encrypted_correct := "cde"
	r := CreateRotor(2, false)
	encrypted_string := r.Encrypt(original_string)
	if encrypted_correct != encrypted_string{
		t.Fatalf(`r.Encrypt() yielded: %s, wanted: %s`, encrypted_string, encrypted_correct)
	}
}

func TestRotorEncryptCircular(t *testing.T) {
	original_string := "xxx"
	encrypted_correct := "aaa"
	r := CreateRotor(2, false)
	encrypted_string := r.Encrypt(original_string)
	if encrypted_string != encrypted_correct {
		t.Fatalf(`r.Encrypt() yielded: %s, wanted: %s`, encrypted_string, encrypted_correct)
	}
}

func TestRoEncryptFull(t *testing.T) {
	original_string := "abcdefghijklmnopqrstvwyuzx " 
	encrypted_correct := "cdefghijklmnopqrstvwyuzx ab"
	r := CreateRotor(2, false)
	encrypted_string := r.Encrypt(original_string)
	if encrypted_correct != encrypted_string {
		t.Fatalf(`r.Encrypt() yielded: %s, wanted: %s`, encrypted_string, encrypted_correct)
	}
}

func TestRotorEncryptFullLargeRotation(t *testing.T) {
	original_string := "abcdefghijklmnopqrstvwyuzx " 
	encrypted_correct := "klmnopqrstvwyuzx abcdefghij"
	r := CreateRotor(10, false)
	encrypted_string := r.Encrypt(original_string)
	if encrypted_correct != encrypted_string {
		t.Fatalf(`r.Encrypt() yielded: %s, wanted: %s`, encrypted_string, encrypted_correct)
	}
}

func TestRotorEncryptDecryptFull(t *testing.T) {
	original_string := "abcdefghijklmnopqrstvwyuzx "
	r := CreateRotor(10, false)
	encrypted_string := r.Encrypt(original_string)
	decrypted_string := r.Decrypt(encrypted_string)
	if original_string != decrypted_string {
		t.Fatalf(`r.Encrypt() (*) r.Decrypt() yielded: %s, wanted: %s`, decrypted_string, original_string)
	}
}

/*
Rotor tests ending
*/

/*
Plugboard tests starting
*/

func TestPlugEncrypt(t *testing.T) {
	original_message := "abab"
	expected_encrypted_message := "baba"
	p := CreatePlugboard("")
	encrypted_string := p.Encrypt(original_message)
	if encrypted_string != expected_encrypted_message {
		t.Fatalf(`p.Encrypt() yielded: %s, wanted: %s`, encrypted_string, original_message)
	} 
}

func TestPlugEncryptDecrypt(t *testing.T) {
	original_message := "abcdefghijklmnopqrstvwyuzx " 
	p := CreatePlugboard("")
	encrypted_string := p.Encrypt(original_message)
	decrypted_string := p.Decrypt(encrypted_string)
	if decrypted_string != original_message {
		t.Fatalf(`p.Encrypt() (*) p.Decrypt() yielded: %s, wanted: %s`, decrypted_string, original_message)
	} 
}
