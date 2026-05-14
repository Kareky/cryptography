package alphabet_test

import (
	"testing"
	"github.com/Kareky/cryptography/internal/alphabet"
)

func TestShiftLetter(t *testing.T) {
	type args struct {
		letter  rune
		shift   int
		isUpper bool
	}
	tests := []struct {
		name string
		args args
		want rune
	}{
		{
			name: "Shift 'a' by 1, lowercase",
			args: args{
				letter:  'a',
				shift:   1,
				isUpper: false,
			},
			want: 'b',
		},
		{
			name: "Shift 'z' by 1, lowercase",
			args: args{
				letter:  'z',
				shift:   1,
				isUpper: false,
			},
			want: 'a',
		},
		{
			name: "Shift 'A' by -1, uppercase",
			args: args{
				letter:  'A',
				shift:   -1,
				isUpper: true,
			},
			want: 'Z',
		},
		{
			name: "Shift 'Z' by -1, uppercase",
			args: args{
				letter:  'Z',
				shift:   -1,
				isUpper: true,
			},
			want: 'Y',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := alphabet.ShiftLetter(tt.args.letter, tt.args.shift, tt.args.isUpper); got != tt.want {
				t.Errorf("ShiftLetter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFromAlphabetToPosition(t *testing.T) {
	type args struct {
		letter  rune
		isUpper bool
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Position of 'a', lowercase",
			args: args{
				letter:  'a',
				isUpper: false,
			},
			want: 1,
		},
		{
			name: "Position of 'z', lowercase",
			args: args{
				letter:  'z',
				isUpper: false,
			},
			want: 26,
		},
		{
			name: "Position of 'A', uppercase",
			args: args{
				letter:  'A',
				isUpper: true,
			},
			want: 1,
		},
		{
			name: "Position of 'Z', uppercase",
			args: args{
				letter:  'Z',
				isUpper: true,
			},
			want: 26,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := alphabet.FromAlphabetToPosition(tt.args.letter, tt.args.isUpper); got != tt.want {
				t.Errorf("FromAlphabetToPosition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFromPositionToAlphabet(t *testing.T) {
	type args struct {
		position int
		toUpper bool
	}
	tests := []struct {
		name string
		args args
		want rune
	}{
		{
			name: "Letter at position 1, lowercase",
			args: args{
				position: 1,
				toUpper: false,
			},
			want: 'a',
		},
		{
			name: "Letter at position 26, lowercase",
			args: args{
				position: 26,
				toUpper: false,
			},
			want: 'z',
		},
		{
			name: "Letter at position 1, uppercase",
			args: args{
				position: 1,
				toUpper: true,
			},
			want: 'A',
		},
		{
			name: "Letter at position 26, uppercase",
			args: args{
				position: 26,
				toUpper: true,
			},
			want: 'Z',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := alphabet.FromPositionToAlphabet(tt.args.position, tt.args.toUpper); got != tt.want {
				t.Errorf("FromPositionToAlphabet() = %v, want %v", got, tt.want)
			}
		})
	}
}