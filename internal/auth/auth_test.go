package auth_test

import (
	"net/http"
	"testing"

	"github.com/bootdotdev/learn-cicd-starter/internal/auth"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		header      http.Header
		expectedKey string
		expectedErr error
	}{
		{http.Header{"Authorization": []string{"ApiKey key"}}, "key", nil},
	}

	for _, tt := range tests {
		res, err := auth.GetAPIKey(tt.header)
		if res != tt.expectedKey || err != tt.expectedErr {
			t.Fatalf("res=%s, expected=%s, err=%v, expected=%v", res, tt.expectedKey, err, tt.expectedErr)
		}
	}
}
