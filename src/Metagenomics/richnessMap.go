package main

// RichnessMap takes a map of frequency maps as input.  It returns a map
// whose values are the richness of each sample.
func RichnessMap(allMaps map[string](map[string]int)) map[string]int {
	new_map := make(map[string]int)

	for sampleName := range allMaps {
		new_map[sampleName] = Richness(allMaps[sampleName])
	}

	return new_map
}
