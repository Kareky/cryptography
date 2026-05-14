package alphabet_test

import (
	"testing"
	"github.com/Kareky/cryptography/internal/alphabet"
)

func TestShiftLetter(t *testing.T) {
	type args struct {
		letter 		rune
		shift  		int
		Alphabet	alphabet.Alphabet
	}
	tests := []struct {
		name		string
		args		args
		want		rune
	}{
		{
			name: "Shift 'a' by 1",
			args: args{
				letter:  	'a',
				shift:   	1,
				Alphabet: 	alphabet.LatinAlphabet,
			},
			want: 'b',
		},
		{
			name: "Shift 'z' by 1",
			args: args{
				letter:  	'z',
				shift:   	1,
				Alphabet: 	alphabet.LatinAlphabet,
			},
			want: 'a',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := alphabet.ShiftLetter(tt.args.letter, tt.args.shift, tt.args.Alphabet); got != tt.want {
				t.Errorf("ShiftLetter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPosition(t *testing.T) {
	type args struct {
		letter  	rune
		Alphabet	alphabet.Alphabet
	}
	tests := []struct {
		name 		string
		args 		args
		want 		int
	}{
		{
			name: "Position of 'a'",
			args: args{
				letter:   	'a',
				Alphabet:	alphabet.LatinAlphabet,
			},
			want: 0,
		},
		{
			name: "Position of 'z'",
			args: args{
				letter:  'z',
				Alphabet:	alphabet.LatinAlphabet,
			},
			want: 25,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, ok := tt.args.Alphabet.Position(tt.args.letter); got != tt.want || !ok {
				t.Errorf("Position() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRuneFor(t *testing.T) {
	type args struct {
		position int
		Alphabet	alphabet.Alphabet
	}
	tests := []struct {
		name string
		args args
		want rune
	}{
		{
			name: "Letter at position 0",
			args: args{
				position: 0,
				Alphabet:	alphabet.LatinAlphabet,
			},
			want: 'a',
		},
		{
			name: "Letter at position 25",
			args: args{
				position: 25,
				Alphabet:	alphabet.LatinAlphabet,
			},
			want: 'z',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.Alphabet.RuneFor(tt.args.position); got != tt.want {
				t.Errorf("RuneFor() = %v, want %v", got, tt.want)
			}
		})
	}
}