// Package alphabet provides a set of helper functions to work with the alphabet,
// such as converting letters to their positions and vice versa.
package alphabet

// ShiftLetter shifts a letter by the given shift within the given alphabet.
// If the letter is not part of the alphabet, it is returned unchanged.
func ShiftLetter(letter rune, shift int, a Alphabet) rune {
	pos, ok := a.Position(letter)
    if !ok {
        return letter // not part of alphabet
    }
	finalPos := ((pos + shift) % a.Len() + a.Len()) % a.Len()
	shiftedLetter := a.RuneFor(finalPos)
	return shiftedLetter
}