package StringsFunctions

import "strings"

func After(s, search string) string {
	pos := strings.Index(s, search)
	if pos == -1 {
		return s // or return "" if you want similar behavior to Laravel when not found
	}
	return s[pos+len(search):]
}

func Lower(s string) string {
	return strings.ToLower(s) //return the result of strings.ToLower(s) back to whatever called the function.
}
