package main

import (
	"fmt"
)

func main() {
	fmt.Println("Genome assembly!")

	length_random_genome := 1000
	number_kmers := 100

	//WriteContigsToFile(ShuffleStrings(KmerComposition(GenerateRandomGenome(length_random_genome), number_kmers)), "Output/new_string")

	reconstructedGenome := GreedyAssembler(ShuffleStrings(KmerComposition(GenerateRandomGenome(length_random_genome), number_kmers)))

	if StringSliceEquals(KmerComposition(GenerateRandomGenome(length_random_genome), number_kmers), KmerComposition(reconstructedGenome, number_kmers)) {
		fmt.Println("Missiong Acomplished!")
	} else {
		fmt.Println("Oh no.")
	}
}
