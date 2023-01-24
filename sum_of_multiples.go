package main

/*
Given a number, find the sum of all the unique multiples of particular numbers up to but not including that number.

If we list all the natural numbers below 20 that are multiples of 3 or 5, we get 3, 5, 6, 9, 10, 12, 15, and 18.

The sum of these multiples is 78.
*/

func SumMultiples(limit int, divisors ...int) (sum int) {

	for i := 1; i < limit; i++ {
		for _, div := range divisors {
			if div > 0 && i%div == 0 {
				sum += i
				break
			}
		}
	}

	return
}

/*
tarun => ubsvo
Jatin => kbujo
Abdul => Bcevm
*/
