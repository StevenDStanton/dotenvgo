package dotenvgo

import (
	"strings"
	"testing"
)

func TestParseFileLocation(t *testing.T) {
	tests := []struct {
		name     string
		params   []string
		expected string
	}{
		{
			name:     "Empty params",
			params:   []string{},
			expected: ".env",
		},
		{
			name:     "One entry",
			params:   []string{"../"},
			expected: "../.env",
		},
		{
			name:     "One entry",
			params:   []string{"cmd/"},
			expected: "cmd/.env",
		},
		{
			name:     "Multiple entries",
			params:   []string{"~/", "no!"},
			expected: "~/.env",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parseFileLocation(tt.params...)
			if result != tt.expected {
				t.Errorf("parseFileLocation(%v) = %v; want %v", tt.params, result, tt.expected)
			}
		})
	}

}

func FuzzParseFileLocation(f *testing.F) {
	// Seed the fuzzer with initial test cases
	f.Add("")
	f.Add("../")
	f.Add("cmd/")
	f.Add("~/")

	f.Fuzz(func(t *testing.T, param string) {
		// Convert the single string param to a slice with one element
		params := []string{param}
		// Call the function with the fuzzed input using ... to pass the slice as variadic arguments
		result := parseFileLocation(params...)

		// Basic validation to ensure result is not empty
		if result == "" {
			t.Errorf("parseFileLocation(%v) returned an empty string", params)
		}

		// Check that the result ends with ".env"
		if !strings.HasSuffix(result, ".env") {
			t.Errorf("parseFileLocation(%v) = %v; expected to end with '.env'", params, result)
		}
	})
}

func TestFetchFile(t *testing.T) {
	result, err := fetchFile("test/test.txt")
	if err != nil {
		t.Errorf("FetchFile Expected hello=world and recieved %v", err)
	}
	if string(result) != "hello=world" {
		t.Errorf("FetchFile Expected hello=world and recieved %s", result)
	}

	_, err = fetchFile("test/tes.txt")
	if err == nil {
		t.Errorf("Expected Error but no error thrown")
	}
}

func TestNormalizeLineEndings(t *testing.T) {
	fileContent, err := fetchFile("test/line-end-test.txt")

	if err != nil {
		t.Errorf("FetchFile Expected hello=world and recieved %v", err)
	}

	if !strings.Contains(string(fileContent), "\r") {
		t.Errorf("Unnormalized content should contain '\\r', but got %q", fileContent)
	}

	normalized := normalizeLineEndings(fileContent)

	if strings.Contains(normalized, "\r") {
		t.Errorf("Normalized content should not contain '\\r', but got %q", normalized)
	}

}

func TestParseFileContentToMap(t *testing.T) {
	vault := make(Vault)
	content := []byte("KEY1  =value1\nKEY2=value2\n# This is a comment\nKEY3=value3 # with comment")
	vault.parseFileContentToMap(content)

	expected := Vault{
		"KEY1": "value1",
		"KEY2": "value2",
		"KEY3": "value3",
	}

	for key, expectedValue := range expected {
		if value, exists := vault[key]; !exists || value != expectedValue {
			t.Errorf("Expected vault[%q] = %q, but got %q", key, expectedValue, value)
		}
	}
}

func TestLoad(t *testing.T) {
	result, err := Load("test/")

	if err != nil {
		t.Error("Unable to Load File")
	}

	expected := Vault{
		"hello": "world",
	}

	for key, expectedValue := range expected {
		if value, exists := result[key]; !exists || value != expectedValue {
			t.Errorf("Expected vault[%q] = %q, but got %q", key, expectedValue, value)
		}
	}

}
