package main

import (
	"fmt"
)

func main() {
	fmt.Println("Metagenomics!")

	Process2023Samples()
}

func Process2023Samples() {
	// parse the tax_report file and construct a taxonomy data reference map
	fileName := "Data/2023_Samples/Kraken2TaxonomicClassifications/tax_report.txt"
	taxonomyData := ReadTaxonomyMapFromFile(fileName)

	// see how many percentage of our data have classification with rank larger than 5
	//rank := 20

	//let's read all the samples.
	fmt.Println("Next, let's read in all of our samples.")
	path := "Data/2023_Samples/Kraken2TaxonomicClassifications/Classification/"
	allMaps := ReadSamplesFromDirectory(path)

	// if we are satisfied with the analysis we can start the analysis
	// First we have to range through every frequency map and preserve the data that reached this rank
	fmt.Println("Let's trim our data to the desired depth of lineage.")
	allMapsWithRank := UpdateRankDeepest(allMaps, taxonomyData)

	samplingThreshold := 1000

	for key, freqMap := range allMapsWithRank {
		if SampleTotal(freqMap) < samplingThreshold {
			//print that we are deleting this map
			fmt.Println(key, "has been removed from the set of frequency maps as its total number of elements, including repeats, is smaller than the threshold", samplingThreshold)
			//delete this map from consideration
			delete(allMapsWithRank, key)
		}
	}

	//now, down sample allMapsWithRank
	downsampledMaps := DownSampleMaps(allMapsWithRank, samplingThreshold)

	fmt.Println("Forming a map of richness maps.")

	richness := RichnessMap(allMaps)

	// writing richness maps to file

	richnessFile := "Matrices/RichnessMatrix_2023.csv"

	WriteRichnessMapToFile(richness, richnessFile)

	fmt.Println("Forming a map of Simpson's indices.")

	simpson := SimpsonsMap(downsampledMaps)

	fmt.Println("Simpson's index map formed!")

	fmt.Println("Now it's time for Beta-Diversity analysis, starting with Bray-Curtis.")

	distMetric := "Bray-Curtis"
	sampleNames, mtxBC := BetaDiversityMatrix(downsampledMaps, distMetric)

	simpsonFile := "Matrices/SimpsonMatrix_2023.csv"
	WriteSimpsonsMapToFile(simpson, simpsonFile)

	brayCurtisFile := "Matrices/BrayCurtisBetaDiversityMatrix_2023.csv"
	WriteBetaDiversityMatrixToFile(mtxBC, sampleNames, brayCurtisFile)

	fmt.Println("Bray-Curtis matrix formed! Now for Jaccard.")
	distMetric = "Jaccard"
	sampleNames2, mtxJac := BetaDiversityMatrix(downsampledMaps, distMetric)
	jaccardFile := "Matrices/JaccardBetaDiversityMatrix_2023.csv"
	//write to file
	WriteBetaDiversityMatrixToFile(mtxJac, sampleNames2, jaccardFile)

	fmt.Println("Success! Now we are ready to visualize our data. :)")
}

func Process2019Samples() {
	//let's check that we can read in one map

	filename := "Data/2019_Samples/Fall_Allegheny_1.txt"

	freqMap := ReadFreqMapFromFile(filename)

	fmt.Println("File read successfully! We have", len(freqMap), "total elements in our map.")

	//let's print its Simpson's index.

	fmt.Println("Simpson's index:", SimpsonsIndex(freqMap))

	//let's read in all the samples.
	fmt.Println("Reading in all samples.")

	path := "Data/2019_Samples/"
	allMaps := ReadSamplesFromDirectory(path)

	fmt.Println("Samples read successfully! We have", len(allMaps), "total samples.")

	richness := RichnessMap(allMaps)

	fmt.Println("Let's print richness of each sample.")

	for sampleName, val := range allMaps {
		fmt.Println(sampleName, val)
	}

	fmt.Println(len(richness))

	// writing richness maps to file

	richnessFile := "Matrices/RichnessMatrix_2019.csv"

	WriteRichnessMapToFile(richness, richnessFile)

	simpson := SimpsonsMap(allMaps) // will store Simpson's Index for all samples

	//this is all great, but let's write all our maps and matrices to files.
	simpsonFile := "Matrices/SimpsonMatrix_2019.csv"

	WriteSimpsonsMapToFile(simpson, simpsonFile)

	fmt.Println("Simpson's map successfully written to file.")

	// let's do two beta diversity matrices too.

	distMetric := "Bray-Curtis"

	sampleNames, mtxBC := BetaDiversityMatrix(allMaps, distMetric)

	fmt.Println("Writing Bray-Curtis matrix to file.")

	brayCurtisFile := "Matrices/BrayCurtisBetaDiversityMatrix_2019.csv"

	WriteBetaDiversityMatrixToFile(mtxBC, sampleNames, brayCurtisFile)

	fmt.Println("Bray-Curtis matrix written to file! Let's make Jaccard matrix.")

	distMetric = "Jaccard"
	sampleNames2, mtxJac := BetaDiversityMatrix(allMaps, distMetric)

	jaccardFile := "Matrices/JaccardBetaDiversityMatrix_2019.csv"

	//write to file
	WriteBetaDiversityMatrixToFile(mtxJac, sampleNames2, jaccardFile)

	fmt.Println("Success! Now we are ready to visualize our data.")
}
