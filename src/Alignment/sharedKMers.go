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

func SumOfMinima(map1 map[string]int, map2 map[string]int) int {
	c := 0

	for key := range map1 {
		_, exists := map2[key]
		if exists {
			c += Min2(map1[key], map2[key])
		}
	}

	return c
}

// Min2 takes two integers and returns their minimum.
func Min2(n1, n2 int) int {
	if n1 < n2 {
		return n1
	}
	return n2
}
