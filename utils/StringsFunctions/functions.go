package StringsFunctions

import (
	"math"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

// StringsFunctions.After method returns everything after the given value in a string. The
// entire string will be returned if the value does not exist within the string:
func After(s, search string) string {
	pos := strings.Index(s, search)
	if pos == -1 {
		return s // or return "" if you want similar behavior to Laravel when not found
	}
	return s[pos+len(search):]
}

// The StringsFunctions.Int8CoalescePositive ensure the minimum value is 1 (or def result) is int8:
func Int8CoalescePositive(p *int8, def int8) int8 {
	if p == nil {
		return def
	}
	if *p < 1 {
		return def
	}
	return *p
}

// The StringsFunctions.Int16CoalescePositive ensure the minimum value is 1 (or def result) is int16:
func Int16CoalescePositive(p *int16, def int16) int16 {
	if p == nil {
		return def
	}
	if *p < 1 {
		return def
	}
	return *p
}

// The StringsFunctions.Int32CoalescePositive ensure the minimum value is 1 (or def result) is int32:
func Int32CoalescePositive(p *int32, def int32) int32 {
	if p == nil {
		return def
	}
	if *p < 1 {
		return def
	}
	return *p
}

// The StringsFunctions.Int64CoalescePositive ensure the minimum value is 1 (or def result) is int64:
func Int64CoalescePositive(p *int64, def int64) int64 {
	if p == nil {
		return def
	}
	if *p < 1 {
		return def
	}
	return *p
}

// The StringsFunctions.IntCoalescePositive ensure the minimum value is 1 (or def result) is int:
func IntCoalescePositive(p *int, def int) int {
	if p == nil {
		return def
	}
	if *p < 1 {
		return def
	}
	return *p
}

// The StringsFunctions.IsInteger method whether the given value is of an integer type (signed or unsigned), such as int, int32, int64, uint, etc:
func IsInteger(value interface{}) bool {
	switch reflect.TypeOf(value).Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return true
	default:
		return false
	}
}

// The StringsFunctions.IsSlug method determines whether the given string is a valid Slug:
func IsSlug(s string) bool {
	// IsValidSlug checks if the slug is valid:
	// - only lowercase letters, numbers and hyphens
	// - must not start/end with a hyphen
	// - there cannot be two hyphens in a row
	// Regex: ^[a-z0-9]+(-[a-z0-9]+)*$
	// Meaning: starts with a letter/number, may be followed by repeated letters/numbers
	re := regexp.MustCompile(`^[a-z0-9]+(-[a-z0-9]+)*$`)
	return re.MatchString(s)

}

// The StringsFunctions.Lower method converts the given string to lowercase:
func Lower(s string) string {
	return strings.ToLower(s) //return the result of strings.ToLower(s) back to whatever called the function.
}

// The StringsFunctions.Slug method generates a URL friendly "slug" from the given string:
func Slug(s string, separator rune) string {
	// 1. Remove diacritics (é → e, ö → o, dll.)
	t := transform.Chain(
		norm.NFD,
		runes.Remove(runes.In(unicode.Mn)),
		norm.NFC,
	)
	s, _, _ = transform.String(t, s)

	// 2. Change all non-alnum to separator
	reg := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	s = reg.ReplaceAllString(s, string(separator))

	// 3. Trim separator at the beginning/end
	s = strings.Trim(s, string(separator))

	// 4. Lowercase di akhir
	s = strings.ToLower(s)

	return s
}

// The StringsFunctions.StringPointer method converts the given string into a *string:
func StringPointer(s string) *string {
	// You can extend this function to modify the string before returning,
	// for example using strings.ToLower(s).
	return &s
}

// The StringsFunctions.ToInt8 method converts the given string to int8:
func ToInt8(s string) int8 {
	num, err := strconv.ParseInt(s, 10, 8)
	if err != nil {
		// Return the maximum negative integer for int8
		return math.MinInt8
	}
	return int8(num)
}

// The StringsFunctions.ToInt16 method converts the given string to int16:
func ToInt16(s string) int16 {
	num, err := strconv.ParseInt(s, 10, 16)
	if err != nil {
		// Return the maximum negative integer for int16
		return math.MinInt16
	}
	return int16(num)
}

// The StringsFunctions.ToInt32 method converts the given string to int32:
func ToInt32(s string) int32 {
	num, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		// Return the maximum negative integer for int32
		return math.MinInt32
	}
	return int32(num)
}

// The StringsFunctions.ToInt64 method converts the given string to int64:
func ToInt64(s string) int64 {
	num, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		// Return the maximum negative integer for int64
		return math.MinInt64
	}
	return num
}

// The StringsFunctions.ToInt method converts the given string to int:
func ToInt(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		// Return the maximum negative integer to indicate an error
		return math.MinInt
	}
	// Return the converted integer if no error occurred
	return num
}
