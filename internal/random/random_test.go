package random_test

import (
	"testing"

	"github.com/Kareky/cryptography/internal/alphabet"
	"github.com/Kareky/cryptography/internal/random"
)

func TestGenerateRandomInt(t *testing.T) {
	type args struct {
		length int
		max    int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Generate 5 random integers with max 10",
			args: args{
				length: 5,
				max:    10,
			},
			wantErr: false,
		},
		{
			name: "Generate 0 random integers with max 10",
			args: args{
				length: 0,
				max:    10,
			},
			wantErr: false,
		},
		{
			name: "Generate 5 random integers with max 0 (invalid)",
			args: args{
				length: 5,
				max:    0,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := random.GeneratePad(tt.args.length, tt.args.max)
			if (err != nil) != tt.wantErr {
				t.Errorf("GeneratePad() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && len(got) != tt.args.length {
				t.Errorf("GeneratePad() got length = %v, want %v", len(got), tt.args.length)
			}
			if !tt.wantErr {
				for _, num := range got {
					if num < 0 || num >= tt.args.max {
						t.Errorf("GeneratePad() got number = %v, want in range [0, %v)", num, tt.args.max)
					}
				}
			}
		})
	}
}

func TestGenerateRandomIntWithAlphabet(t *testing.T) {
	type args struct {
		length int
		a      alphabet.Alphabet
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Generate 5 random integers with Latin alphabet",
			args: args{
				length: 5,
				a:      alphabet.LatinAlphabet,
			},
			wantErr: false,
		},
		{
			name: "Generate 0 random integers with Latin alphabet",
			args: args{
				length: 0,
				a:      alphabet.LatinAlphabet,
			},
			wantErr: false,
		},
		{
			name: "Generate 5 random integers with nil alphabet (invalid)",
			args: args{
				length: 5,
				a:      nil,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := random.GeneratePadWithAlphabet(tt.args.length, tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("GeneratePadWithAlphabet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && len(got) != tt.args.length {
				t.Errorf("GeneratePadWithAlphabet() got length = %v, want %v", len(got), tt.args.length)
			}
			if !tt.wantErr {
				max := tt.args.a.Len()
				for _, num := range got {
					if num < 0 || num >= max {
						t.Errorf("GeneratePadWithAlphabet() got number = %v, want in range [0, %v)", num, max)
					}
				}
			}
		})
	}
}

func TestGenerateRandomIntFromText(t *testing.T) {
	type args struct {
		text string
		a    alphabet.Alphabet
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Generate random integers from text with Latin alphabet",
			args: args{
				text: "Hello, World!",
				a:    alphabet.LatinAlphabet,
			},
			wantErr: false,
		},
		{
			name: "Generate random integers from empty text with Latin alphabet",
			args: args{
				text: "",
				a:    alphabet.LatinAlphabet,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := random.GeneratePadFromText(tt.args.text, tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("GeneratePadFromText() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
