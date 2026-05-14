package frequency_test

import (
	"slices"
	"testing"

	"github.com/Kareky/cryptography/analysis/frequency"
	"github.com/Kareky/cryptography/internal/alphabet"
)

func TestLetterFrequency(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Frequency of 'Lorem ipsum'",
			args: args{
				text: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
			},
			want: []int{29, 3, 16, 19, 38, 3, 3, 1, 42, 0, 0, 22, 17, 24, 29, 11, 5, 22, 18, 32, 29, 3, 0, 3, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := frequency.LetterFrequency(tt.args.text, alphabet.LatinAlphabet); !slices.Equal(got.Slice(), tt.want) {
				t.Errorf("LetterFrequency() = %v, want %v", got, tt.want)
			}
		})
	}
}