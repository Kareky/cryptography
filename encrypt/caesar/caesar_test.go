package caesar_test

import (
	"testing"
	"github.com/Kareky/cryptography/encrypt/caesar"
)

func TestEncrypt(t *testing.T) {
	type args struct {
		text                string
		shift               int
		preserveUppercase   bool
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Encrypt with shift 3 and preserve uppercase",
			args: args{
				text:              "Hello, World!",
				shift:             3,
				preserveUppercase: true,
			},
			want: "Khoor, Zruog!",
		},
		{
			name: "Encrypt with shift 5 and do not preserve uppercase",
			args: args{
				text:              "Hello, World!",
				shift:             5,
				preserveUppercase: false,
			},
			want: "mjqqt, btwqi!",
		},
		{
			name: "Encrypt with shift 0 and preserve uppercase",
			args: args{
				text:              "Hello, World!",
				shift:             0,
				preserveUppercase: true,
			},
			want: "Hello, World!",
		},
		{
			name: "Encrypt with shift -3 and preserve uppercase",
			args: args{
				text:              "Hello, World!",
				shift:             -3,
				preserveUppercase: true,
			},
			want: "Ebiil, Tloia!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := caesar.Encrypt(tt.args.text, tt.args.shift, tt.args.preserveUppercase); got != tt.want {
				t.Errorf("Encrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecrypt(t *testing.T) {
	type args struct {
		text                string
		shift               int
		preserveUppercase   bool
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Decrypt with shift 3 and preserve uppercase",
			args: args{
				text:              "Khoor, Zruog!",
				shift:             3,
				preserveUppercase: true,
			},
			want: "Hello, World!",
		},
		{
			name: "Decrypt with shift 5 and do not preserve uppercase",
			args: args{
				text:              "mjqqt, btwqi!",
				shift:             5,
				preserveUppercase: false,
			},
			want: "hello, world!",
		},
		{
			name: "Decrypt with shift 0 and preserve uppercase",
			args: args{
				text:              "Hello, World!",
				shift:             0,
				preserveUppercase: true,
			},
			want: "Hello, World!",
		},
		{
			name: "Decrypt with shift -3 and preserve uppercase",
			args: args{
				text:              "Ebiil, Tloia!",
				shift:             -3,
				preserveUppercase: true,
			},
			want: "Hello, World!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := caesar.Decrypt(tt.args.text, tt.args.shift, tt.args.preserveUppercase); got != tt.want {
				t.Errorf("Decrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}