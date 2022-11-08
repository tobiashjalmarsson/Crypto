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
	r := CreateRotor(original_string, 2)
	r.Encrypt()
	if encrypted_correct != r.encrypted {
		t.Fatalf(`r.Encrypt() yielded: %s, wanted: %s`, r.encrypted, encrypted_correct)
	}
}

func TestRotorEncryptCircular(t *testing.T) {
	original_string := "xxx"
	encrypted_correct := "aaa"
	r := CreateRotor(original_string, 2)
	r.Encrypt()
	if encrypted_correct != r.encrypted {
		t.Fatalf(`r.Encrypt() yielded: %s, wanted: %s`, r.encrypted, encrypted_correct)
	}
}

func TestRoEncryptFull(t *testing.T) {
	original_string := "abcdefghijklmnopqrstvwyuzx " 
	encrypted_correct := "cdefghijklmnopqrstvwyuzx ab"
	r := CreateRotor(original_string, 2)
	r.Encrypt()
	if encrypted_correct != r.encrypted {
		t.Fatalf(`r.Encrypt() yielded: %s, wanted: %s`, r.encrypted, encrypted_correct)
	}
}

func TestRotorEncryptFullLargeRotation(t *testing.T) {
	original_string := "abcdefghijklmnopqrstvwyuzx " 
	encrypted_correct := "klmnopqrstvwyuzx abcdefghij"
	r := CreateRotor(original_string, 10)
	r.Encrypt()
	if encrypted_correct != r.encrypted {
		t.Fatalf(`r.Encrypt() yielded: %s, wanted: %s`, r.encrypted, encrypted_correct)
	}
}

func TestRotorEncryptDecryptFull(t *testing.T) {
	original_string := "abcdefghijklmnopqrstvwyuzx "
	r := CreateRotor(original_string, 10)
	r.Encrypt()
	r.Decrypt()
	if original_string != r.decrypted {
		t.Fatalf(`r.Encrypt() (*) r.Decrypt() yielded: %s, wanted: %s`, r.decrypted, original_string)
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
	p := CreatePlugboard(original_message)
	p.Encrypt()
	if p.encrypted != expected_encrypted_message {
		t.Fatalf(`p.Encrypt() yielded: %s, wanted: %s`, p.encrypted, original_message)
	} 
}

func TestPlugEncryptDecrypt(t *testing.T) {
	original_message := "abcdefghijklmnopqrstvwyuzx " 
	p := CreatePlugboard(original_message)
	p.Encrypt()
	p.Decrypt()
	if p.decrypted != original_message {
		t.Fatalf(`p.Encrypt() (*) p.Decrypt() yielded: %s, wanted: %s`, p.decrypted, original_message)
	} 
}
