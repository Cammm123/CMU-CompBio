package main

// SimpsonsMap takes a map mapping sample IDs to frequency maps.package solution
// It returns a map of sample IDs to Simpson's indices for each sample.
func SimpsonsMap(allMaps map[string](map[string]int)) map[string]float64 {
	new_map := make(map[string]float64)

	for i := range allMaps {
		// if allMaps[i] != new_map[i] {

		// }
		new_map[i] = SimpsonsIndex(allMaps[i])
	}

	return new_map
}
