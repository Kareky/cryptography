package otp_test

import (
	"testing"

	"github.com/Kareky/cryptography/encrypt/otp"
	"github.com/Kareky/cryptography/internal/alphabet"
)

var latin = alphabet.LatinAlphabet

func TestEncrypt(t *testing.T) {
	tests := []struct {
		name    string
		text    string
		pad     []int
		want    string
		wantErr bool
	}{
		{
			name:    "basic encryption",
			text:    "abc",
			pad:     []int{1, 2, 3},
			want:    "bdf",
			wantErr: false,
		},
		{
			name:    "with space and punctuation preserved",
			text:    "a b!",
			pad:     []int{1, 1},
			want:    "b c!",
			wantErr: false,
		},
		{
			name:    "empty pad with non‑empty text",
			text:    "a",
			pad:     []int{},
			want:    "",
			wantErr: true,
		},
		{
			name:    "pad length mismatch",
			text:    "abc",
			pad:     []int{1, 2},
			want:    "",
			wantErr: true,
		},
		{
			name:    "pad longer than needed",
			text:    "a",
			pad:     []int{1, 2},
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := otp.Encrypt(tt.text, tt.pad, latin)
			if (err != nil) != tt.wantErr {
				t.Errorf("Encrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("Encrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEncryptStripped(t *testing.T) {
	tests := []struct {
		name    string
		text    string
		pad     []int
		want    string
		wantErr bool
	}{
		{
			name:    "strips spaces and punctuation",
			text:    "a b!",
			pad:     []int{1, 1},
			want:    "bc",
			wantErr: false,
		},
		{
			name:    "length mismatch after stripping",
			text:    "a b",
			pad:     []int{1},
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := otp.EncryptStripped(tt.text, tt.pad, latin)
			if (err != nil) != tt.wantErr {
				t.Errorf("EncryptStripped() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("EncryptStripped() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEncryptWithPad(t *testing.T) {
	text := "hello world"
	cipher, pad, err := otp.EncryptWithPad(text, latin)
	if err != nil {
		t.Fatalf("EncryptWithPad() unexpected error: %v", err)
	}
	if len(cipher) != len(text) {
		t.Errorf("cipher length = %d, want %d", len(cipher), len(text))
	}
	if len(pad) != 10 {
		t.Errorf("pad length = %d, want 10", len(pad))
	}

	dec, err := otp.Decrypt(cipher, pad, latin)
	if err != nil {
		t.Fatalf("Decrypt() unexpected error: %v", err)
	}
	if dec != text {
		t.Errorf("Decrypt() = %v, want %v", dec, text)
	}
}

func TestEncryptWithPadStripped(t *testing.T) {
	text := "hello world"
	cipher, pad, err := otp.EncryptWithPadStripped(text, latin)
	if err != nil {
		t.Fatalf("EncryptWithPadStripped() unexpected error: %v", err)
	}
	if len(cipher) != 10 {
		t.Errorf("stripped cipher length = %d, want 10", len(cipher))
	}
	if len(pad) != 10 {
		t.Errorf("pad length = %d, want 10", len(pad))
	}
}

func TestDecrypt(t *testing.T) {
	tests := []struct {
		name    string
		text    string
		pad     []int
		want    string
		wantErr bool
	}{
		{
			name:    "basic decryption",
			text:    "bdf",
			pad:     []int{1, 2, 3},
			want:    "abc",
			wantErr: false,
		},
		{
			name:    "with spaces preserved",
			text:    "b c!",
			pad:     []int{1, 1},
			want:    "a b!",
			wantErr: false,
		},
		{
			name:    "empty pad",
			text:    "a",
			pad:     []int{},
			want:    "",
			wantErr: true,
		},
		{
			name:    "pad length mismatch",
			text:    "abc",
			pad:     []int{1, 2},
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := otp.Decrypt(tt.text, tt.pad, latin)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("Decrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNilAlphabet(t *testing.T) {
	_, err := otp.Encrypt("a", []int{0}, nil)
	if err == nil {
		t.Error("Encrypt() with nil alphabet expected error")
	}
	_, _, err = otp.EncryptWithPad("a", nil)
	if err == nil {
		t.Error("EncryptWithPad() with nil alphabet expected error")
	}
	_, _, err = otp.EncryptWithPadStripped("a", nil)
	if err == nil {
		t.Error("EncryptWithPadStripped() with nil alphabet expected error")
	}
	_, err = otp.Decrypt("a", []int{0}, nil)
	if err == nil {
		t.Error("Decrypt() with nil alphabet expected error")
	}
}

func TestSentinelErrors(t *testing.T) {
	t.Run("ErrEmptyPad", func(t *testing.T) {
		_, err := otp.Encrypt("a", []int{}, latin)
		if err != otp.ErrEmptyPad {
			t.Errorf("Encrypt empty pad: got %v, want %v", err, otp.ErrEmptyPad)
		}
		_, err = otp.Decrypt("a", []int{}, latin)
		if err != otp.ErrEmptyPad {
			t.Errorf("Decrypt empty pad: got %v, want %v", err, otp.ErrEmptyPad)
		}
	})
	t.Run("ErrLengthMismatch", func(t *testing.T) {
		_, err := otp.Encrypt("abc", []int{1}, latin)
		if err != otp.ErrLengthMismatch {
			t.Errorf("Encrypt length mismatch: got %v, want %v", err, otp.ErrLengthMismatch)
		}
		_, err = otp.Decrypt("abc", []int{1}, latin)
		if err != otp.ErrLengthMismatch {
			t.Errorf("Decrypt length mismatch: got %v, want %v", err, otp.ErrLengthMismatch)
		}
	})
}