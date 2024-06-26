package main

import "math/rand"

// DownSampleMaps() takes a map of frequency maps and a threshold, and returns a map of frequency maps with the same keys, but with each frequency map randomly downsampled to the threshold.
func DownSampleMaps(allMaps map[string]map[string]int, threshold int) map[string]map[string]int {
	//make a new map of frequency maps
	newMaps := make(map[string]map[string]int)

	//for each frequency map in allMaps
	for key, freqMap := range allMaps {
		//downsample the frequency map to the threshold
		newMap := DownSample(freqMap, threshold)
		//add the downsampled frequency map to newMaps
		newMaps[key] = newMap
	}

	//return newMaps
	return newMaps
}

// DownSample() takes as input a frequency map and a threshold, and returns a new frequency map with the same keys, but with each value randomly downsampled to the threshold.
func DownSample(freqMap map[string]int, threshold int) map[string]int {
	if SampleTotal(freqMap) < threshold {
		//panic
		panic("DownSample() called on a frequency map with a total value less than the threshold!")
	}
	//make a new map of ints
	newMap := make(map[string]int)

	// randomly generate threshold total values from freqMap, with replacement
	for i := 0; i < threshold; i++ {
		//pick a random key from freqMap
		key := RandomKey(freqMap)
		//add one to the value of that key in newMap
		newMap[key]++
	}

	//return newMap
	return newMap
}

// RandomKey takes as input a frequency map and returns a random key from that map.
// The probability of each key being chosen is proportional to its value in the frequency map.
func RandomKey(freqMap map[string]int) string {
	//pick a random number between 0 and the total value of freqMap
	rand := rand.Intn(SampleTotal(freqMap))
	//start a counter at 0
	counter := 0
	//for each key in freqMap
	for key, val := range freqMap {
		//add val to counter
		counter += val
		//if counter is greater than rand
		if counter > rand {
			//return key
			return key
		}
	}
	//if we get here, something went wrong
	panic("RandomKey() failed!")
}
