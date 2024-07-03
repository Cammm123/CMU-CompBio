package main

//OverlapScoringMatrix takes a collection of reads along with alignment penalties.
//It returns a matrix in which mtx[i][j] is the overlap alignment score of
//reads[i] with reads[j].
func OverlapScoringMatrix(reads []string, match, mismatch, gap float64) [][]int {
    numRows := len(reads)
    numCols := len(reads)
	mtx := InitializeRectangleMatrix(numRows, numCols)
    for i := 0; i < len(reads); i++ {
		for j := i + 1 ; j < len(reads); j++ {
            if i == j {
                mtx[i][j] = 0
            } else {
            mtx[i][j] = int(ScoreOverlapAlignment(reads[i], reads[j], match, mismatch, gap))
            mtx[j][i] = int(ScoreOverlapAlignment(reads[j], reads[i], match, mismatch, gap))
            }
		}
	}
    return mtx
}
func InitializeRectangleMatrix(a, b int) [][]int {
	//first, make the slice
	mtx := make([][]int, a)
	for i := 0; i < a; i++ {
		// in here, we make the rows
		mtx[i] = make([]int, b)
	}
	return mtx
}
