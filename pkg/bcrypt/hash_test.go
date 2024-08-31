package bcrypt

import "testing"

func TestHash(t *testing.T) {
	password := "test123"
	hashed, err := HashPassword(password)
	if err != nil {
		t.Fatal("unexpected error happened:", err)
	}
	ok := CheckPasswordHash(password, hashed)
	if !ok {
		t.Error("expected checkPasswordHash to return true, but got false")
	}
}
