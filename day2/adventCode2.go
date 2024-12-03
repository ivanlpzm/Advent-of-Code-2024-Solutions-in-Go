package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func parseFile(filename string) ([][]int, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var result [][]int
	lines := strings.Split(string(content), "\n")

	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}
		parts := strings.Fields(line)

		var row []int
		for _, part := range parts {
			num, err := strconv.Atoi(part)
			if err != nil {
				return nil, fmt.Errorf("error parsing number: %v", err)
			}
			row = append(row, num)
		}

		result = append(result, row)
	}
	return result, nil
}

func isRowSafe(row []int) bool {
	isIncreasing, isDecreasing := false, false

	for i := 0; i < len(row)-1; i++ {
		distance := math.Abs(float64(row[i] - row[i+1]))

		if distance > 3 || distance < 1 {
			return false
		}

		if row[i] < row[i+1] {
			isIncreasing = true
		} else if row[i] > row[i+1] {
			isDecreasing = true
		}

		if isIncreasing && isDecreasing {
			return false
		}
	}

	return true
}

func isRowSafeWithDampener(row []int) bool {
	if isRowSafe(row) {
		return true
	}

	for i := 0; i < len(row); i++ {
		newRow := append([]int{}, row[:i]...)
		newRow = append(newRow, row[i+1:]...)
		if isRowSafe(newRow) {
			return true
		}
	}

	return false
}

func calculateSafeRows(matrix [][]int) int {
	count := 0
	for _, row := range matrix {
		if isRowSafe(row) {
			count++
		}
	}
	return count
}

func calculateSafeRowsWithDampener(matrix [][]int) int {
	count := 0
	for _, row := range matrix {
		if isRowSafeWithDampener(row) {
			count++
		}
	}
	return count
}

func main() {
	matrix, err := parseFile("input2.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	
	safeRows := calculateSafeRows(matrix)
	safeRowsWithDampener := calculateSafeRowsWithDampener(matrix)
	fmt.Println("Safe rows without dampener:", safeRows)
	fmt.Println("Safe rows with dampener:", safeRowsWithDampener)
}
