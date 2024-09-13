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

func TestSchemeValidator(t *testing.T) {
	test := []struct {
		source   string
		dest     string
		key      string
		expected bool
	}{
		{"ftp", "ftp", "ftp", true},
		{"sftp", "ftp", "ftp", false},
		{"sftp", "sftp", "ftp", false},
	}
	for _, tt := range test {
		t.Run(tt.source, func(t *testing.T) {
			result := SchemeValidator(tt.source, tt.dest, tt.key)
			if result != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestFtpParamsValidator(t *testing.T) {
	test := []struct {
		input    map[string]string
		expected bool
	}{
		{map[string]string{
			"host":   "Myhost",
			"scheme": "sftp",
			"user":   "myuser",
		}, false},
		{map[string]string{
			"host":   "",
			"scheme": "sftp",
			"user":   "myuser",
		}, true},
		{map[string]string{
			"host":   "",
			"scheme": "",
			"user":   "myuser",
		}, true},
	}

	for _, tt := range test {
		t.Run("FTP Parameters Validator", func(t *testing.T) {
			result := FtpParamsValidator(tt.input)
			if result != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}
