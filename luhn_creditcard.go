package main

import (
	"strings"
	"unicode"
)

/*
Instructions
Given a number determine whether or not it is valid per the Luhn formula.

The Luhn algorithm is a simple checksum formula used to validate a variety of identification numbers, such as credit card numbers and Canadian Social Insurance Numbers.

The task is to check if a given string is valid.

Validating a Number
Strings of length 1 or less are not valid. Spaces are allowed in the input, but they should be stripped before checking. All other non-digit characters are disallowed.

Example 1: valid credit card number
4539 3195 0343 6467
The first step of the Luhn algorithm is to double every second digit, starting from the right. We will be doubling

4_3_ 3_9_ 0_4_ 6_6_
If doubling the number results in a number greater than 9 then subtract 9 from the product. The results of our doubling:

8569 6195 0383 3437
Then sum all of the digits:

8+5+6+9+6+1+9+5+0+3+8+3+3+4+3+7 = 80
If the sum is evenly divisible by 10, then the number is valid. This number is valid!

Example 2: invalid credit card number
8273 1232 7352 0569
Double the second digits, starting from the right

7253 2262 5312 0539
Sum the digits

7+2+5+3+2+2+6+2+5+3+1+2+0+5+3+9 = 57
57 is not evenly divisible by 10, so this number is not valid.
*/

func modify(num int) int {
	if (num * 2) > 9 {
		return num*2 - 9
	} else {
		return num * 2
	}
}

func IsValid(number string) bool {
	/*
		 string should be more than 1 in length
						 loop from the last digit towards the start:
						  if it is the even digit from the end:
						    double it
			          if it becomes greater than 9
			            sub 9 from the product

						    and add to count
						  else:
						    add to the count

					      if count %10 == 0 : return true
					      else return false
	*/
	//   '2'-> int('2')- int('0')
	numWithoutSpaces := strings.ReplaceAll(number, " ", "")

	if len(numWithoutSpaces) <= 1 {
		return false
	}
	// 234567
	count := 1
	sum := 0
	for i := len(numWithoutSpaces) - 1; i >= 0; i-- {
		digit := rune(numWithoutSpaces[i])
		if !unicode.IsDigit(digit) {
			return false
		}
		if count%2 == 0 {
			sum += modify(int(digit))
		} else {
			sum += int(digit)
		}
		// 7+3+5+8+3+4
		count++
	}

	return sum%10 == 0
}
