package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println("Strings and Runes:")

	text := "übunG macHt dEn meisTer"
	fmt.Println("sentence: ", text)
	fmt.Println("camel-case: ", toCamelCase(text))
	fmt.Println("snake-case: ", toSnakeCase(text))
}

func toCamelCase(sentence string) string {
	words := strings.Fields(sentence)

	for i, word := range words {
		words[i] = toCapitalizedWord(word)
	}

	return strings.Join(words, "")
}

func toSnakeCase(sentence string) string {
	words := strings.Fields(sentence)

	for i, word := range words {
		words[i] = toLowCaseWord(word)
	}

	return strings.Join(words, "-")
}

func toCapitalizedWord(word string) string {
	chars := []rune(word)
	for j, char := range chars {
		if j == 0 {
			chars[j] = unicode.ToTitle(char)
		} else {
			chars[j] = unicode.ToLower(char)
		}
	}
	return string(chars)
}

func toLowCaseWord(word string) string {
	chars := []rune(word)
	for j, char := range chars {
		chars[j] = unicode.ToLower(char)
	}
	return string(chars)
}
