package main

import (
	"fmt"
	"os"
	"strings"
)

type Pair struct {
	X, Y int
}

func parseFile(filename string) ([][]rune, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(strings.TrimSpace(string(content)), "\n")
	var result [][]rune
	for _, line := range lines {
		if line = strings.TrimSpace(line); line == "" {
			continue
		}
		parts := strings.Fields(line)
		var row []rune
		for _, part := range parts {
			row = append(row, []rune(part)...)
		}
		result = append(result, row)
	}
	return result, nil
}

func tryLoop(matrix [][]rune, ni, nj int, initDir int) bool {
	dx := []int{-1, 0, 1, 0}
	dy := []int{0, 1, 0, -1}
	mi, mj := ni, nj
	limit := 2 * (len(matrix) + len(matrix[0]))
	count := 0
	dir := initDir
	for count < limit {
		mi, mj = mi+dx[initDir], mj+dy[initDir]

		if mi < 0 || mi >= len(matrix) || mj < 0 || mj >= len(matrix[mi]) {
			break
		}

		if mi == ni && mj == nj {
			return true
		}

		if matrix[mi][mj] == '#' {
			dir = (dir + 1) % 4

		}

		count++
	}

	return mi == ni && mj == nj
}

func calculatePositionsVisited(matrix [][]rune) (int, int) {
	var startI, startJ int
	found := false
	for i := 0; i < len(matrix) && !found; i++ {
		for j := 0; j < len(matrix[i]) && !found; j++ {
			if matrix[i][j] == '^' {
				startI, startJ = i, j
				found = true
			}
		}
	}

	result := 0
	result2 := 0
	ended := false
	dir := 0
	dx := []int{-1, 0, 1, 0}
	dy := []int{0, 1, 0, -1}

	pairs := make(map[Pair]struct{})
	pairsLoop := make(map[Pair]struct{})

	ix, jx := startI, startJ

	for !ended {
		for {
			ni := ix + dx[dir]
			nj := jx + dy[dir]
			if ni < 0 || ni >= len(matrix) || nj < 0 || nj >= len(matrix[ni]) {
				ended = true
				break
			}
			if matrix[ni][nj] == '#' {
				break
			}
			key := Pair{X: ni, Y: nj}
			if _, ok := pairs[key]; !ok {
				result++
				pairs[key] = struct{}{}
				if _, inLoop := pairsLoop[key]; !inLoop && !(ni == startI && nj == startJ) && tryLoop(matrix, ni, nj, (dir+1)%4) {
					pairsLoop[key] = struct{}{}
					result2++
				}
			}
			ix, jx = ni, nj
		}
		dir = (dir + 1) % 4
	}

	return result, result2
}

func main() {
	matrix, err := parseFile("input6.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	result, result2 := calculatePositionsVisited(matrix)
	fmt.Println("Result:", result)
	fmt.Println("Result2:", result2)
}
