package utils

import (
	"regexp"
	"strings"
)

// GetCanonicalName removes common suffixes like "_copy", "(1)", etc.
func GetCanonicalName(name string) string {
	// Remove common patterns for duplicates
	re := regexp.MustCompile(`(?i)(\s*$begin:math:text$copy|\\s*\\(\\d+$end:math:text$|_copy))$`)
	cleanName := re.ReplaceAllString(name, "")
	return strings.TrimSpace(cleanName)
}
