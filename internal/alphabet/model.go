package alphabet

type Alphabet interface {
	Len() int                   	//length of the alphabet
	RuneFor(pos int) rune       	// the rune at position pos in the alphabet
	Position(r rune) (int, bool)	// the position of the rune in the alphabet
	ToSlice() []rune           		// return the alphabet as a slice
	Name() string               	// the name of the alphabet
}

type Latin []rune

var LatinAlphabet = Latin([]rune("abcdefghijklmnopqrstuvwxyz"))

func (l Latin) Len() int             	{ return len(l) }
func (l Latin) RuneFor(pos int) rune 	{ return l[pos] }
func (l Latin) Name() string         	{ return "latin" }
func (l Latin) ToSlice() []rune     	{ return LatinAlphabet }

func (l Latin) Position(r rune) (int, bool) {
	for i, char := range l {
		if char == r {
			return i, true
		}
	}

	return 0, false
}