package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(wordsInCamelCase("jujutsuKaisenItadoriYujji"))
	fmt.Println(caesarCipher("jujutsuKaisenItadoriYujji", 5))
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

// if the letter is not and english alphabet with ascii encoding
// the then letter is not changed
func caesarCipher(str string, shift int32) string {
	r := []rune{}
	for _, v := range str {
		v = rune(v)
		switch {
		case 'a' <= v && v <= 'z':
			r = append(r, rotate(v, 97, shift))
		case 'A' <= v && v <= 'Z':
			r = append(r, rotate(v, 65, shift))
		default:
			r = append(r, v)
		}
	}
	return string(r)
}

func rotate(r rune, startNumber int32, shiftBy int32) rune {
	newNumber := ((r - startNumber + shiftBy) % 26) + startNumber
	return rune(newNumber)
}
