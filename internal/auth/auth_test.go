package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name    string
		headers http.Header
		want    string
		wantErr bool
	}{
		{
			name:    "no headers should throw error",
			headers: http.Header{},
			want:    "",
			wantErr: true,
		},
		{
			name: "empty Authorization header should throw error",
			headers: http.Header{
				"Authorization": []string{""},
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "malformed header with ApiKey part without actual key value should throw error",
			headers: http.Header{
				"Authorization": []string{"ApiKey"},
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "malformed header with dummy key type and value should throw error",
			headers: http.Header{
				"Authorization": []string{"DummyKey Test"},
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "correct api key in authorization header should return second part of string of header value",
			headers: http.Header{
				"Authorization": []string{"ApiKey Test"},
			},
			want:    "Test",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAPIKey(tt.headers)

			if (err != nil) != tt.wantErr {
				t.Fatalf("GetAPIKey() error = %v, wantErr %v", err, tt.wantErr)
			}

			if got != tt.want {
				t.Errorf("GetAPIKey() = %q, want %q", got, tt.want)
			}
		})
	}
}
