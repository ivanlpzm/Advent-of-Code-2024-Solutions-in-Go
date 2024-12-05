package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseFile(filename string) (map[int][]int, [][]int, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, nil, err
	}

	lines := strings.Split(string(content), "\n")

	section1 := make(map[int][]int)
	var section2 [][]int
	var isSection2 bool

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			isSection2 = true
			continue
		}

		if !isSection2 {
			// Parse the first section into the map
			pair := strings.Split(line, "|")
			if len(pair) != 2 {
				return nil, nil, fmt.Errorf("invalid pair format: %s", line)
			}
			key, err := strconv.Atoi(strings.TrimSpace(pair[0]))
			if err != nil {
				return nil, nil, fmt.Errorf("invalid number format in key: %v", err)
			}
			value, err := strconv.Atoi(strings.TrimSpace(pair[1]))
			if err != nil {
				return nil, nil, fmt.Errorf("invalid number format in value: %v", err)
			}
			section1[key] = append(section1[key], value)
		} else {
			// Parse the second section into a 2D slice
			values := strings.Split(line, ",")
			var parsedLine []int
			for _, value := range values {
				num, err := strconv.Atoi(strings.TrimSpace(value))
				if err != nil {
					return nil, nil, fmt.Errorf("invalid number format: %v", err)
				}
				parsedLine = append(parsedLine, num)
			}
			section2 = append(section2, parsedLine)
		}
	}

	return section1, section2, nil
}

func calculateCorrectRowsOrder(dict map[int][]int, matrix [][]int) []int {
	var result []int

	for i, row := range matrix {
		storedValues := make(map[int]struct{})
		isOk := true

		for _, val := range row {
			for storedVal := range storedValues {
				if contains(dict[val], storedVal) {
					isOk = false
					break
				}
			}
			if !isOk {
				break
			}
			storedValues[val] = struct{}{}
		}

		if isOk {
			result = append(result, i)
		}
	}

	return result
}

func contains(slice []int, value int) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func sumMedValueCorrectRows(rows []int, matrix [][]int) int {
	result := 0
	for _, row := range rows {
		length := len(matrix[row])
		mid := length / 2
		result += matrix[row][mid]
	}
	return result
}

func calculateCorrectRowsOrderWithFixing(dict map[int][]int, matrix [][]int) []int {
	var result []int

	for i, row := range matrix {
		storedValues := make(map[int]struct{})
		isOk := true

		for _, val := range row {
			for storedVal := range storedValues {
				if contains(dict[val], storedVal) {
					isOk = false
					break
				}
			}
			if !isOk {
				break
			}
			storedValues[val] = struct{}{}
		}

		// Fix the row if it is not in the correct order
		if !isOk {
			for changed := true; changed; {
				changed = false
				for j := 0; j < len(row)-1; j++ {
					// Swap if the order is incorrect based on dict
					if contains(dict[row[j]], row[j+1]) {
						row[j], row[j+1] = row[j+1], row[j]
						changed = true
					}
				}
			}
			result = append(result, i)
		}
	}

	return result
}

func main() {
	section1, section2, err := parseFile("input5.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	rows := calculateCorrectRowsOrder(section1, section2)

	result := sumMedValueCorrectRows(rows, section2)

	fmt.Println("Result Part one", result)

	rows2 := calculateCorrectRowsOrderWithFixing(section1, section2)

	result2 := sumMedValueCorrectRows(rows2, section2)

	fmt.Println("Result Part two", result2)
}
