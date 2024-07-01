package main

import (
	"fmt"
)

func main() {
	fmt.Println("Sequence alignment!")
	SARSAlignment()

}

func Hemoglobin() {
	zebrafish := ReadFASTAFile("Data/Hemoglobin/Danio_rerio_hemoglobin.fasta")
	cow := ReadFASTAFile("Data/Hemoglobin/Bos_taurus_hemoglobin.fasta")
	gorilla := ReadFASTAFile("Data/Hemoglobin/Gorilla_gorilla_hemoglobin.fasta")
	human := ReadFASTAFile("Data/Hemoglobin/Homo_sapiens_hemoglobin.fasta")

	match := 1.0
	mismatch := 1.0
	gap := 3.0 //insertion or deletion of an amino acid is more deadly in terms of mutations

	alignment1 := GlobalAlignment(zebrafish, human, match, mismatch, gap)
	WriteAlignmentToFASTA(alignment1, "Output/fish-human.txt")

	alignment2 := GlobalAlignment(cow, human, match, mismatch, gap)
	WriteAlignmentToFASTA(alignment2, "Output/cow-human.txt")

	alignment3 := GlobalAlignment(gorilla, human, match, mismatch, gap)
	WriteAlignmentToFASTA(alignment3, "Output/gorilla-human.txt")
}

func SARSAlignment() {
	fmt.Println("Reading in genomes.")
	sars1 := ReadFASTAFile("Data/Coronaviruses/SARS-CoV_genome.fasta")
	sars2 := ReadFASTAFile("Data/Coronaviruses/SARS-CoV-2_genome.fasta")
	fmt.Println("Genomes read, Running global aligment")

	match := 1.0
	mismatch := 1.0
	gap := 3.0
	sarsAlignment := GlobalAlignment(sars1, sars2, match, mismatch, gap)
	fmt.Println("Global alignment ran. Printing alignment to file.")
	WriteAlignmentToFASTA(sarsAlignment, "Output/sars_genome_alignment.txt")
}
