package main

// CountSharedKmers takes two strings and an integer k. It returns the number of
// k-mers that are shared by the two strings.
func CountSharedKmers(str1, str2 string, k int) int {
	kmers1 := make(map[string]int)
	kmers2 := make(map[string]int)

	for i := 0; i < len(str1)-k+1; i++ {
		kmers1[str1[i:i+k]]++
	}

	for j := 0; j < len(str2)-k+1; j++ {
		kmers2[str2[j:j+k]]++
	}

	return SumOfMinima(kmers1, kmers2)
}
