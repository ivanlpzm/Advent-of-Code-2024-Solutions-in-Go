package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Data struct {
	Target int
	Values []int
}

func parseFile(filename string) ([]Data, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var result []Data
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			return nil, fmt.Errorf("malformed line: %s", line)
		}

		target, err := strconv.Atoi(strings.TrimSpace(parts[0]))
		if err != nil {
			return nil, err
		}

		values := strings.Fields(parts[1])
		numValues := make([]int, len(values))
		for i, val := range values {
			numValues[i], err = strconv.Atoi(val)
			if err != nil {
				return nil, err
			}
		}

		result = append(result, Data{Target: target, Values: numValues})
	}

	return result, scanner.Err()
}

func generateOperations(n int, ops []string, includeConcat bool, results *[][]string) {
	if n == 0 {
		*results = append(*results, append([]string{}, ops...))
		return
	}
	generateOperations(n-1, append(ops, "+"), includeConcat, results)
	generateOperations(n-1, append(ops, "*"), includeConcat, results)
	if includeConcat {
		generateOperations(n-1, append(ops, "||"), includeConcat, results)
	}
}

func evaluateWithPruning(A []int, ops []string, T int, includeConcat bool) bool {
	result := A[0]
	for i := 1; i < len(A); i++ {
		switch ops[i-1] {
		case "+":
			result += A[i]
		case "*":
			result *= A[i]
		case "||":
			if includeConcat {
				result = concatenateInts(result, A[i])
			}
		}
		if result > T {
			return false
		}
	}
	return result == T
}

func concatenateInts(a, b int) int {
	result, _ := strconv.Atoi(fmt.Sprintf("%d%d", a, b))
	return result
}

func canDecompose(T int, A []int, includeConcat bool) bool {
	if len(A) == 1 {
		return A[0] == T
	}

	var operations [][]string
	generateOperations(len(A)-1, []string{}, includeConcat, &operations)

	for _, ops := range operations {
		if evaluateWithPruning(A, ops, T, includeConcat) {
			return true
		}
	}
	return false
}

func calculateSumDecomposed(datas []Data, includeConcat bool) int {
	result := 0
	for _, entry := range datas {
		if canDecompose(entry.Target, entry.Values, includeConcat) {
			result += entry.Target
		}
	}
	return result
}

func main() {
	data, err := parseFile("input7.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	result := calculateSumDecomposed(data, false)
	result2 := calculateSumDecomposed(data, true)

	fmt.Println("Result:", result)
	fmt.Println("Result2:", result2)
}
