package main

/*
Calculate the Hamming Distance between two DNA strands.

Your body is made up of cells that contain DNA. Those cells regularly wear out and need replacing, which they achieve by dividing into daughter cells. In fact, the average human body experiences about 10 quadrillion cell divisions in a lifetime!

When cells divide, their DNA replicates too. Sometimes during this process mistakes happen and single pieces of DNA get encoded with the incorrect information. If we compare two strands of DNA and count the differences between them we can see how many mistakes occurred. This is known as the "Hamming Distance".

We read DNA using the letters C,A,G and T. Two strands might look like this:

GAGCCTACTAACGGGAT
CATCGTAATGACGGCCT
^ ^ ^  ^ ^    ^^
They have 7 differences, and therefore the Hamming Distance is 7.

The Hamming Distance is useful for lots of things in science, not just biology, so it's a nice phrase to be familiar with :)

Implementation notes
The Hamming distance is only defined for sequences of equal length, so an attempt to calculate it between sequences of different lengths should not work.

You may be wondering about the cases_test.go file. We explain it in the leap exercise.

Source
The Calculating Point Mutations problem at Rosalind
*/

func hamming(s1 string, s2 string) int {

	var num int

	if len(s1) < len(s2) {
		s1, s2 = s2, s1
	}
	/*
	 temp:=s2
	 s2=s1
	 s1=temp
	*/
	for index := range s1 {

		if s1[index] != s2[index] {
			num++
		}

	}

	return num + len(s2) - len(s1)

}
