package utils

import (
	"testing"
)

func TestNullChecker(t *testing.T) {
	test := []struct {
		input    string
		expected bool
	}{
		{"", true},
		{"ftp://user:password@host/path", false},
	}

	for _, tt := range test {
		t.Run(tt.input, func(t *testing.T) {
			result := NullChecker(tt.input)
			if result != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}
