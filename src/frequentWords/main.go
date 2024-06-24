package main

import (
	"fmt"
)

func main() {
	fmt.Println("Finding Frequent Words!")

	text := "ATCAATGATCAACGTAAGCTTCTAAGCATGATCAAGGTGCTCACACAGTTTATCCACAACCTGAGTGGATGACATCAAGATAGGTCGTTGTATCTCCTTCCTCTCGTACTCTCATGACCACGGAAAGATGATCAAGAGAGGATGATTTCTTGGCCATATCGCAATGAATACTTGTGACTTGTGCTTCCAATTGACATCTTCAGCGCCATATTGCGCTGGCCAAGGTGACGGAGCGGGATTACGAAAGCATGATCATGGCTGTTGTTCTGTTTATCTTGTTTTGACTGAGACTTGTTAGGATAGACGGTTTTTCATCACTGACTAGCCAAAGCCTTACTCTGCCTGACATCGACCGTAAATTGATAATGAATTTACATGCTTCCGCGACGATTTACCTCTTGATCATCGATCCGATTGAAGATCTTCAATTGTTAATTCTCTTGCCTCGACTCATAGCCATGATGAGCTCTTGATCATGTTTCCTTAACCCTCTATTTTTTACGGAAGAATGATCAAGCTGCTGCTCTTGATCATCGTTTC"

	k := 9

	freqMap := FrequencyTable(text, k)
	freqPatterns := FindFrequentWords(text, k)
	fmt.Println("Frequent patterns found when k is", k, "are", freqPatterns)

	fmt.Println("The maximum number of occurrences is", freqMap[freqPatterns[0]])
}

/*
FindFrequentWords takes as input a string tect and an integer k.
	It then reteurns a slice of strings corresponding to the substring(s) of length k
	that occur most frequently in text.
*/

func FindFrequentWords(text string, k int) []string {
	freqPatterns := make([]string, 0)
	freqMap := FrequencyTable(text, k)
	//once I have the frequency table, find the maximum value
	max := MaxMapValue(freqMap)
	//what keys of this table reach the max value?
	for pattern, val := range freqMap {
		if val == max {
			freqPatterns = append(freqPatterns, pattern)
		}
	}

	return freqPatterns
}

/* FrequencyTable() takes as input a string tet and an integer k.
	It then returns the frequency table here mapping each substring
	of the text of length k to its number of occurences.
*/
func FrequencyTable(text string, k int) map[string]int {
	freqMap := make(map[string]int)

	n:=len(text)
	for i:=0; i<=n-k; i++ {
		pattern := text[i:i+k]
		
		/*
		//2 cases. 1) Already exists --> add one, 2) Doesn't exist --> make new value
		_, exists := freqMap[pattern]
		if exists {
			freqMap[pattern] ++
		} else {
			freqMap[pattern] = 1
		}
		*/
		//Shortway to do that:

		freqMap[pattern]++
	}

	return freqMap
}


/* MaxMapValue() takes as input a map of strings to integers.
	Then, it returns the maximum value found in the map.
*/
func MaxMapValue(dictionary map[string]int) int {
	m := 0
	firstTime := true

	for _, val := range dictionary {
		if firstTime || val > m {
			m = val
			firstTime = false
		}
	}

	return m
}