package stringutils

import (
	"strings"
)

func IsValidSlug(slug string) bool {
	if len(slug) == 0 {
		return false
	}

	allowedChars := "abcdefghijklmnopqrstuvwyz0123456789-_"

	for _, char := range slug {
		if !strings.ContainsRune(allowedChars, char) {
			return false
		}
	}
	return true
}
