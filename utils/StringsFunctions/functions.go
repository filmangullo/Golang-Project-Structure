package StringsFunctions

import (
	"math"
	"strconv"
	"strings"
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

// The StringsFunctions.Lower method converts the given string to lowercase:
func Lower(s string) string {
	return strings.ToLower(s) //return the result of strings.ToLower(s) back to whatever called the function.
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
