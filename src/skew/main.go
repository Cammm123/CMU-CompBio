package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"gonum.org/v1/plt"
	"gonum.org/v1/plt/plotter" //include line graph support
	"gonum.org/v1/plt/vg" //support for vector graphics and image generation
)

/* SkewArray() takes as input a DNA string genome
	It then returns a slice of integers corresponding to the G-C Skew of the genome at each position
*/
func SkewArray(genome string) []int {
	n := len(genome)
	array := make([]int, n+1)
	array[0] = 0 //redundant, but we want that

	//range over remaining values and set i-th value of array 
	for i := 1; i <= n; i++ { //starting at one because [0] nedes to be = 0
		array[i] = array[i-1] + Skew(genome[i-1])
	}

	return array
}

/* Skew() takes a symbol as an input.
	It then outputs 1) if symbol is G, -1) if symbol is C, and 0) otherwise.
*/
func Skew(symbol byte) int {
	if symbol == 'G' {
		return 1
	} else if symbol == 'C' {
		return -1
	} else {
		return 0
	}
}

/* MinimumSkew() --> this is were we get to that origin
	This takes as input a DNA string genome.
	It then returns a slice of integers corresponding to the indices in the genome
	where the skew array attains a minimum value.
*/
func MinimumSkew(genome string) []int {
	indices := make([]int, 0)

	array := SkewArray(genome)

	m := MinIntigerArray(array)

	for i, val := range array {
		if val == m {
			//found an index
			indices = append(indices, i)
		}
	}

	return indices
}

func MinIntigerArray(list []int) int {
	if len(list) == 0 {
		panic("Error: empty list.")
	}

	min := list[0]
	for i := range list {
		if list[i] < min {
			min = list[i]
		}
	}
	return min
}

/* MakeSkewDiagram() takes as input a skew array.
	It then outputs nothing. It draws the skew diagram of the given skew array to an image and saves to file.
*/

func MakeSkewDiagram(skewArray []int) {
	p := plot.New() //creates new plotter object

	p.Title.Text = "Skew Diagram"
	p.X.Label.Text = "Genome Position"
	p.Y.Label.Text = "Skew Value"

	//remove legend
	p.Legend.Top = false

	//make a collection of points associated with each skew value
	points := make(plotter.XYs, len(skewArray))

	//set x and Y valuye of each point
	for i, val := range skewArray {
		points[i].X = float64(i)
		points[i].Y = float64(skewValue)
	}

	//connect the dots!
	line, err := plotter.NewLine(points)
	if err != nil {
		panic(err)
	}

	//add our line to the plot
	p.Add(line)

	//draw to an image
	//1) Set a unit of length
	unitOfLength := vg.Centimeter

	p.X.Label.TextStyle.Font.Size = 3 * unitOfLength
	p.Y.Label.TextStyle.Font.Size = 3 * unitOfLength
	p.Title.TextStyle.Font.Size = 4 * unitOfLength
	p.X.Tick.Label.Font.Size = 2 * unitOfLength
	p.Y.Tick.Label.Font.Size = 2 * unitOfLength

	//2) Save plot to png
	err = p.Save(100*unitOfLength, 60*unitOfLength, "skewDiagram.png")
	if err != nil {
		panic(err)
	}
	
}



func main() {
	fmt.Println("The Skew Array!")

	url := "https://bioinformaticsalgorithms.com/data/realdatasets/Replication/E_coli.txt"
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	
	//close the connections when you're done
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Recieved non-OK status: %v", resp.Status)
	}

	genomeSymbols, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	genome := string(genomeSymbols)
	fmt.Println("Genome read. It has, ", len(genome), "total nucleotides.")

	EcoliSkewArray := SkewArray(genome)
	minSkewPositoins := MinimumSkew(genome)

	fmt.Println("The minimum skew value of", EcoliSkewArray[minSkewPositoins[0]], "occurs at positions", minSkewPositoins)

	//drawing the picture
	MakeSkewDiagram(EcoliSkewArray)
	fmt.Println("Skew Diagram drawn! Exiting normally.")
}

