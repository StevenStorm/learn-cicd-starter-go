package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := map[string]struct {
		authHeader     string
		expectedAPIKey string
		expectError    error
	}{
		"no auth header": {
			authHeader:     "",
			expectedAPIKey: "",
			expectError:    ErrNoAuthHeaderIncluded,
		},
		"malformed auth header": {
			authHeader:     "111111",
			expectedAPIKey: "",
			expectError:    ErrMalformedAuthHeader,
		},
		"invalid api key format": {
			authHeader:     "11111 22222",
			expectedAPIKey: "",
			expectError:    ErrMalformedAuthHeader,
		},
		"valid api key": {
			authHeader:     "ApiKey 11111",
			expectedAPIKey: "11111",
			expectError:    nil,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			headers := http.Header{}
			headers.Set("Authorization", test.authHeader)
			apiKey, err := GetAPIKey(headers)
			if err != test.expectError {
				t.Fatalf("Expected Error: %v, got: %v", test.expectError, err)
			}
			if apiKey != test.expectedAPIKey {
				t.Fatalf("Expected API Key: %v, got: %v", test.expectedAPIKey, apiKey)
			}
		})
	}
}
