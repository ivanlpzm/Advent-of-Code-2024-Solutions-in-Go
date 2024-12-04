package main

import (
	"fmt"
	"os"
	"strings"
)

func parseFile(filename string) ([][]byte, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var result [][]byte
	lines := strings.Split(string(content), "\n")

	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}
		parts := strings.Fields(line)

		var row []byte
		for _, part := range parts {
			row = append(row, []byte(part)...)
		}

		result = append(result, row)
	}
	return result, nil
}

func checkCandidate(i, j, deltaI, deltaJ int, matrix [][]byte) bool {
	return matrix[i+deltaI][j+deltaJ] == 'M' &&
		matrix[i+2*deltaI][j+2*deltaJ] == 'A' &&
		matrix[i+3*deltaI][j+3*deltaJ] == 'S'
}

func checkAllDirections(i, j, rowLen, colLen int, matrix [][]byte) int {
	result := 0

	directions := []struct {
		deltaI, deltaJ int
	}{
		{-1, 0}, // UP
		{1, 0},  // DOWN
		{0, -1}, // LEFT
		{0, 1},  // RIGHT
		{-1, -1}, // LEFT UP
		{1, -1},  // LEFT DOWN
		{-1, 1},  // RIGHT UP
		{1, 1},   // RIGHT DOWN
	}

	for _, dir := range directions {
		newI, newJ := i+3*dir.deltaI, j+3*dir.deltaJ
		if newI >= 0 && newI < rowLen && newJ >= 0 && newJ < colLen && checkCandidate(i, j, dir.deltaI, dir.deltaJ, matrix) {
			result++
		}
	}

	return result
}

func checkX(i, j, rowLen, colLen int, matrix [][]byte) bool {
	if i <= 0 || j <= 0 || i >= rowLen-1 || j >= colLen-1 {
		return false
	}

	upLeft, downRight := matrix[i-1][j-1], matrix[i+1][j+1]
	upRight, downLeft := matrix[i-1][j+1], matrix[i+1][j-1]

	return (upLeft == 'M' && downRight == 'S' && downLeft == 'M' && upRight == 'S') ||
		(upLeft == 'S' && downRight == 'M' && downLeft == 'S' && upRight == 'M') ||
		(upLeft == 'M' && downRight == 'S' && upRight == 'M' && downLeft == 'S') ||
		(downLeft == 'M' && downRight == 'M' && upRight == 'S' && upLeft == 'S')
}

func calculateNumberXMAS(matrix [][]byte) int {
	sum := 0
	for i, row := range matrix {
		for j, val := range row {
			if val == 'A' && checkX(i, j, len(matrix), len(row), matrix){ // ASCII value of 'A'
				sum++
			}
		}
	}
	return sum
}

func calculateNumberMAS(matrix [][]byte) int {
	sum := 0
	for i, row := range matrix {
		for j, val := range row {
			if val == 'X' { 
				sum += checkAllDirections(i, j, len(matrix), len(row), matrix)
			}
		}
	}
	return sum
}

func main() {
	input, err := parseFile("input4.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	result := calculateNumberMAS(input)
	result2 := calculateNumberXMAS(input)
	fmt.Println("Result:", result)
	fmt.Println("Result 2:", result2)
}
