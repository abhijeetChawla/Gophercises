package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(wordsInCamelCase("jujutsuKaisenItadoriYujji"))
}

func wordsInCamelCase(s string) int32 {
	if s == "" {
		return 0
	}
	// this is one since if the string is not empty we will have al least one word
	var count int32 = 1
	for _, word := range s[1:] {
		if unicode.IsUpper(word) {
			count++
		}
	}
	return count
}
