package utils

import (
	"strings"
	"unicode/utf8"

	"golang.org/x/text/encoding/charmap"
)

// NormalizeUTF8 returns a UTF-8 safe string by decoding legacy ISO-8859-9
// bytes when needed. This guards against SQL_ASCII/legacy encoded data.
func NormalizeUTF8(value string) string {
	if value == "" {
		return value
	}
	if utf8.ValidString(value) {
		return value
	}

	decoded, err := charmap.ISO8859_9.NewDecoder().String(value)
	if err != nil {
		return strings.ToValidUTF8(value, "")
	}
	if utf8.ValidString(decoded) {
		return decoded
	}
	return strings.ToValidUTF8(decoded, "")
}

// NormalizeUTF8Ptr normalizes a *string safely.
func NormalizeUTF8Ptr(value *string) *string {
	if value == nil {
		return nil
	}
	fixed := NormalizeUTF8(*value)
	return &fixed
}
