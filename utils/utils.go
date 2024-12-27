package utils

import (
	"path/filepath"
	"regexp"
	"strings"
)

var (
	duplicatePatterns = []*regexp.Regexp{
		regexp.MustCompile(`(?i)(_copy|[\s-]copy)`), // matches _copy, -copy, space copy (case insensitive)
		regexp.MustCompile(`\s*\(\d+\)$`),           // matches (1), (2)
		regexp.MustCompile(`\s*_\d+$`),              // matches _1, _2
		regexp.MustCompile(`\s*-\d+$`),              // matches -1, -2
		regexp.MustCompile(`\s+\d+$`),               // matches " 1", " 2"
	}
)

func GetCanonicalName(filePath string) string {
	baseName := filepath.Base(filePath)

	ext := filepath.Ext(baseName)
	nameWithoutExt := strings.TrimSuffix(baseName, ext)

	// Apply patterns
	for _, pattern := range duplicatePatterns {
		nameWithoutExt = pattern.ReplaceAllString(nameWithoutExt, "")
	}

	// Clean up any remaining whitespace
	nameWithoutExt = strings.TrimSpace(nameWithoutExt)

	return strings.ToLower(nameWithoutExt) + strings.ToLower(ext)
}

func Contains(slice []string, str string) bool {
	for _, v := range slice {
		if v == str {
			return true
		}
	}
	return false
}
