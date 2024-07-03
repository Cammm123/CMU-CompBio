package main

import (	
	"math/rand"
)

//GenerateRandomGenome takes a parameter length and returns
//a random DNA string of this length where each nucleotide has equal probability.
func GenerateRandomGenome(n int) string {
	random_string := make([]byte, n)

	for i:=0; i<n; i++ {

		int := rand.Intn(4)

		switch int {
		case 0:
			random_string[i] = 'A'
		case 1:
			random_string[i] = 'T'
		case 2:
			random_string[i] = 'C'
		case 3:
			random_string[i] = 'G'
		}

		/*
		if int == 0 {
			random_string += "A"
		} else if int == 1 {
			random_string += "T"
		} else if int == 2 {
			random_string += "C"
		} else if int == 3 {
			random_string += "G"
		} else {
			panic("Error: There has been an error with generating a random number")
		}
		*/
		
	}

	return string(random_string)
}