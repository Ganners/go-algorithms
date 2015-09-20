package main

import (
	"fmt"
	"log"
	"math"
)

func main() {

	matrixSize := 6
	matrix := createMatrix(matrixSize)
	log.Println(printMatrix(matrix))

	matrix = rotateMatrix(matrix)
	log.Println(printMatrix(matrix))
}

func printMatrix(matrix [][]int) string {

	outString := "\n"

	for y := 0; y < len(matrix); y++ {
		for x := 0; x < len(matrix[y]); x++ {
			outString += fmt.Sprintf(" %2d ", matrix[y][x])
		}
		outString += "\n"
	}

	return outString
}

func rotateMatrix(matrix [][]int) [][]int {

	size := len(matrix[0])

	for layer := 0; layer < int(math.Floor(float64(size/2))); layer++ {
		for i := layer; i < size-1-layer; i++ {

			temp := matrix[layer][i]
			matrix[layer][i] = matrix[i][size-1-layer]
			matrix[i][size-1-layer] = matrix[size-1-layer][size-1-i]
			matrix[size-1-layer][size-1-i] = matrix[size-1-i][layer]
			matrix[size-1-i][layer] = temp
		}
	}

	return matrix
}

func createMatrix(size int) [][]int {

	matrix := make([][]int, size)

	for i := 0; i < size*size; i++ {
		row := int(math.Floor(float64(i / size)))
		col := i % size
		if len(matrix[row]) == 0 {
			matrix[row] = make([]int, size)
		}
		matrix[row][col] = i
	}

	return matrix
}
