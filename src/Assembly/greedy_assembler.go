package main

//GreedyAssembler takes a collection of strings and returns a genome whose
//k-mer composition is these strings. It makes the following assumptions.
//1. "Perfect coverage" -- every k-mer is detected
//2. No errors in reads
//3. Every read has equal length (k)
//4. DNA is single-stranded
func GreedyAssembler(reads []string) string {
	if len(reads) == 0 {
		return ""
	}

	reads2 := reads[0]
	reads = reads[1:]

	for len(reads) > 0 {
		found := false

		for i := 0; i < len(reads); i++ {
			read := reads[i]

			// Check if we can append to the left
			if len(reads2) >= 2 && read[len(read)-2:] == reads2[:2] {
				reads2 = read[:1] + reads2
				reads = append(reads[:i], reads[i+1:]...)
				found = true
				break
			}

			// Check if we can append to the right
			if len(reads2) >= 2 && read[:2] == reads2[len(reads2)-2:] {
				reads2 += read[len(read)-1:]
				reads = append(reads[:i], reads[i+1:]...)
				found = true
				break
			}
		}

		if !found { //cant find anything, leave
			break
		}
	}

	return reads2
}