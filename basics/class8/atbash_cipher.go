package main

import (
	"strings"
	"unicode"
)

/*
Create an implementation of the atbash cipher, an ancient encryption system created in the Middle East.

The Atbash cipher is a simple substitution cipher that relies on transposing all the letters in the alphabet such that the resulting alphabet is backwards. The first letter is replaced with the last letter, the second with the second-last, and so on.

An Atbash cipher for the Latin alphabet would be as follows:

Plain:  abcdefghijklmnopqrstuvwxyz
Cipher: zyxwvutsrqponmlkjihgfedcba
It is a very weak cipher because it only has one possible key, and it is a simple mono-alphabetic substitution cipher. However, this may not have been an issue in the cipher's time.

Ciphertext is written out in groups of fixed length, the traditional group size being 5 letters, leaving numbers unchanged, and punctuation is excluded. This is to make it harder to guess things based on word boundaries.

Examples
Encoding test gives gvhg
Encoding x123 yes gives c123b vh
Decoding gvhg gives test
Decoding gsvjf rxpyi ldmul cqfnk hlevi gsvoz abwlt gives thequickbrownfoxjumpsoverthelazydog

*/

func AtbashCipherEncode(input string) (encodedOutput string) {
	/*
	   Plain:  abcde fghij klmno pqrst uvwxy z
	   Cipher: zyxwv utsrq ponml kjihg fedcb a

	   Encoding x123 yes gives c123b vh

	   Decoding gsvjf rxpyi ldmul cqfnk hlevi gsvoz abwlt gives thequickbrownfoxjumpsoverthelazydog
	*/
	myMap := map[rune]rune{
		'a': 'z',
		'b': 'y',
		'c': 'x',
		'd': 'w',
		'e': 'v',
		'f': 'u',
		'g': 't',
		'h': 's',
		'i': 'r',
		'j': 'q',
		'k': 'p',
		'l': 'o',
		'm': 'n',
		'n': 'm',
		'o': 'l',
		'p': 'k',
		'q': 'j',
		'r': 'i',
		's': 'h',
		't': 'g',
		'u': 'f',
		'v': 'e',
		'w': 'd',
		'x': 'c',
		'y': 'b',
		'z': 'a',
		'A': 'Z',
		'B': 'Y',
		'C': 'X',
		'D': 'W',
		'E': 'V',
		'F': 'U',
		'G': 'T',
		'H': 'S',
		'I': 'R',
		'J': 'Q',
		'K': 'P',
		'L': 'O',
		'M': 'N',
		'N': 'M',
		'O': 'L',
		'P': 'K',
		'Q': 'J',
		'R': 'I',
		'S': 'H',
		'T': 'G',
		'U': 'F',
		'V': 'E',
		'W': 'D',
		'X': 'C',
		'Y': 'B',
		'Z': 'A',
	}
	count := 0
	for _, val := range input {
		// run data type
		if unicode.IsSpace(val) {
			continue // ignore
		}
		if count%5 == 0 {
			encodedOutput += " "
		}
		if unicode.IsLetter(val) {
			encodedOutput += string(myMap[val])
		} else {
			encodedOutput += string(val) // add as it is
		}
		count++
	}

	return
}

// 'c' => string ( int('c')- int('a'))

func Atbash(s string) string {
	var shiftedStr string
	i := 0
	for _, r := range s {
		base := 0
		skipSpace := false
		switch {
		case r >= 'a' && r <= 'z':
			base = int('a')
		case r >= 'A' && r <= 'Z':
			base = int('A')
		case r >= '0' && r <= '9':
		default:
			skipSpace = true
		}
		if base != 0 {
			// base= int('a') , r='z',
			// 25 - (25)=0
			// r:= run(base + shift) => a + 0 =>a
			// base='a' , r='y', expected='b'
			// 25 - (int('b')-int('a'))=> 25-1=>24
			// r:= rune (int('a') + 24)=> y
			shift := 25 - (int(r) - base)
			r = rune(base + shift)
		}
		if !skipSpace {
			if i%5 == 0 && i > 0 {
				shiftedStr += " "
			}
			shiftedStr += string(r)
			i++
		}
	}
	return strings.ToLower(shiftedStr)
}
