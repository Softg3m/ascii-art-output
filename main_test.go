package main

import (
	"os"
	"strings"
	"testing"
)

// Test GenerateAscii basic functionality
func TestGenerateAscii_NotEmpty(t *testing.T) {
	result := GenerateAscii("A", "standard")

	if result == "" {
		t.Error("Expected ASCII output, got empty string")
	}
}

// Test invalid banner file

func TestGenerateAscii_InvalidBanner(t *testing.T) {
	result := GenerateAscii("A", "unknownbanner")

	if !strings.Contains(result, "Error reading banner file") {
		t.Error("Expected error message for invalid banner")
	}
}

// Test newline handling
func TestGenerateAscii_NewLine(t *testing.T) {
	result := GenerateAscii("A\\nB", "standard")

	if result == "" {
		t.Error("Expected multi-line ASCII output")
	}
}

// Test file writing
func TestOutputFile(t *testing.T) {
	fileName := "test_output.txt"
	text := "Hi"
	banner := "standard"

	ascii := GenerateAscii(text, banner)

	err := os.WriteFile(fileName, []byte(ascii), 0644)
	if err != nil {
		t.Fatalf("Failed to write file: %v", err)
	}

	// Check if file exists
	_, err = os.Stat(fileName)
	if os.IsNotExist(err) {
		t.Error("Expected file to be created")
	}

	// Clean up
	os.Remove(fileName)
}

// Test Usage function (just ensures it runs)
func TestUsage(t *testing.T) {
	Usage()
}
