package main

//BinarizeMatrix takes a matrix of values and a threshold.
//It binarizes the matrix according to the threshold.
//If entries across main diagonal are both above threshold, only retain the bigger one.
func BinarizeMatrix(mtx [][]float64, threshold float64) [][]int {
    

    
    newArray := make([][]int, len(mtx))
    for a := range newArray {
        newArray[a] = make([]int, len(mtx))
    }
    for i := 0; i < len(newArray); i++ {
        for j := i + 1; j < len(newArray); j++ {
            
            
            
            if mtx[i][j] >= (threshold) {
                newArray[i][j] = 1
                
            }
            if mtx[j][i] >= (threshold) {
                newArray[j][i] = 1
            }
            if newArray[i][j] == newArray[j][i] && newArray[i][j] == 1   {
                if mtx[i][j] >= mtx[j][i] {
           
                    newArray[j][i] = 0
                } else{
                     newArray[i][j] = 0
                } 
                
            }
            
        }
    }
    
  return newArray

}