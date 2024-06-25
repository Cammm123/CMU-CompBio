package main

// RichnessMap takes a map of frequency maps as input.  It returns a map
// whose values are the richness of each sample.
func RichnessMap(allMaps map[string](map[string]int)) map[string]int {
	new_map := make(map[string]int)

	for i := range allMaps {
		new_map[i] = Richness(allMaps[i])
	}

	return new_map
}
