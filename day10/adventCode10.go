package main

import (
	"fmt"
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
	lines := strings.Split(strings.TrimSpace(string(content)), "\n")

	for _, line := range lines {
		var row []int
		for _, char := range strings.TrimSpace(line) {
			num, err := strconv.Atoi(string(char))
			if err != nil {
				return nil, fmt.Errorf("error parsing digit: %v", err)
			}
			row = append(row, num)
		}
		result = append(result, row)
	}

	return result, nil
}

// Find the score of a trailhead using DFS
func findScore(matrix [][]int, i, j int, reached [][]bool) int {
	if matrix[i][j] == 9 {
		if !reached[i][j] {
			reached[i][j] = true
			return 1
		}
		return 0
	}

	directions := [][2]int{
		{-1, 0}, // up
		{1, 0},  // down
		{0, -1}, // left
		{0, 1},  // right
	}

	score := 0
	for _, dir := range directions {
		ni, nj := i+dir[0], j+dir[1]
		if ni >= 0 && ni < len(matrix) && nj >= 0 && nj < len(matrix[0]) &&
			matrix[ni][nj] == matrix[i][j]+1 {
			score += findScore(matrix, ni, nj, reached)
		}
	}

	return score
}

func findScorePartTwo(matrix [][]int, i, j int) int {
	if matrix[i][j] == 9 {
		return 1
	}

	directions := [][2]int{
		{-1, 0}, // up
		{1, 0},  // down
		{0, -1}, // left
		{0, 1},  // right
	}

	score := 0
	for _, dir := range directions {
		ni, nj := i+dir[0], j+dir[1]
		if ni >= 0 && ni < len(matrix) && nj >= 0 && nj < len(matrix[0]) &&
			matrix[ni][nj] == matrix[i][j]+1 {
			score += findScorePartTwo(matrix, ni, nj)
		}
	}

	return score
}

func calculateTotalScore(matrix [][]int) int {
	reached := make([][]bool, len(matrix))
	for i := range reached {
		reached[i] = make([]bool, len(matrix[0]))
	}

	totalScore := 0
	for i, row := range matrix {
		for j, val := range row {
			if val == 0 {
				for i := range reached {
					reached[i] = make([]bool, len(matrix[0]))
				}
				score := findScore(matrix, i, j, reached)
				fmt.Printf("Trailhead at (%d, %d) has a score of %d\n", i, j, score)
				totalScore += score
			}
		}
	}

	return totalScore
}

func calculateTotalScorePartTwo(matrix [][]int) int {
	totalScore := 0
	for i, row := range matrix {
		for j, val := range row {
			if val == 0 {
				score := findScorePartTwo(matrix, i, j)
				fmt.Printf("Trailhead at (%d, %d) has a score of %d\n", i, j, score)
				totalScore += score
			}
		}
	}

	return totalScore
}

func main() {
	data, err := parseFile("input10.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	totalScore := calculateTotalScore(data)
	totalScore2 := calculateTotalScorePartTwo(data)
	fmt.Println("Total Score of All Trailheads:", totalScore)
	fmt.Println("Total Score of All Trailheads 2:", totalScore2)
}
