package caesar_test

import (
	"testing"
	"github.com/Kareky/cryptography/encrypt/caesar"
	"github.com/Kareky/cryptography/internal/alphabet"
)

func TestEncrypt(t *testing.T) {
	type args struct {
		text        string
		shift      	int
		Alphabet	alphabet.Alphabet
	}
	tests := []struct {
		name 		string
		args 		args
		want 		string
	}{
		{
			name: "Encrypt with shift 3",
			args: args{
				text:       "Hello, World!",
				shift:      3,
				Alphabet: 	alphabet.LatinAlphabet,
			},
			want: "khoor, zruog!",
		},
		{
			name: "Encrypt with shift 5",
			args: args{
				text:       "Hello, World!",
				shift:      5,
				Alphabet: 	alphabet.LatinAlphabet,
			},
			want: "mjqqt, btwqi!",
		},
		{
			name: "Encrypt with shift 0",
			args: args{
				text:       "Hello, World!",
				shift:      0,
				Alphabet: 	alphabet.LatinAlphabet,
			},
			want: "hello, world!",
		},
		{
			name: "Encrypt with shift -3",
			args: args{
				text:       "Hello, World!",
				shift:      -3,
				Alphabet: 	alphabet.LatinAlphabet,
			},
			want: "ebiil, tloia!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := caesar.Encrypt(tt.args.text, tt.args.shift, tt.args.Alphabet); got != tt.want {
				t.Errorf("Encrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecrypt(t *testing.T) {
	type args struct {
		text        string
		shift      	int
		Alphabet	alphabet.Alphabet
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Decrypt with shift 3",
			args: args{
				text:       "Khoor, Zruog!",
				shift:      3,
				Alphabet: 	alphabet.LatinAlphabet,
			},
			want: "hello, world!",
		},
		{
			name: "Decrypt with shift 5",
			args: args{
				text:       "mjqqt, btwqi!",
				shift:      5,
				Alphabet: 	alphabet.LatinAlphabet,
			},
			want: "hello, world!",
		},
		{
			name: "Decrypt with shift 0",
			args: args{
				text:       "Hello, World!",
				shift:      0,
				Alphabet: 	alphabet.LatinAlphabet,
			},
			want: "hello, world!",
		},
		{
			name: "Decrypt with shift -3",
			args: args{
				text:       "Ebiil, Tloia!",
				shift:      -3,
				Alphabet: 	alphabet.LatinAlphabet,
			},
			want: "hello, world!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := caesar.Decrypt(tt.args.text, tt.args.shift, tt.args.Alphabet); got != tt.want {
				t.Errorf("Decrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}