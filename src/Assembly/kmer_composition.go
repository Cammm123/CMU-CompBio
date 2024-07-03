package main

//KmerComposition returns the k-mer composition (all k-mer substrings) of a given genome.
func KmerComposition(genome string, k int) []string {
    kmers := make([]string, k)

    for i := range kmers {
        kmers[i] = genome[i:i+k]
    }
   return kmers
}
