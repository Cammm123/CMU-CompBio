package main

import (	
	"math/rand"
)

//ShuffleStrings takes a collection of strings patterns as input.
//It returns a random shuffle of the strings.
func ShuffleStrings(patterns []string) []string {
    k := len(patterns)
    stringArray := make([]string, k)
    numArray := rand.Perm(k) //numbers 0 --> k in a random order


    for i := range stringArray {
       stringArray[numArray[i]] = patterns[i] //set patterns[i] to a random index in stringArray
    }
    
	return stringArray
}
