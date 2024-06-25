package main

import (
	"sort"
)

// BetaDiversityMatrix takes a map of frequency maps along with a distance metric
// ("Bray-Curtis" or "Jaccard") as input.
// It returns a slice of strings corresponding to the sorted names of the keys
// in the map, along with a matrix of distances whose (i,j)-th
// element is the distance between the i-th and j-th samples using the input metric.
// Input: a collection of frequency maps samples and a distance metric
// Output: a list of sample names and a distance matrix where D_i,j is the distance between
// sample_i and sample_j according to the given distance metric
func BetaDiversityMatrix(allMaps map[string](map[string]int), distanceMetric string) ([]string, [][]float64) {

	//make the matrix
	new2D_map := make([][]float64, len(allMaps))
	for row := range new2D_map {
		new2D_map[row] = make([]float64, len(allMaps))
	}

	//get sample names
	list_of_names := make([]string, 0, len(allMaps))
	for i := range allMaps {
		list_of_names = append(list_of_names, i)
	}
	sort.Strings(list_of_names)

	//set elements
	for i := range new2D_map {
		for j := range new2D_map[i] {

			id1 := list_of_names[i]
			id2 := list_of_names[j]

			map1 := allMaps[id1]
			map2 := allMaps[id2]
			if distanceMetric == "Bray-Curtis" {
				// these are the params: map1 map[string]int, map2 map[string]int) float64
				new2D_map[i][j] = BrayCurtisDistance(map1, map2)
			} else if distanceMetric == "Jaccard" {
				new2D_map[i][j] = JaccardDistance(map1, map2)
			}
		}
	}

	//return values
	return list_of_names, new2D_map
}

/*

Philip Version:

func BetaDiversityMatrixSolutionCode(allMaps map[string](map[string]int), distanceMetric string) ([]string, [][]float64) {
	numSamples := len(allMaps)
	sampleNames := make([]string, 0, numSamples)
	 //grab all sample names
	 for name := range allMaps {
		samplenames = append(sampleNames, name)
	 }

	 //sort strings
	 sort.Strings(sampleNames)

	 //form distance matrix

	 matx := InitializeSquareMatrix(numSamples)

	 //set elements
	 for r:= 0; r < numSamples; r++
	 {
		for c:= r+1; c< len(mtx[r]); c++
		{
			id1 := sampleNames[r]
			id2:= sampleNames[c]

			map1 := allMaps[id1]
			map2 := allMaps[id2]
			if distanceMetric == "Bray-Curtis" {
				mtx[r][c] = BrayCurtisDistance(map1, map2)
			} else if distanceMetric == "Jaccard" {
				mtx[r][c] = JaccardDistance(map1, map2)
			}

			//regardless of metric, set symmetric
			mtx[c][r] = mtx[r][c]
		}
	 }
}

*/
