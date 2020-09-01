package main

import (
	"fmt"
	"strings"
	"unicode"
)

func countWordsInCamelCase(str string) int {
	var words = 1
	for _, c := range str {
		if unicode.IsUpper(c) {
			words++
		}
	}
	return words
}

func caesar(str string, offset int32) string {

	var sb strings.Builder

	for _, c := range str {
		if c >= 'A' && c <= 'Z' { // Uppercase
			sb.WriteString(string('A' + (c-'A'+offset)%26))
		} else if c >= 'a' && c <= 'z' { // Lowercase
			sb.WriteString(string('a' + (c-'a'+offset)%26))
		} else {
			sb.WriteRune(c) // Other character
		}
	}

	return sb.String()
}

func main() {
	fmt.Println(countWordsInCamelCase("halloMeinNameIstFred"))
	fmt.Println(caesar("middle-Outz", 2))
}
