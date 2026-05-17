package xor

import "github.com/Kareky/cryptography/internal/random"

// Encrypt applies a repeating‑key XOR cipher to data using key.
// If key is empty, data is returned unchanged.
func Encrypt(data []byte, key []byte) []byte {
	if len(key) == 0 {
		return data
	}

	return encrypt(data, key)
}

// Decrypt decrypts data that was encrypted with the same repeating‑key XOR cipher.
// It is identical to calling Encrypt with the same key.
func Decrypt(data []byte, key []byte) []byte {
	return Encrypt(data, key)
}

// EncryptWithPad generates a cryptographically secure random one‑time pad
// exactly as long as data, encrypts data with it, and returns both the
// ciphertext and the pad. The caller must retain the pad to decrypt.
// It returns an error if the random source fails.
func EncryptWithPad(data []byte) ([]byte, []byte, error) {
	pad, err := random.GenerateBytes(len(data))
	if err != nil {
		return nil, nil, err
	}

	return encrypt(data, pad), pad, nil
}

// encrypt contains the shared XOR encryption logic.
// It XORs each byte of data with the corresponding byte of key,
// repeating key cyclically. The caller must guarantee key is non‑empty.
func encrypt(data []byte, key []byte) []byte {
	encryptedData := make([]byte, 0, len(data))
	for i, b := range data {
		encryptedData = append(encryptedData, b^key[i%len(key)])
	}

	return encryptedData
}