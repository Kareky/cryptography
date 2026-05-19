package polybius_test

import (
	"reflect"
	"testing"

	"github.com/Kareky/cryptography/encrypt/polybius"
	"github.com/Kareky/cryptography/internal/coordinates"
)

var testSquare = polybius.LatinNoJSquare

func TestEncrypt(t *testing.T) {
	tests := []struct {
		name    string
		square  polybius.Square
		text    string
		want    []coordinates.Coordinate
		wantErr bool
	}{
		{
			name:    "simple word",
			square:  testSquare,
			text:    "hello",
			want: []coordinates.Coordinate{
				{Row: 1, Col: 2},
				{Row: 0, Col: 4},
				{Row: 2, Col: 0},
				{Row: 2, Col: 0},
				{Row: 2, Col: 3},
			},
			wantErr: false,
		},
		{
			name:    "space and punctuation skipped",
			square:  testSquare,
			text:    "a b!",
			want: []coordinates.Coordinate{
				{Row: 0, Col: 0},
				{Row: 0, Col: 1},
			},
			wantErr: false,
		},
		{
			name:    "uppercase handled",
			square:  testSquare,
			text:    "HELLO",
			want: []coordinates.Coordinate{
				{Row: 1, Col: 2},
				{Row: 0, Col: 4},
				{Row: 2, Col: 0},
				{Row: 2, Col: 0},
				{Row: 2, Col: 3},
			},
			wantErr: false,
		},
		{
			name:    "nil square",
			square:  nil,
			text:    "abc",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "empty text",
			square:  testSquare,
			text:    "",
			want:    []coordinates.Coordinate{},
			wantErr: false,
		},
		{
			name:    "digits are skipped",
			square:  testSquare,
			text:    "123",
			want:    []coordinates.Coordinate{},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := polybius.Encrypt(tt.text, tt.square)
			if (err != nil) != tt.wantErr {
				t.Errorf("Encrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Encrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecrypt(t *testing.T) {
	tests := []struct {
		name        string
		square      polybius.Square
		coordinates []coordinates.Coordinate
		want        string
		wantErr     bool
	}{
		{
			name:   "simple word",
			square: testSquare,
			coordinates: []coordinates.Coordinate{
				{Row: 1, Col: 2},
				{Row: 0, Col: 4},
				{Row: 2, Col: 0},
				{Row: 2, Col: 0},
				{Row: 2, Col: 3},
			},
			want:    "hello",
			wantErr: false,
		},
		{
			name:   "invalid coordinate skipped",
			square: testSquare,
			coordinates: []coordinates.Coordinate{
				{Row: 0, Col: 0},
				{Row: 99, Col: 99},
				{Row: 0, Col: 1},
			},
			want:    "ab",
			wantErr: false,
		},
		{
			name:    "nil square",
			square:  nil,
			coordinates: []coordinates.Coordinate{
				{Row: 0, Col: 0},
			},
			want:    "",
			wantErr: true,
		},
		{
			name:        "empty coordinates",
			square:      testSquare,
			coordinates: []coordinates.Coordinate{},
			want:        "",
			wantErr:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := polybius.Decrypt(tt.coordinates, tt.square)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("Decrypt() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestRoundTrip(t *testing.T) {
	plain := "hello world"
	coords, err := polybius.Encrypt(plain, testSquare)
	if err != nil {
		t.Fatalf("Encrypt() error: %v", err)
	}
	dec, err := polybius.Decrypt(coords, testSquare)
	if err != nil {
		t.Fatalf("Decrypt() error: %v", err)
	}
	if dec != "helloworld" {
		t.Errorf("round‑trip = %q, want %q", dec, "helloworld")
	}
}