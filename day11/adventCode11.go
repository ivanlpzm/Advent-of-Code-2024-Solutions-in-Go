package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var result []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		result = append(result, strings.Fields(line)...)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func divideString(s string) (string, string) {
	mid := len(s) / 2
	return s[:mid], s[mid:]
}

func blinkIterationForPosition(data string, iter, now int, memo map[string]map[int]int) int {
	if now == iter {
		return 1
	}

	if _, exists := memo[data]; exists {
		if cachedResult, found := memo[data][now]; found {
			return cachedResult
		}
	} else {
		memo[data] = make(map[int]int)
	}

	var result int
	if data == "0" {
		result = blinkIterationForPosition("1", iter, now+1, memo)
	} else if len(data)%2 == 0 {
		part1, part2 := divideString(data)
		number, _ := strconv.Atoi(part2)
		part2Str := strconv.Itoa(number)
		result = blinkIterationForPosition(part1, iter, now+1, memo) + blinkIterationForPosition(part2Str, iter, now+1, memo)
	} else {
		number, _ := strconv.Atoi(data)
		number *= 2024
		numberStr := strconv.Itoa(number)
		result = blinkIterationForPosition(numberStr, iter, now+1, memo)
	}

	memo[data][now] = result
	return result
}

func blinkIterationCount(datas []string, iter int) int {
	memo := make(map[string]map[int]int)
	total := 0

	for _, data := range datas {
		total += blinkIterationForPosition(data, iter, 0, memo)
	}

	return total
}

func main() {
	datas, err := parseFile("input11.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	iter := 120
	result := blinkIterationCount(datas, iter)

	fmt.Println("Total Result:", result)
}
