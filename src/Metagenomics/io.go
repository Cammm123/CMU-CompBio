package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
)

// taxonomy contains a name (taxonomic ID) and a lineage, consisting of successive taxonomic IDs from the current ID up to the root of the tree of life.
type taxonomy struct {
	taxName string
	lineage []string
}

// WriteRichnessMapToFile writes the richness data from a map to a CSV file.
// Input: A map containing sample names as keys and richness values as integers, and a filename string for the output file.
// Output: None. The function writes directly to a file and panics on failure.
func WriteRichnessMapToFile(richness map[string]int, filename string) {
	file, err := os.Create(filename)
	if err != nil { // panic if anything went wrong
		panic(err)
	}

	writer := bufio.NewWriter(file)

	//print headers
	fmt.Fprint(writer, "Sample")
	fmt.Fprint(writer, ",")
	fmt.Fprint(writer, "Richness")
	fmt.Fprintln(writer, "")

	// print sample name and value on each line
	for sampleName, val := range richness {
		fmt.Fprint(writer, sampleName)
		fmt.Fprint(writer, ",")
		fmt.Fprint(writer, val)

		//print new line
		fmt.Fprintln(writer, "")
	}

	writer.Flush()

	file.Close() // the "defer" statement says "do this at the end of the file"
}

// WriteSimpsonsMapToFile writes the Simpson's diversity index data from a map to a CSV file.
// Input: A map containing sample names as keys and Simpson's index values as floats, and a filename string for the output file.
// Output: None. The function writes directly to a file and panics on failure.
func WriteSimpsonsMapToFile(simpson map[string]float64, filename string) {
	file, err := os.Create(filename)
	if err != nil { // panic if anything went wrong
		panic(err)
	}

	writer := bufio.NewWriter(file)

	//print headers
	fmt.Fprint(writer, "Sample")
	fmt.Fprint(writer, ",")
	fmt.Fprint(writer, "SimpsonsIndex")
	fmt.Fprintln(writer, "")

	// print sample name and value on each line
	for sampleName, val := range simpson {
		fmt.Fprint(writer, sampleName)
		fmt.Fprint(writer, ",")
		fmt.Fprint(writer, val)

		//print new line
		fmt.Fprintln(writer, "")
	}

	writer.Flush()

	file.Close() // the "defer" statement says "do this at the end of the file"
}

// WriteBetaDiversityMatrixToFile writes a matrix of beta diversity values and their corresponding sample names to a CSV file.
// Input: A 2D slice of float64 representing the beta diversity matrix, a slice of strings for sample names, and a filename string for the output file.
// Output: None. The function writes directly to a file and panics on failure.
func WriteBetaDiversityMatrixToFile(mtx [][]float64, sampleNames []string, filename string) {
	file, err := os.Create(filename)
	if err != nil { // panic if anything went wrong
		panic(err)
	}

	writer := bufio.NewWriter(file)
	// add gap at start of file
	fmt.Fprint(writer, ",")

	//print all sample names
	for _, name := range sampleNames {
		fmt.Fprint(writer, name)
		fmt.Fprint(writer, ",")
	}
	fmt.Fprintln(writer, "")

	for i := range mtx {
		fmt.Fprint(writer, sampleNames[i])
		fmt.Fprint(writer, ",")
		for j := range mtx[i] {
			fmt.Fprint(writer, mtx[i][j])
			fmt.Fprint(writer, ",")
		}
		fmt.Fprintln(writer, "")
	}

	writer.Flush()

	file.Close() // the "defer" statement says "do this at the end of the file"

}

// ReadSamplesFromDirectory has the following input and output
// Input: a collection of filenames. Each file has one string on each line
// Output: a map whose keys are the sample names and whose values are the frequency map at that site.
func ReadSamplesFromDirectory(directory string) map[string](map[string]int) {

	allMaps := make(map[string](map[string]int))

	dirContents, err := ioutil.ReadDir(directory)
	if err != nil {
		panic("Error reading directory!")
	}

	for _, fileData := range dirContents {
		// what is the file name?
		fileName := fileData.Name()

		// Remove the file extension (".txt") to obtain the name of the sample
		sampleName := strings.Replace(fileName, ".txt", "", 1)

		freqMap := ReadFreqMapFromFile(filepath.Join(directory, fileName))

		allMaps[sampleName] = freqMap
	}

	return allMaps
}

// ReadFreqMapFromFile
// Input: name of file which contains one string on each line
// Output: a frequency map corresponding to the strings in the file
func ReadFreqMapFromFile(filename string) map[string]int {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	patterns := strings.Split(strings.Replace(strings.TrimSpace(string(data)), "\r\n", "\n", -1), "\n")
	freqMap := make(map[string]int)
	for _, val := range patterns {
		freqMap[val]++
	}
	return freqMap
}

// ReadTaxonomyMapFromFile(fileName)
// Input: name of file which contains one string on each line. Each line contains code, taxID, taxName, and lineage
// Output: a map of taxonomy for all the entry in the file. The key is taxID and the value is taxonomy which contains taxName and taxRank
// Noted that taxRank is converted from the lineage of each entry
func ReadTaxonomyMapFromFile(fileName string) map[string]taxonomy {
	taxonomyMap := make(map[string]taxonomy)

	// read in the file
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.Replace(strings.TrimSpace(string(data)), "\r\n", "\n", -1), "\n")

	for _, line := range lines[1:] { // TODO: run through the whole dataset
		data := strings.Split(line, "|")
		// 0 is code
		code := strings.Replace(data[0], "\t", "", -1)
		// if the code is 3 it's not associated with public entry
		// if the code is 4, it's not in the database. We will skip them here
		// fmt.Println(code)
		if code == "3" || code == "4" {
			continue
		}

		// 1 is taxID, we use it as key
		taxID := strings.Replace(data[1], "\t", "", -1)

		var t taxonomy
		// 2 is taxName
		t.taxName = strings.Replace(data[2], "\t", "", -1)
		// fmt.Println(t.taxName)

		// 3 is taxRank calculated through the length of it's lineage
		lineageList := strings.Fields(data[3])

		t.lineage = lineageList

		// // range through the list and convert the strings to integer
		// for i, s := range lineageList {
		// 	t.lineage[i], _ = strconv.Atoi(s)
		// }
		// assign the entry to a key
		taxonomyMap[taxID] = t
	}
	return taxonomyMap
}

// ReadSamplesFromDirectoryWithSampling has the following input and output
// Input: a collection of filenames. Each file has one string on each line and a map which contains taxonomy data
// Output: a map whose keys are the sample names and whose values are the frequency map at that site.
func ReadSamplesFromDirectoryWithSampling(directory string, taxonomyData map[string]taxonomy, threshold int, valid bool) map[string](map[string]int) {

	allMaps := make(map[string](map[string]int))

	dirContents, err := ioutil.ReadDir(directory)
	if err != nil {
		panic("Error reading directory!")
	}

	for _, fileData := range dirContents {
		// what is the file name?
		fileName := fileData.Name()

		// Remove the file extension (".txt") to obtain the name of the sample
		sampleName := strings.Replace(fileName, ".txt", "", 1)

		freqMap := ReadFreqMapFromFileWithSampling(filepath.Join(directory, fileName), taxonomyData, threshold, valid)

		allMaps[sampleName] = freqMap
	}

	return allMaps
}

// ReadFreqMapFromFileWithSampling
// Input: name of file which contains one string on each line, a taxonomyData map, and a sample threshold and a boolean valid
// Output: it randomly sample from the file and if valid equals true it would only preserve species that are present in taxonomy data
func ReadFreqMapFromFileWithSampling(filename string, taxonomyData map[string]taxonomy, threshold int, valid bool) map[string]int {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	patterns := strings.Split(strings.Replace(strings.TrimSpace(string(data)), "\r\n", "\n", -1), "\n")
	freqMap := make(map[string]int)

	// return an empty map if the number of samples doesn't meet the threshold
	if len(patterns) < threshold {
		return freqMap
	}

	// reading in valid samples
	if valid == true {
		var validPatterns []string
		for _, val := range patterns {
			// check each pattern, it it's valid (in taxonomyData) we read it in
			_, check := taxonomyData[val]
			if check {
				validPatterns = append(validPatterns, val)
			}
		}
		fmt.Println("There are", len(validPatterns), "samples in", filename)
		// if size of the valid pattern is smaller than threshold return an empty list
		if len(validPatterns) < threshold {
			fmt.Println(filename, "doesn't have enough valid sample!")
			return freqMap
		} else { // if there's enough samples
			// shuffle the pattern and sample threshold number of samples.
			rand.Shuffle(len(validPatterns), func(i, j int) { validPatterns[i], validPatterns[j] = validPatterns[j], validPatterns[i] })
			for _, val := range validPatterns[0:threshold] {
				freqMap[val]++
			}
		}

	} else {
		for _, val := range patterns[0:threshold] {
			freqMap[val]++
		}
	}

	return freqMap
}
