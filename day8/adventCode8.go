package main

import (
	"fmt"
	"os"
	"strings"
)

type Point struct {
	X, Y int
}

func inbounds(data [][]byte, x, y int) bool {
	return x >= 0 && x < len(data[0]) && y >= 0 && y < len(data)
}

// Calculates antinodes for a pair of points
func antinodes(data [][]byte, p1, p2 Point) (map[Point]struct{}, map[Point]struct{}) {
	p1Pts := make(map[Point]struct{})
	p2Pts := map[Point]struct{}{
		p1: {},
		p2: {},
	}

	dx := p2.X - p1.X
	dy := p2.Y - p1.Y

	// Extend backward from p1
	if inbounds(data, p1.X-dx, p1.Y-dy) {
		p1Pts[Point{p1.X - dx, p1.Y - dy}] = struct{}{}
	}
	// Extend forward from p2
	if inbounds(data, p2.X+dx, p2.Y+dy) {
		p1Pts[Point{p2.X + dx, p2.Y + dy}] = struct{}{}
	}

	curX, curY := p1.X, p1.Y
	for {
		curX -= dx
		curY -= dy
		if !inbounds(data, curX, curY) {
			break
		}
		p2Pts[Point{curX, curY}] = struct{}{}
	}

	curX, curY = p1.X, p1.Y
	for {
		curX += dx
		curY += dy
		if !inbounds(data, curX, curY) {
			break
		}
		p2Pts[Point{curX, curY}] = struct{}{}
	}

	return p1Pts, p2Pts
}

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

func main() {
	data, _ := parseFile("input8.txt")

	lut := make(map[byte][]Point)
	for y := 0; y < len(data); y++ {
		for x := 0; x < len(data[0]); x++ {
			if data[y][x] == '.' {
				continue
			}
			lut[data[y][x]] = append(lut[data[y][x]], Point{x, y})
		}
	}

	// Calculate antinodes
	p1 := make(map[Point]struct{})
	p2 := make(map[Point]struct{})
	for _, points := range lut {
		for i := 0; i < len(points); i++ {
			for j := i + 1; j < len(points); j++ {
				p1Pts, p2Pts := antinodes(data, points[i], points[j])
				for pt := range p1Pts {
					p1[pt] = struct{}{}
				}
				for pt := range p2Pts {
					p2[pt] = struct{}{}
				}
			}
		}
	}

	fmt.Printf("Part 1: %d\n", len(p1))
	fmt.Printf("Part 2: %d\n", len(p2))
}
