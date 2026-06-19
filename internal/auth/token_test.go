package auth

import (
	"net/http"
	"testing"
)

func TestGetBearerToken(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorisation", "Bearer mytoken123")
	token, err := GetBearerToken(headers)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if token != "mytoken123" {
		t.Fatalf("expected 'mytoken123', got %v", token)
	}
	_, err = GetBearerToken(http.Header{})
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}
