package encryptor

import "testing"

func TestSign(t *testing.T) {
	sign := Sign("YV78Pyj1VvqdNGpMJ1pHic0bIBOWMv", 1711001766, "abc", nil)
	if sign != "fff8dae1356e7867ea98743439f0e9f8" {
		t.Fatalf("bad sign: %s", sign)
	}
}

func TestEncryptDecrypt(t *testing.T) {
	aesKey := "q1Os1ZMe0nG28KUEx9lg3HjK7V5QyXvi212fzsgDqgz"
	plainText := []byte(`{"query":"hello"}`)

	cipherText, err := Encrypt(aesKey, plainText)
	if err != nil {
		t.Fatalf("Encrypt error: %v", err)
	}
	got, err := Decrypt(aesKey, cipherText)
	if err != nil {
		t.Fatalf("Decrypt error: %v", err)
	}
	if string(got) != string(plainText) {
		t.Fatalf("bad plaintext: %s", got)
	}
}

func TestDecryptInvalidKey(t *testing.T) {
	if _, err := Encrypt("bad-key", []byte("hello")); err == nil {
		t.Fatal("Encrypt should return error")
	}
}
