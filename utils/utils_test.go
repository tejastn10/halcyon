package utils

import "testing"

// TestGetCanonicalName tests the GetCanonicalName function
// Each test case specifies an input string and the expected output
func TestGetCanonicalName(t *testing.T) {
	tests := []struct {
		name     string // Test case name
		input    string // Input string
		expected string // Expected output
	}{
		// Basic cases
		{"standard file", "document.txt", "document.txt"},
		{"uppercase file", "DOCUMENT.TXT", "document.txt"},

		// Copy variations
		{"with copy suffix", "document copy.txt", "document.txt"},
		{"with copy uppercase", "document COPY.txt", "document.txt"},
		{"with underscore copy", "document_copy.txt", "document.txt"},
		{"with mixed case copy", "document_Copy.txt", "document.txt"},

		// Number variations
		{"with parenthesis number", "document(1).txt", "document.txt"},
		{"with underscore number", "document_1.txt", "document.txt"},
		{"with dash number", "document-1.txt", "document.txt"},
		{"with space number", "document 1.txt", "document.txt"},

		// Complex cases
		{"mixed case with number", "Document(2).TXT", "document.txt"},
		{"copy with number", "document copy(1).txt", "document.txt"},
		{"copy with underscore", "document_copy_1.txt", "document.txt"},

		// Different extensions
		{"pdf file", "document copy.pdf", "document.pdf"},
		{"no extension", "document copy", "document"},
	}

	// Iterate over each test case
	for _, tt := range tests {
		// Run each test case as a sub-test
		t.Run(tt.name, func(t *testing.T) {
			// Call the GetCanonicalName function
			got := GetCanonicalName(tt.input)

			// Verify the output against the expected value
			if got != tt.expected {
				t.Errorf("GetCanonicalName(%q) = %q, want %q", tt.input, got, tt.expected)
			}
		})
	}
}

// TestContains tests the Contains function
// Each test case specifies a slice of strings, a search string, and the expected output
func TestContains(t *testing.T) {
	tests := []struct {
		name     string   // Test case name
		slice    []string // Slice of strings
		search   string   // Search string
		expected bool     // Expected output
	}{
		// Basic cases
		{"empty slice", []string{}, "search", false},
		{"single element", []string{"search"}, "search", true},
		{"single element mismatch", []string{"search"}, "other", false},
		{"multiple elements", []string{"one", "two", "three"}, "two", true},
		{"multiple elements mismatch", []string{"one", "two", "three"}, "four", false},
	}

	// Iterate over each test case
	for _, tt := range tests {
		// Run each test case as a sub-test
		t.Run(tt.name, func(t *testing.T) {
			// Call the Contains function
			got := Contains(tt.slice, tt.search)

			// Verify the output against the expected value
			if got != tt.expected {
				t.Errorf("Contains(%v, %q) = %t, want %t", tt.slice, tt.search, got, tt.expected)
			}
		})
	}
}
