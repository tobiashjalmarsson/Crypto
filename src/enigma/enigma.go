package enigma

import "container/list"

type Enigma struct {
	message string
	num_rotors int8
	rotors *list.List
	plugboard Plugboard
}

func CreateEnigma(message string, num_rotors int8) *Enigma {
	e := Enigma{message: message, num_rotors: num_rotors}
	
	//var rotors [num_rotors]Rotor
	
	/*
	for A := 0; A < ; A++ {
		if A == num_rotors {
		}	
	}
	*/

	return &e
}

func (e *Enigma) SetMessage(m string) {
	e.message = m
}

