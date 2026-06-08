package main

import (
	"fmt"
	"regexp"
	"strings"
)

func ToCamelCase(input string) string {
	trimmedStr := strings.TrimSpace(input)
	cleanedStr := removeSpecialChar(trimmedStr)
	lowerStr := strings.ToLower(cleanedStr)
	lowerWords := strings.Fields(lowerStr)
	capitalizedWords := capitalizeWordsFromIndex(lowerWords, 1)
	camelCaseStr := strings.Join(capitalizedWords, "")
	return camelCaseStr
}

func removeSpecialChar(s string) string {
	reg := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	return strings.TrimSpace(reg.ReplaceAllString(s, " "))
}

func capitalizeWordsFromIndex(words []string, index int) []string {
	capitalizedWords := make([]string, len(words))
	copy(capitalizedWords, words)
	for i := index; i < len(words); i++ {
		word := words[i]
		if len(word) > 0 {
			capitalizedWords[i] = strings.ToUpper(string(word[0])) + word[1:]
		} else {
			capitalizedWords[i] = word
		}
	}

	return capitalizedWords
}

func main() {
	input := "Hello world 123-1"
	output := ToCamelCase(input)
	fmt.Printf("Input:  %q\n", input)
	fmt.Printf("Output: %q\n", output)
}
