package main

import (
	
)

func DiffuseBoardOneParticle(currentBoard [][]float64, diffusionRate float64, kernel [3][3]float64) [][]float64 {
	newBoard := make([][]float64, len(currentBoard))
	for i := range newBoard {
		newBoard[i] = make([]float64, len(currentBoard[i]))
	}

	for row := range currentBoard {
		for col := range currentBoard[row] {
			newValue := multiplyOut(currentBoard, kernel, row, col)
			newBoard[row][col] = newValue * diffusionRate
		}
	}

	// Calculate the middle indices
	midRow := len(newBoard) / 2
	midCol := len(newBoard[0]) / 2

	// Safely modify the middle cell
	newBoard[midRow][midCol] = 1 + newBoard[midRow][midCol]

	return newBoard
}

func multiplyOut(currentBoard [][]float64, kernel [3][3]float64, row, col int) float64 {
	sum := 0.0
	kernelSize := len(kernel)
	offset := kernelSize / 2

	for i := 0; i < kernelSize; i++ {
		for j := 0; j < kernelSize; j++ {
			newRow := row + i - offset
			newCol := col + j - offset

			if InField(currentBoard, newRow, newCol) {
				sum += currentBoard[newRow][newCol] * kernel[i][j]
			}
		}
	}

	return sum
}

func InField(currentBoard [][]float64, row, col int) bool {
	if row < 0 || col < 0 || row >= len(currentBoard) || col >= len(currentBoard[0]) {
		return false
	}
	return true
}