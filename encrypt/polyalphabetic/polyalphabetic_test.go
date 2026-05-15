package polyalphabetic_test

import (
	"testing"
	"github.com/Kareky/cryptography/encrypt/polyalphabetic"
	"github.com/Kareky/cryptography/internal/alphabet"
)

func TestEncrypt(t *testing.T) {
	type args struct {
		text        string
		encryptWord string
		Alphabet    alphabet.Alphabet
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Encrypt with word 'key'",
			args: args{
				text:       "Hello, World!",
				encryptWord: "key",
				Alphabet:    alphabet.LatinAlphabet,
			},
			want: "rijvs, uyvjn!",
		},
		{
			name: "Encrypt with word 'abc'",
			args: args{
				text:       "Hello, World!",
				encryptWord: "abc",
				Alphabet:    alphabet.LatinAlphabet,
			},
			want: "hfnlp, yosnd!",
		},
		{
			name: "Encrypt with empty word",
			args: args{
				text:       "Hello, World!",
				encryptWord: "",
				Alphabet:    alphabet.LatinAlphabet,
			},
			want: "Hello, World!",
		},
		{
			name: "Encrypt with word containing non-alphabetic characters",
			args: args{
				text:       "Hello, World!",
				encryptWord: "k3y",
				Alphabet:    alphabet.LatinAlphabet,
			},
			want: "Error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := polyalphabetic.Encrypt(tt.args.text, tt.args.encryptWord, tt.args.Alphabet)
			if tt.want == "Error" {
				if err == nil {
					t.Errorf("Encrypt() expected an error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("Encrypt() returned an unexpected error: %v", err)
				}
				if got != tt.want {
					t.Errorf("Encrypt() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}