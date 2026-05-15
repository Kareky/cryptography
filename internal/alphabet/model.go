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
    name  			string
    runes 			[]rune
	accentToBase	map[rune]rune
}

// NewAlphabet creates a new alphabet with the given name and letters, and an accentToBase map.
// The accentToBase map is a simple mapping which map an accented letter to is non-accented character
// (ex: 'é':'e').
// 
// Alphabet package provides a standard map 'CommonAccents' with the list of all most common accents
// already mapped.
func NewAlphabet(name string, letters []rune, accentToBase map[rune]rune) *AlphabetImpl {
    return &AlphabetImpl{name: name, runes: letters, accentToBase: accentToBase}
}

func (a *AlphabetImpl) Len() int             	{ return len(a.runes) }
func (a *AlphabetImpl) RuneFor(pos int) rune 	{ return a.runes[pos] }
func (a *AlphabetImpl) Name() string         	{ return a.name }
func (a *AlphabetImpl) ToSlice() []rune     	{ return append([]rune(nil), a.runes...) }

func (a *AlphabetImpl) Position(r rune) (int, bool) {
	if char, ok := a.accentToBase[r]; ok {
		r = char
	}

	for i, char := range a.runes {
		if char == r {
			return i, true
		}
	}

	return 0, false
}

// Pre‑defined alphabets
var LatinAlphabet = NewAlphabet("latin", []rune("abcdefghijklmnopqrstuvwxyz"), CommonAccents)
var SpanishAlphabet = NewAlphabet("spanish", []rune("abcdefghijklmnñopqrstuvwxyz"), CommonAccents)

// CommonAccents maps common accented characters to their base Latin letter.
// Do not use this map if any of these runes is part of the alphabet itself,
// as the accent would be stripped before the lookup.
var CommonAccents = map[rune]rune{
	'à' : 'a',
	'á' : 'a',
	'â' : 'a',
	'ã' : 'a',
	'ä' : 'a',
	'å' : 'a',

	'ç' : 'c',

	'è' : 'e',
	'é' : 'e',
	'ê' : 'e',
	'ë' : 'e',

	'ì' : 'i',
	'í' : 'i',
	'î' : 'i',
	'ï' : 'i',

	'ò' : 'o',
	'ó' : 'o',
	'ô' : 'o',
	'õ' : 'o',
	'ö' : 'o',

	'ù' : 'u',
	'ú' : 'u',
	'û' : 'u',
	'ü' : 'u',

	'ý' : 'y',
	'ÿ' : 'y',
}