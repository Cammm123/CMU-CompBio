package main

//MakeOverlapNetwork() takes a slice of reads with match, mismatch, gap and a threshold.
//It returns adjacency lists of reads; edges are only included
//in the overlap graph is the alignment score of the two reads is at least the threshold.
func MakeOverlapNetwork(reads []string, match, mismatch, gap, threshold float64) map[string][]string {
	//Initialize an adjacency list to represent the overlap graph.
	adjacencyList := make(map[string][]string)

	return adjacencyList
}
