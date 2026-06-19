package auth

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestMakeAndValidateJWT(t *testing.T) {
	userID := uuid.New()
	secret := "testsecret"

	token, err := MakeJWT(userID, secret, time.Hour)
	if err != nil {
		t.Fatalf("MakeJWT failed: %v", err)
	}

	gotID, err := ValidateJWT(token, secret)
	if err != nil {
		t.Fatalf("ValidateJWT failed: %v", err)
	}
	if gotID != userID {
		t.Errorf("expected %v, got %v", userID, gotID)
	}
}

func TestExpiredJWT(t *testing.T) {
	userID := uuid.New()
	secret := "testsecret"

	token, err := MakeJWT(userID, secret, -time.Hour)
	if err != nil {
		t.Fatalf("MakeJWT failed: %v", err)
	}

	_, err = ValidateJWT(token, secret)
	if err == nil {
		t.Fatal("Expected error for expired token, got nil.")
	}
}

func TestWrongSecret(t *testing.T) {
	userID := uuid.New()

	token, err := MakeJWT(userID, "correctsecret", time.Hour)
	if err != nil {
		t.Fatalf("MakeJWT failed: %v", err)
	}

	_, err = ValidateJWT(token, "wrongsecret")
	if err == nil {
		t.Fatal("expected error for wrong secret, got nil")
	}
}
