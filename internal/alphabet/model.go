package alphabet

type Alphabet interface {
	Len() int                   	//length of the alphabet
	RuneFor(pos int) rune       	// the rune at position pos in the alphabet
	Position(r rune) (int, bool)	// the position of the rune in the alphabet
	MapOrdered() map[rune]int    	// a map of rune and their position
	MapEmpty() map[rune]int			// a map of runes with 0 value, normally used to count characters
	Alphabet() []rune           	// the whole alphabet
	Name() string               	// the name of the alphabet
}

type Latin []rune

func (l Latin) Len() int             { return len(l) }
func (l Latin) RuneFor(pos int) rune { return l[pos] }
func (l Latin) Name() string         { return "latin" }
func (l Latin) Alphabet() []rune     { return []rune("abcdefghijklmnopqrstuvwxyz") }
func (l Latin) MapOrdered() map[rune]int {
	alph := make(map[rune]int)
	for i, character := range l.Alphabet() {
		alph[character] = i
	}

	return alph
}

func (l Latin) MapEmpty() map[rune]int {
	alph := make(map[rune]int)
	for _, character := range l.Alphabet() {
		alph[character] = 0
	}

	return alph
}

func (l Latin) Position(r rune) (int, bool) {
	pos, ok := l.MapOrdered()[r]
	return pos, ok
}