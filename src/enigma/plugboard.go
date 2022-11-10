package enigma

import (
	"strings"
)

type Plugboard struct {
	board string
}


func CreatePlugboard(board string) *Plugboard {
	if board == "" {
		return &Plugboard{board: characters}
	} else {
		return &Plugboard{board: board}
	}
}

func (p *Plugboard) Encrypt(message string) string {
	var sb strings.Builder
	
	for _, char := range message {
		char_pos := strings.IndexRune(p.board, char)
		sb.WriteByte(default_board[char_pos])
	}
	return sb.String()
}

func (p *Plugboard) Decrypt(message string) string {
	var sb strings.Builder

	for _, char := range message {
		char_pos := strings.IndexRune(p.board, char)
		sb.WriteByte(default_board[char_pos])
	}

	return sb.String()
}
