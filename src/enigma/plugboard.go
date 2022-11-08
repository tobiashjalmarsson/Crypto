package enigma

import (
	"strings"
)

type Plugboard struct {
	unencrypted string
	encrypted string
	decrypted string
}


func CreatePlugboard(original_message string) *Plugboard {
	return &Plugboard{unencrypted: strings.ToLower(original_message)}
}

func (p *Plugboard) Encrypt() {
	var sb strings.Builder
	
	for _, char := range p.unencrypted {
		char_pos := strings.IndexRune(characters, char)
		sb.WriteByte(default_board[char_pos])
	}
	p.encrypted = sb.String()
}

func (p *Plugboard) Decrypt() {
	var sb strings.Builder

	for _, char := range p.encrypted {
		char_pos := strings.IndexRune(characters, char)
		sb.WriteByte(default_board[char_pos])
	}

	p.decrypted = sb.String()
}
