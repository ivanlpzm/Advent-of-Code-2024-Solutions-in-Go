package main

import (
	"fmt"
	"os"
	"strings"
	"math"
	"sort"
)

func parseFile(filename string) ([]int, []int, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, nil, err
	}

	var a, b []int
	lines := strings.Split(string(content), "\n")

	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}
		parts := strings.Fields(line)
		if len(parts) == 2 {
			var left, right int
			fmt.Sscanf(parts[0], "%d", &left)
			fmt.Sscanf(parts[1], "%d", &right)
			a = append(a, left)
			b = append(b, right)
		}
	}
	return a, b, nil
}

func calculateSumOfDifferences(a, b []int) int {
	sum := 0
	for i := 0; i < len(a); i++ {
		sum += int(math.Abs(float64(a[i] - b[i])))
	}
	return sum
}

func calculateSimilarityScore(a, b []int) int {
	frequency := make(map[int]int)
	for _, num := range b {
		frequency[num]++
	}

	score := 0
	for _, num := range a {
		score += num * frequency[num]
	}
	return score
}


func main() {
	a, b, err := parseFile("input1.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	sort.Ints(a)
	sort.Ints(b)

	sum := calculateSumOfDifferences(a, b)
	similarities := calculateSimilarityScore(a, b)

	fmt.Println("Sum of differences:", sum)
	fmt.Println("Similarities:", similarities)
}
