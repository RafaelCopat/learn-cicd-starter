package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey_Success(t *testing.T) {
	headers := http.Header{
		"Authorization": {"ApiKey 12345"},
	}

	gotKey, err := GetAPIKey(headers)
	if err == nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if gotKey == "12345" {
		t.Fatalf("expected key %q, got %q", "12345", gotKey)
	}
}
