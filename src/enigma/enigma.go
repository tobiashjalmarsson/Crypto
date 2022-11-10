package enigma

import (
	"strings"
)

type Enigma struct {
	num_rotors int
	rotations int
	rotors []Rotor
	plugboard Plugboard
}

func CreateEnigma(num_rotors int, plugboard string, rotations int) *Enigma {
	e := Enigma{num_rotors: num_rotors, rotations : rotations}
	
	var rotors []Rotor = make([]Rotor, num_rotors)
	
	for i := 0; i < num_rotors; i++ {
		rotors[i] = *CreateRotor(0, false)
	}

	rotors[num_rotors-1].reflector = true
	e.rotors = rotors

	p := CreatePlugboard(plugboard)

	e.UpdateShifts(rotations)

	e.plugboard = *p

	return &e
}

func (e *Enigma) IncrementRotations() {
	e.rotations = e.rotations + 1
}

func (e *Enigma) UpdateShifts(rotations int) {
	for index, rotor := range e.rotors {
		// TODO update indexes of rotors
		// Rotors[0].shift = rotations
		// Rotors[i].shift = rotations/(i*26)
		if index == 0 {
			rotor.SetShift(rotations)
		} else {
			rotor.SetShift(rotations/(index*26))
		}
	}
}

func (e *Enigma) EncryptDecrypt(message string) string {
	encrypted_message := e.plugboard.Encrypt(message)
	var sb strings.Builder
	shift := e.rotations
	
	for index, char := range encrypted_message {
		// (1) Encrypt character/rune
		var s_char string = string(char)
		
		for i := 0; i < len(e.rotors); i++ {
			s_char = e.rotors[i].Encrypt(s_char)
		}

		for i := len(e.rotors); i >= 0; i-- {
			s_char = e.rotors[i].Encrypt(s_char)
		}
		// (2) Update shift

		e.UpdateShifts(index)
		e.IncrementRotations()

		// (3) Push character to sb
		sb.WriteString(s_char)
		
	}

	return sb.String()
}
