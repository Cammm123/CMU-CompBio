package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Strings!")

	fmt.Println(string('A'))
	fmt.Println(string(45))

	fmt.Println(strconv.Itoa(45))
	j, err := strconv.Atoi("23") //err is the second error variable

	// the conversion is OK when the error variable is equal to nil
	if err != nil {
		panic(err)
	}

	fmt.Println(j)

	pi, err2 := strconv.ParseFloat("3.14", 64) //64 means float64 variable
	if err2 != nil {
		panic(err)
	}

	fmt.Println("The value of pi is", pi)
	var s string = "Hi"
	t := "Lovers"
	//concatenate these strings with + operation
	u := s+ " " + t
	fmt.Println(u)
	fmt.Println("The first symbol of u is", string(u[0]))
	fmt.Println("The final symbol of u is", string(u[len(u) - 1]))

	if t[2] == 'v' {
		fmt.Println("The symbol at position 2 of t is v.")
	}

	dna := "ACCGAT"
	fmt.Println(Complement(dna)) // Should print "TGGCTA"
	fmt.Println(Reverse(dna)) //Should print out "TAGCCA"
	fmt.Println(ReverseComplement(dna)) //should print "ATCGGT"
}

/*
ReverseComplement takes as input a string pattern of DNA symbols, 
and returns the reverse compliment of the strings.
*/

func ReverseComplement(pattern string) string {
	return Reverse(Complement(pattern))
}

/* 2 Independent tasks

	1) Complement takes as input a string pattern of DNA symbols. It then returns the string formed
	by somplementing each position of the input string
	(A --> T, C --> G)

	2) Reverse takes as ipnut a string pattern of DNA symbols. It then returns the string formed
	by reversing the positions of all symbols in pattern
*/

func Complement(dna string) string {
	dna2 := make([]byte, len(dna))
	for i, symbol := range dna {
		switch symbol  {
		case 'A': 
			dna2[i] = 'T'
		case 'C':
			dna2[i] = 'G'
		case 'G':
			dna2[i] = 'C'
		case 'T':
			dna2[i] = 'A'
		default:
			panic("Error: Invalid symbol given to the Complement()")
		}
	}
	return string(dna2) //converts slice of bytes to a string
}

func Reverse(pattern string) string {
	rev := make([]byte, len(pattern))
	n := len(pattern)

	for i:= range pattern {
		rev[i] = pattern[n-1-i]
	}

	return string(rev)
}