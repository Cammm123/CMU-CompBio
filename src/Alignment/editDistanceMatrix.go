package main

// EditDistanceMatrix takes as input a slice of strings patterns.
// It returns a matrix whose (i,j)th entry is the edit distance between
// the i-th and j-th strings in patterns.
func EditDistanceMatrix(patterns []string) [][]int {
	numRows := len(patterns)
	numCols := len(patterns)

	mtx := InitializeRectangleMatrix(numRows, numCols)

	for i := range patterns {
		for j := i + 1; j < len(patterns); j++ {
			mtx[i][j] = EditDistance(patterns[i], patterns[j])
			mtx[j][i] = mtx[i][j]
		}
	}

	return mtx
}
