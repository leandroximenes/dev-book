package utils

import (
	"regexp"
	"strings"
	"unicode"

	// "golang.org/x/text/unicode/norm"
)

// Slugify converts a string into a URL-safe slug
func Slugify(s string) string {
	// Normalize (removes accents like é → e, ç → c, ñ → n)
	// t := norm.NFD.String(s)

	t := s
	sb := strings.Builder{}
	for _, r := range t {
		if unicode.Is(unicode.Mn, r) {
			continue // skip marks
		}
		sb.WriteRune(r)
	}
	s = sb.String()

	// Lowercase
	s = strings.ToLower(s)

	// Replace non-alphanumeric with hyphen
	reg := regexp.MustCompile(`[^a-z0-9]+`)
	s = reg.ReplaceAllString(s, "-")

	// Trim hyphens
	s = strings.Trim(s, "-")

	return s
}
