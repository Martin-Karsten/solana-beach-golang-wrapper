package solanabeach

import (
	"io"
	"net/http"
	"testing"
)

var handler = func(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<html><body>Hello World!</body></html>")
}

func TestGetResponseBody(t *testing.T) {
	tt := []struct {
		name   string
		url    string
		params map[string]interface{}
		body   []byte
	}{
		{"successful response", "mock-url", map[string]interface{}{"limit": "1"}, []byte{'s', 'u', 'c', 'c', 'e', 's'}},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if _, err := getResponseBody(tc.url, tc.params); err == nil {
				t.Fatalf("getResponseBody(%v, %v) => got %v, want %v ", tc.url, tc.params, err, tc.body)
			}
		})
	}
}
