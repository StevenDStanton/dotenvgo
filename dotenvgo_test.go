package dotenvgo

import (
	"os"
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
	// Create a temporary file with known content
	tmpFile, err := os.CreateTemp("", "testfile")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	content := "hello=world"
	if _, err := tmpFile.WriteString(content); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	tmpFile.Close()

	result, err := fetchFile(tmpFile.Name())
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if string(result) != content {
		t.Errorf("Expected %s, got %s", content, result)
	}

	_, err = fetchFile("nonexistentfile.txt")
	if err == nil {
		t.Errorf("Expected error for nonexistent file, got nil")
	}
}

func TestNormalizeLineEndings(t *testing.T) {
	content := []byte("line1\r\nline2\rline3\n")
	expected := []string{"line1", "line2", "line3"}

	normalized := normalizeLineEndings(content)
	if len(normalized) != len(expected) {
		t.Fatalf("Expected %d lines, got %d", len(expected), len(normalized))
	}

	for i, line := range expected {
		if normalized[i] != line {
			t.Errorf("Expected line %d to be %q, got %q", i, line, normalized[i])
		}
	}
}
func TestParseFileContentToMap(t *testing.T) {
	vault := make(Vault)
	lines := []string{"KEY1=value1", "KEY2=value2", "# This is a comment", "", "", "KEY3=value3 # with comment"}
	vault.parseFileContentToMap(lines)

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
	// Create a temporary directory
	tmpDir, err := os.MkdirTemp("", "envdir")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Create the .env file within the temporary directory
	envFilePath := tmpDir + "/.env"
	tmpFile, err := os.Create(envFilePath)
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}

	content := "KEY1=value1\nKEY2=value2"
	if _, err := tmpFile.WriteString(content); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	if err := tmpFile.Close(); err != nil {
		t.Fatalf("Failed to close temp file: %v", err)
	}

	// Call Load with the directory path
	result, err := Load(true, tmpDir+"/")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	expected := Vault{
		"KEY1": "value1",
		"KEY2": "value2",
	}

	for key, expectedValue := range expected {
		if value, exists := result[key]; !exists || value != expectedValue {
			t.Errorf("Expected vault[%q] = %q, but got %q", key, expectedValue, value)
		}
	}
}

func TestSaveFileContentToEnvironment(t *testing.T) {
	lines := []string{"KEY1=value1", "KEY2=value2"}

	saveFileContentToEnviroment(lines)

	if os.Getenv("KEY1") != "value1" {
		t.Errorf("Expected KEY1 to be 'value1', got %q", os.Getenv("KEY1"))
	}
	if os.Getenv("KEY2") != "value2" {
		t.Errorf("Expected KEY2 to be 'value2', got %q", os.Getenv("KEY2"))
	}
}

func TestParseKeyValue(t *testing.T) {
	tests := []struct {
		line     string
		key      string
		value    string
		expected bool
	}{
		{"KEY=value", "KEY", "value", true},
		{"KEY = value ", "KEY", "value", true},
		{"KEY", "", "", false},
		{"KEY=value#comment", "KEY", "value", true},
	}

	for _, tt := range tests {
		key, value, exists := parseKeyValue(tt.line)
		if exists != tt.expected || key != tt.key || value != tt.value {
			t.Errorf("parseKeyValue(%q) = %q, %q, %v; want %q, %q, %v",
				tt.line, key, value, exists, tt.key, tt.value, tt.expected)
		}
	}
}
