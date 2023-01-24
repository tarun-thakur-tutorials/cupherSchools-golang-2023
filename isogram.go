package main

import (
	"strings"
)

/*
Determine if a word or phrase is an isogram.

An isogram (also known as a "non-pattern word") is a word or phrase without a repeating letter, however spaces and hyphens are allowed to appear multiple times.

Examples of isograms:

lumberjacks
background
downstream
six-year-old

The word isograms, however, is not an isogram, because the s repeats.
*/

func IsIsogram(s string) bool {

	s = strings.ToLower(s)

	for index, character := range s {
		if index == len(s)-1 {
			break
		}

		if string(character) == " " || string(character) == "-" {
			continue
		}
		// func to check if the character is a alphabetical letter or not: unicode.IsLetter(character)
		if strings.ContainsRune(s[index+1:], character) {
			return false
		}
	}
	return true
}
