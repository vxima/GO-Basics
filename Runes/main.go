package main

import (
	"fmt"
	"unicode/utf8"
)

func Application(log string) string {
	for _, v := range log {
		unicode := fmt.Sprintf("%U", v)
		if unicode == "U+2757" {
			return "recommendation"
		} else if unicode == "U+1F50D" {
			return "search"
		} else if unicode == "U+2600" {
			return "weather"
		}
	}
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	newLog := ""
	for _, v := range log {
		if v == oldRune {
			newLog = newLog + string(newRune)
		} else {
			newLog = newLog + string(v)
		}

	}

	return newLog
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	numberOfRunes := utf8.RuneCountInString(log)
	if numberOfRunes <= limit {
		return true
	} else {
		return false
	}
}
func main() {
	fmt.Println(Application("❗ recommended search product 🔍"))
	log := "please replace '👎' with '👍'"

	fmt.Println(Replace(log, '👎', '👍'))
	fmt.Println(WithinLimit("hello❗", 6))

}
