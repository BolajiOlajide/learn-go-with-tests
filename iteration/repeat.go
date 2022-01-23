package iteration

import "strings"

// Repeat repeat a letter / word n number of times
func Repeat(letter string, times int) string {
	return strings.Repeat(letter, times)
}
