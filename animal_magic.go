package main

import (
	"math/rand"
	"time"
)

/*
Elaine is working on a new children's game that features animals and magic wands. It is time to code functions for rolling a die, generating random wand energy and shuffling a slice.

1. Seed the random number generator.

Implement a SeedWithTime function that seeds the math.rand package with the current computer time.

2. Roll a die.

Implement a RollADie function.

This will be the traditional twenty-sided die with numbers 1 to 20.

d := RollADie() // d will be assigned a random int, 1 <= d <= 20

3. Generate wand energy.

Implement a GenerateWandEnergy function. The wand energy should be a random floating point number between 0.0 and 12.0.

f := GenerateWandEnergy()  // f will be assigned a random float64, 0.0 <= f < 12.0

4. Shuffle a slice.

The game features eight different animals:

ant
beaver
cat
dog
elephant
fox
giraffe
hedgehog
Write a function ShuffleAnimals that returns a slice with the eight animals in random order.

*/

func SeedWithTime() {
	// seed is a starting point for the sequence of pseudorandom number. If you start from the same seed, you get the same sequence. If you want different seed everytime them you can use the current time as a seed

	rand.Seed(time.Now().UnixNano())
}

func RollADie() int {
	var die = make([]int, 20)
	for i := 0; i < len(die); i++ {
		die[i] = i + 1
	}
	// rand.Intn(max_val) - rangee from 0  upto max_val
	return die[rand.Intn(len(die))]
	// die[0] - die[19]
}

func GenerateWandEnergy() float64 { // from 0.0 to 12.0
	// rand.Float64()- 0.0 to 1.0
	return rand.Float64() * 12
}
