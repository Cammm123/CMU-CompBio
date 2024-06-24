package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Homework 2!")
}

/* -------------------------------7.1---------------------------------*/

// GameBoard is a two-dimensional slice of booleans
type GameBoard1 []([]bool)

// write your InField() function here along with any subroutines that you need. CountRows and CountCols have been provided for your convenience.
func InField(board GameBoard1, i, j int) bool {
	// Check if i and j are within the bounds of the board
	if i < 0 || j < 0 || i >= CountRows(board) || j >= CountCols(board) {
		return false
	}

	// Return the value at the specified cell
	return board[i][j]
}

func CountRows(board GameBoard1) int {
	return len(board)
}

func CountCols(board GameBoard1) int {
	// assume that we have a rectangular board
	if CountRows(board) == 0 {
		panic("Error: empty board given to CountCols")
	}
	// give # of elements in 0-th row
	return len(board[0])
}

func CountLiveNeighbors(board GameBoard1, r, c int) int {
	start := InField(board, r, c)
	if start { //if the starting spot is on the board, continue
		//check north side
		north := NorthSide(board, r, c)

		//check south side
		south := SouthSide(board, r, c)

		//check sides
		side := Beside(board, r, c)

		return north + south + side
	}

	panic("Error: the starting piece is not on the board")
}

func NorthSide(board GameBoard1, r, c int) int { //counts on the upper row
	add_count := 0
	newY := c - 1
	//newCord := (r, newY)

	if InField(board, r, newY) && board[r][newY] { // if its infield and true
		add_count++
	}

	add_count += Beside(board, r, newY)

	return add_count
}

func SouthSide(board GameBoard1, r, c int) int { //counts on either side of the starting place
	add_count := 0
	newY := c + 1
	//newCord := (r, newY)

	if InField(board, r, newY) && board[r][newY] { // if its infield and true
		add_count++
	}

	add_count += Beside(board, r, newY)

	return add_count
}

func Beside(board GameBoard1, r, c int) int { //counts on the lower row
	add_count := 0

	new1x := r + 1
	new2x := r - 1

	if InField(board, new1x, c) && board[new1x][c] {
		add_count++
	}
	if InField(board, new2x, c) && board[new2x][c] {
		add_count++
	}

	return add_count
}

/*
Game of Life Instructions

A: If a cell is alive and has either two or three live neighbors, then it remains alive.
B: If a cell is alive and has zero or one live neighbors, then it dies out.
C: If a cell is alive and has four or more live neighbors, then it dies out.
D: If a cell is dead and has more than or fewer than three live neighbors, then it remains dead.
E: If a cell is dead and has exactly three live neighbors, then it becomes alive.
*/

/*
// PlayGameOfLife takes an initial GameBoard and the number of generations to simulate
func PlayGameOfLife(initialBoard GameBoard1, numGens int) []GameBoard1 {
	// Create a slice to store the boards for each generation
	new_board := make([]GameBoard1, numGens+1)

	// Store the initial board as the first element
	new_board[0] = initialBoard

	// Iterate through the number of generations
	for i := 1; i <= numGens; i++ {
		// Update the board for the next generation
		new_board[i] = UpdateBoard(new_board[i-1])
	}

	return new_board
}
*/

// InitializeBoard takes a number of rows and columns as inputs and returns a gameboard with appropriate number of rows and colums, where all values = false.
func InitializeBoard(numRows, numCols int) GameBoard1 {
	// make a 2-D slice (default values = false)
	var board GameBoard1
	board = make(GameBoard1, numRows)
	// now we need to make the rows too
	for r := range board {
		board[r] = make([]bool, numCols)
	}

	return board
}

/*
func UpdateBoard(currBoard GameBoard1) GameBoard1 {
	// Create a new board to store the updated state
	newBoard := make(GameBoard1, CountRows(currBoard))
	for i := range newBoard {
		newBoard[i] = make([]bool, CountCols(currBoard))
	}

	// Iterate through each cell in the current board
	for r := range currBoard {
		for c := range currBoard[r] {
			newBoard[r][c] = UpdateCell(currBoard, r, c)
		}
	}

	return newBoard
}
*/

/*
func UpdateCell(board GameBoard1, r, c int) bool {
	if A(board, r, c) || E(board, r, c) { //if this qualifies for A or E, become alive
		return true
	} else { //technically I only need to write A and E, because every other case is feaf
		return false
	}
}
*/

func A(board GameBoard1, r, c int) bool {
	if InField(board, r, c) && board[r][c] { //must start alive and in field
		if CountLiveNeighbors(board, r, c) == 2 || CountLiveNeighbors(board, r, c) == 3 {
			return true
		}

		return false //otherwise, return false
	}

	return false
}

func E(board GameBoard1, r, c int) bool {
	if InField(board, r, c) { //must start alive and in field
		if !board[r][c] && CountLiveNeighbors(board, r, c) == 3 {
			return true
		}

		return false //otherwise, return false
	}

	return false
}

/* -------------------------------7.2---------------------------------*/
//going to return 0 if the value I get is unuvailable

type GameBoard2 [][]int

func NeighborhoodToString(board GameBoard2, row, col int, neighborhoodType string) string {
	var answer string
	if neighborhoodType == "Moore" {
		answer = Moore(board, row, col)
	} else if neighborhoodType == "vonNeumann" {
		answer = vonNeumann(board, row, col)
	}
	return answer
}

func Moore(board GameBoard2, row, col int) string {
	var answer string
	answer += strconv.Itoa(board[row][col])

	// Define the relative positions in clockwise order starting from top-left
	directions := [][]int{
		{-1, -1}, // Top-left
		{-1, 0},  // Top
		{-1, 1},  // Top-right
		{0, 1},   // Right
		{1, 1},   // Bottom-right
		{1, 0},   // Bottom
		{1, -1},  // Bottom-left
		{0, -1},  // Left
	}

	for _, dir := range directions {
		r := row + dir[0]
		c := col + dir[1]
		if isValidCell(board, r, c) {
			answer += strconv.Itoa(board[r][c])
		} else {
			answer += "0" // Append "0" for out-of-bounds cells
		}
	}

	return answer
}

func vonNeumann(board GameBoard2, row, col int) string {
	var answer string

	// Center
	answer += strconv.Itoa(board[row][col])

	// Top
	if row-1 >= 0 {
		answer += strconv.Itoa(board[row-1][col])
	} else {
		answer += "0"
	}

	// Right
	if col+1 < len(board[0]) {
		answer += strconv.Itoa(board[row][col+1])
	} else {
		answer += "0"
	}

	// Bottom
	if row+1 < len(board) {
		answer += strconv.Itoa(board[row+1][col])
	} else {
		answer += "0"
	}

	// Left
	if col-1 >= 0 {
		answer += strconv.Itoa(board[row][col-1])
	} else {
		answer += "0"
	}

	return answer
}

func isValidCell(board GameBoard2, row, col int) bool {
	return row >= 0 && row < len(board) && col >= 0 && col < len(board[0])
}

func UpdateCell(board GameBoard2, row, col int, neighborhoodType string, rules map[string]int) int {
	var neighborhood string
	if neighborhoodType == "Moore" {
		neighborhood = Moore(board, row, col)
	} else if neighborhoodType == "vonNeumann" {
		neighborhood = vonNeumann(board, row, col)
	}

	if newState, found := rules[neighborhood]; found {
		return newState
	}
	return 0
}

func UpdateBoard(board GameBoard2, neighborhoodType string, rules map[string]int) GameBoard2 {
	rows := len(board)
	cols := len(board[0])
	newBoard := make(GameBoard2, rows)
	for i := range newBoard {
		newBoard[i] = make([]int, cols)
	}

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			newBoard[row][col] = UpdateCell(board, row, col, neighborhoodType, rules)
		}
	}

	return newBoard
}

func PlayAutomaton(numGens int, neighborhoodType string, initialBoard GameBoard2, ruleStrings map[string]int) []GameBoard2 {
	var boards []GameBoard2
	currentBoard := initialBoard

	for gen := 0; gen < numGens; gen++ {
		boards = append(boards, currentBoard)
		currentBoard = UpdateBoard(currentBoard, neighborhoodType, ruleStrings)
	}
	boards = append(boards, currentBoard)
	return boards
}
