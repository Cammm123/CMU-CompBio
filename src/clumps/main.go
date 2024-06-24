package main

import (
	"net/http" //for accessing url
	"fmt"
	"io" //needed to read from and write to files
	"log" //needed for log files (errors)
	"time"
)

func main() {
	fmt.Println("Clumps!")

	url := "https://bioinformaticsalgorithms.com/data/realdatasets/Replication/E_coli.txt"
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	
	//close the connections when you're done
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Recieved non-OK status: %v", resp.Status)
	}

	genomeSymbols, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	
	//we now have a slice of symbols associated with the genome

	fmt.Println("The number of nucleotides in the E-Coli genome is", len(genomeSymbols))

	EcoliGenome := string(genomeSymbols)

	k := 9
	L := 500
	t := 3

	clumps := FindClumps(EcoliGenome, k, L, t)

	start := time.Now()
	FindClumps(EcoliGenome, k, L, t)
	elapsed := time.Since(start)
	log.Printf("FindClumps took %s", elapsed)

	start2 := time.Now()
	FindClumpsFaster(EcoliGenome, k, L, t)
	elapsed2 := time.Since(start2)
	log.Printf("FindClumpsFaster took %s", elapsed2)
	
	fmt.Println("Found", len(clumps), "total patterns occuring as clumps.")
}

/* FindClumps() takes as input a string text, and integers k, L, t.
	It then returns a slice of strings representing all substrings of length k
	that appear at least t times in a "window" of length L in the string text.
*/

//text = "BANANASPLIT"
//L = 6
//first window: BANANA
//second window: ANANAS

func FindClumpsFaster(text string, k, L, t int) [] string {
	patterns := make([]string, 0)
	n := len(text)
	foundPatterns := make(map[string]bool)

	firstWindow := text[:L]
	freqMap := FrequencyTable(firstWindow, k)

	//append any patterns we find to patterns slice
	for s, freq := range freqMap {
		if freq >= t {
			patterns = append(patterns, s)
			foundPatterns[s] = true
		}
	}

	for i:=1; i<=n-L; i++ {
		//decrease by one the value associated with the frist subtring of length k in the preceding window
		oldPattern := text[i-1:i-1+k]
		freqMap[oldPattern] --

		//clean up the map if the frequency of oldPattern was one
		if freqMap[oldPattern] == 0 {
			delete(freqMap, oldPattern)
		}
		
		//add the pattern from the end of the current window
		newPattern := text[i+L-k:i+L]
		freqMap[newPattern]++

		//I have updated the frequency map :)
		for s, freq := range freqMap {
			if freq >= t && !foundPatterns[s] {
				patterns = append(patterns, s)
				foundPatterns[s] = true
			}
		}
	}

	return patterns
}

func FindClumps(text string, k, L, t int) []string {
	patterns := make([]string, 0)
	n := len(text)

	/*
	//map to track whether I have added a string to patterns yet
	foundPatterns := make(map[string]bool)

	//I can use this line of code instead if I want to cut out the Contains() function.
	*/
	for i:=0; i<=n-L; i++ {
		//set the current window
		window := text[i:i+L]

		//let's make the frequency table for this window
		freqMap := FrequencyTable(window, k)
		
		//what occurs frequency
		for s,val := range freqMap {
			//append s to patterns if s occurs frequently and s doesn't already appear in patterns.
			if val >= t && !Contains(patterns, s){
				patterns = append(patterns, s)
			}
		}
	}
	return patterns
}

/* Contains() will take as input a slice of strings and a single string.
	It then returns true if s is contained in the slice, and false otherwise
*/
func Contains(patterns []string, s string) bool {
	for _, pattern := range patterns {
		if pattern == s {
			return true
		}
	}
	return false //otherwise if it is never true, then its false
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
		freqMap[pattern]++
	}

	return freqMap
}