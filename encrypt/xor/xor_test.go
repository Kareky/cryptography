package xor_test

import (
	"bytes"
	"testing"

	"github.com/Kareky/cryptography/encrypt/xor"
)

func TestEncrypt(t *testing.T) {
	tests := []struct {
		name string
		data []byte
		key  []byte
		want []byte
	}{
		{
			name: "simple ASCII",
			data: []byte("hello"),
			key:  []byte{0x01, 0x02, 0x03},
			want: []byte{0x69, 0x67, 0x6f, 0x6d, 0x6d},
		},
		{
			name: "empty key returns data unchanged",
			data: []byte("abc"),
			key:  []byte{},
			want: []byte("abc"),
		},
		{
			name: "empty data with key",
			data: []byte{},
			key:  []byte{0x10},
			want: []byte{},
		},
		{
			name: "key same length as data",
			data: []byte{0xFF, 0x00, 0xAA},
			key:  []byte{0x0F, 0x0F, 0x0F},
			want: []byte{0xF0, 0x0F, 0xA5},
		},
		{
			name: "repeating key XOR",
			data: []byte{0x00, 0x00, 0x00, 0x00},
			key:  []byte{0xAB, 0xCD},
			want: []byte{0xAB, 0xCD, 0xAB, 0xCD},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := xor.Encrypt(tt.data, tt.key)
			if !bytes.Equal(got, tt.want) {
				t.Errorf("Encrypt() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func TestDecrypt(t *testing.T) {
	// Since Decrypt is just Encrypt, a round‑trip test suffices.
	original := []byte("secret message")
	key := []byte{0x12, 0x34, 0x56}
	enc := xor.Encrypt(original, key)
	dec := xor.Decrypt(enc, key)
	if !bytes.Equal(dec, original) {
		t.Errorf("Decrypt(Encrypt(data)) = %s, want %s", dec, original)
	}
}

func TestEncryptWithPad(t *testing.T) {
	data := []byte("hello world")
	cipher, pad, err := xor.EncryptWithPad(data)
	if err != nil {
		t.Fatalf("EncryptWithPad() unexpected error: %v", err)
	}
	if len(cipher) != len(data) {
		t.Errorf("cipher length = %d, want %d", len(cipher), len(data))
	}
	if len(pad) != len(data) {
		t.Errorf("pad length = %d, want %d", len(pad), len(data))
	}
	// Decrypt should recover original
	dec := xor.Decrypt(cipher, pad)
	if !bytes.Equal(dec, data) {
		t.Errorf("Decrypt(cipher, pad) = %s, want %s", dec, data)
	}
}