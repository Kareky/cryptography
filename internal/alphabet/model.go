package alphabet

// Alphabet defines the interface for an alphabet, which includes methods to get the length of the alphabet,
// get a rune for a given position, get the position of a given rune, convert the alphabet to a slice of runes,
// and get the name of the alphabet.
type Alphabet interface {
	Len() int                   	//length of the alphabet
	RuneFor(pos int) rune       	// the rune at position pos in the alphabet
	Position(r rune) (int, bool)	// the position of the rune in the alphabet
	ToSlice() []rune           		// return the alphabet as a slice
	Name() string               	// the name of the alphabet
}

type AlphabetImpl struct {
    name  string
    runes []rune
}

// NewAlphabet creates a new alphabet with the given name and letters.
func NewAlphabet(name string, letters []rune) *AlphabetImpl {
    return &AlphabetImpl{name: name, runes: letters}
}



func (a *AlphabetImpl) Len() int             	{ return len(a.runes) }
func (a *AlphabetImpl) RuneFor(pos int) rune 	{ return a.runes[pos] }
func (a *AlphabetImpl) Name() string         	{ return a.name }
func (a *AlphabetImpl) ToSlice() []rune     	{ return append([]rune(nil), a.runes...) }

func (a *AlphabetImpl) Position(r rune) (int, bool) {
	for i, char := range a.runes {
		if char == r {
			return i, true
		}
	}

	return 0, false
}

// Pre‑defined alphabets
var LatinAlphabet = NewAlphabet("latin", []rune("abcdefghijklmnopqrstuvwxyz"))
var SpanishAlphabet = NewAlphabet("spanish", []rune("abcdefghijklmnñopqrstuvwxyz"))
