package main

import (
	"errors"
	"testing"
)

func TestNormalizeURL(t *testing.T) {
	tests := []struct {
		name     string
		inputURL string
		expected string
		err      error
	}{
		{
			name:     "remove scheme",
			inputURL: "https://blog.boot.dev/path",
			expected: "blog.boot.dev/path",
			err:      nil,
		},
		{
			name:     "invalid scheme",
			inputURL: "htp://blog.boot.dev/path",
			expected: "",
			err:      errors.New("malformed URL"),
		},
		{
			name:     "missing //",
			inputURL: "http:blog.boot.dev/path",
			expected: "",
			err:      errors.New("malformed URL"),
		},
		{
			name:     "extra slash early",
			inputURL: "https:///blog.boot.dev/path",
			expected: "",
			err:      errors.New("malformed URL"),
		},
		{
			name:     "extra slash end",
			inputURL: "https://blog.boot.dev/path//",
			expected: "",
			err:      errors.New("malformed URL"),
		},
		{
			name:     "no path, no slash",
			inputURL: "https://blog.boot.dev",
			expected: "blog.boot.dev/",
			err:      nil,
		},
		{
			name:     "no path, slash",
			inputURL: "https://blog.boot.dev/",
			expected: "blog.boot.dev/",
			err:      nil,
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := normalizeURL(tc.inputURL)
			if tc.err == nil && err != nil {
				t.Errorf("expected no error, got %v", err)
				return
			}
			if tc.err != nil && err == nil {
				t.Errorf("expected error %v, got nil", tc.err)
				return
			}
			if tc.err != nil && err != nil && err.Error() != tc.err.Error() {
				t.Errorf("expected error %v, got %v", tc.err, err)
				return
			}
			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}
