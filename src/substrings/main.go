package main

import (
	"fmt"
)

func main() {
	fmt.Println("Substrings (and slices)!")

	s := "Hi Lovers!"

	fmt.Println(s[1:5]) //Going from 1 index to 5 index (not including 5)
	fmt.Println(s[0:7])
	fmt.Println(s[4:len(s)]) //Not len(s) - 1 because we dont include the ending in string.
	//In practice, you dont need len(s) or 0 beacuse that is assumed

	a := make([]int, 6)
	for i := range a {
		a[i] = 2*i + 1
	}
	fmt.Println(a)
	fmt.Println(a[3:5])
	fmt.Println(a[:3])
	fmt.Println(a[4:])

	pattern := "ATA"
	text := "ATATA"

	fmt.Println(PatternCount(pattern, text))
	fmt.Println(StartingIndicies(pattern, text))
}

/* PatternCount() takes as input two strings: pattern and text.
	This will then return the number of times pattern occurs in a substring of text (including overlaps)
*/

func PatternCount(pattern, text string) int {
	return len(StartingIndices(pattern,text))
	/*
	count := 0
	k := len(pattern)
	n := len(text)
	for i := 0; i<=n-k; i++ {
		if pattern == text[i:i+k] {
			count ++
		}
	}

	return count
	*/
}

/*StartingIndices() takes as input two string: pattern and text
	It then returns the collection of all starting positions of pattern as a substring of text, 
	including all overlaps
*/

func StartingIndices(pattern, text string) []int {
	positions := make([]int, 0)
	n := len(text)
	k := len(pattern)

	//range over text, appending current position to positions if we find a math
	for i := 0; i <= n-k; i++ {
		if pattern == text[i:i+k] {
			positions = append(positions, i)
		}
	}

	return positions
}