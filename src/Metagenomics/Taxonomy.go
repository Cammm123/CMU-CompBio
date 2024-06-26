package main

// CheckSamplePercentage takes in a frequency map of the species, int rank and a map of taxonomyData
// It returns the percentage of classification in the freqMap
func CheckSamplePercentage(freqMap map[string]int, rank int, taxonomyData map[string]taxonomy) float64 {
	totalCnt := 0
	passCnt := 0
	// range through the frequency map and check all the keys
	for key, val := range freqMap {
		totalCnt += val
		// check how long is the lineage
		if len(taxonomyData[key].lineage) > rank {
			passCnt += val
		}
	}
	return float64(passCnt) / float64(totalCnt)
}

// UpdateRankDeepest takes as input a map of frequency maps and a map of strings to taxonomy structs.
// It returns a new map in which each frequency map has been updated to the deepest rank.
func UpdateRankDeepest(allMaps map[string](map[string]int), taxonomyData map[string]taxonomy) map[string](map[string]int) {
	// make a new map
	rankedAllMap := make(map[string](map[string]int))

	for location, freqMap := range allMaps {
		// make a new updated map
		rankedFreqMap := make(map[string]int)

		// range through all the keys in each map
		for taxID, count := range freqMap {
			lin := taxonomyData[taxID].lineage
			// preserve the data and update the taxID to the deepest taxonomy rank
			if len(lin) > 0 {
				rankedFreqMap[lin[0]] = count
			}
		}
		// append the map into the new all map
		rankedAllMap[location] = rankedFreqMap
	}
	return rankedAllMap
}

// UpdateRankAllMap Update all the data points to the desired depth of lineage
func UpdateRankAllMap(allMaps map[string](map[string]int), rank int, taxonomyData map[string]taxonomy) map[string](map[string]int) {
	// make a new map
	rankedAllMap := make(map[string](map[string]int))

	for location, freqMap := range allMaps {
		// make a new updated map
		rankedFreqMap := make(map[string]int)

		// range through all the keys in each map
		for taxID, count := range freqMap {
			lin := taxonomyData[taxID].lineage
			// check if the lineage is long enough
			if len(lin) > rank {
				// preserve the data and update the taxID to the desire taxonomy rank
				rankedFreqMap[lin[len(lin)-rank]] += count
			}
		}
		// append the map into the new all map
		rankedAllMap[location] = rankedFreqMap
	}
	return rankedAllMap
}
