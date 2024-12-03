package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func parseFile(filename string) (string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func extractMulValues(input string) []string {
	// Define the regex pattern for "mul(x,y)"
	pattern := `mul\(\s*\d+\s*,\s*\d+\s*\)`

	re := regexp.MustCompile(pattern)

	matches := re.FindAllString(input, -1)

	return matches
}

func multiply(input []string) int {
	sum := 0
	for _, s := range input {
		s = strings.TrimPrefix(s, "mul(")
		s = strings.TrimSuffix(s, ")")
		values := strings.Split(s, ",")

		if len(values) == 2 {
			value0, err0 := strconv.Atoi(strings.TrimSpace(values[0]))
			value1, err1 := strconv.Atoi(strings.TrimSpace(values[1]))
			if err0 == nil && err1 == nil {
				sum += value0 * value1
			}
		}
	}
	return sum
}

func multiplyWithInstructions(input string) int {
	pattern := `(do\(\)|don't\(\)|mul\(\s*\d+\s*,\s*\d+\s*\))`
	re := regexp.MustCompile(pattern)
	instructions := re.FindAllString(input, -1)

	enabled := true
	sum := 0

	for _, instr := range instructions {
		if instr == "do()" {
			enabled = true
		} else if instr == "don't()" {
			enabled = false
		} else if strings.HasPrefix(instr, "mul(") && enabled {
			instr = strings.TrimPrefix(instr, "mul(")
			instr = strings.TrimSuffix(instr, ")")
			values := strings.Split(instr, ",")
			if len(values) == 2 {
				value0, err0 := strconv.Atoi(strings.TrimSpace(values[0]))
				value1, err1 := strconv.Atoi(strings.TrimSpace(values[1]))
				if err0 == nil && err1 == nil {
					sum += value0 * value1
				}
			}
		}
	}

	return sum
}

func main() {
	input, err := parseFile("input3.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	formatValue := extractMulValues(input)
	result1 := multiply(formatValue)
	result2 := multiplyWithInstructions(input)
	fmt.Println(result1)
	fmt.Println(result2)
}
