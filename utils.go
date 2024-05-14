package tabber

import (
	"fmt"
	"strconv"
	"strings"
)

// equal checks if two integers are equal.
func equal(v1 int, v2 int) bool {
	return v1 == v2
}

// intToString converts an integer to a string.
func intToString(value int) string {
	return strconv.Itoa(value)
}

// trimmed strips away '"' from a string.
func trimmed(s string) string {
	return strings.Trim(s, "\"")
}

// toSlug returns a copy of string with lowercase
// replacing "_" and whitespaces with "-"
// example: toSlug("New Resource") returns new-resource.
func toSlug(s string) string {
	return strings.Map(func(r rune) rune {
		switch {
		case r == ' ', r == '_':
			return '-'
		case r == '"':
			return -1 // Remove quotes
		default:
			return r
		}
	}, strings.ToLower(trimmed(s)))
}

// makeTabId generates a tab ID string based on the given value.
func makeTabId(value int) string {
	return fmt.Sprintf("tab-%s-btn", strconv.Itoa(value))
}

// makePanelId generates a tab ID string based on the given value.
func makePanelId(value int) string {
	return fmt.Sprintf("tab-%s-panel", strconv.Itoa(value))
}
