package util

import (
	"encoding/base64"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"strings"
)

func DecodeInText(text string) string {
	data, err := base64.URLEncoding.DecodeString(text)
	if err != nil {
		return "in text decode failed"
	}

	text = NormalizeText(string(data))

	return text
}

func EncodeOutText(text string) string {
	text = NormalizeText(text)
	return base64.URLEncoding.EncodeToString([]byte(text))
}

// FROM HERE: https://rosettacode.org/wiki/Strip_control_codes_and_extended_characters_from_a_string#Go
//
// Advanced Unicode normalization and filtering,
// see http://blog.golang.org/normalization and
// http://godoc.org/golang.org/x/text/unicode/norm for more
// details.
func NormalizeText(text string) string {
	isOk := func(r rune) bool {
		return r < 32 || r >= 127
	}
	// The isOk filter is such that there is no need to chain to norm.NFC
	t := transform.Chain(norm.NFKD, transform.RemoveFunc(isOk))
	// This Transformer could also trivially be applied as an io.Reader
	// or io.Writer filter to automatically do such filtering when reading
	// or writing data anywhere.
	text, _, _ = transform.String(t, text)
	text = strings.Trim(text, " ")
	return strings.ToLower(text)
}
