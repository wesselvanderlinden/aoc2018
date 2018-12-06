package day6

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Solver struct{}

func (s Solver) DayNumber() uint { return 6 }

func (s Solver) SolvePart1(input string) (string, error) {
	points, err := parse(input)
	if err != nil {
		return "", err
	}
	return fmt.Sprint(largestArea(points)), nil
}

func (s Solver) SolvePart2(input string) (string, error) {
	points, err := parse(input)
	if err != nil {
		return "", err
	}
	return fmt.Sprint(safeAreaSize(points, 10000)), nil
}

type point struct {
	x int
	y int
}

func safeAreaSize(points []point, maxSize int) int {
	gridSize := gridSize(points)

	var size int

	for y := 0; y <= gridSize.y; y++ {
		for x := 0; x <= gridSize.x; x++ {
			current := point{x: x, y: y}
			var totalDistance int

			for _, p := range points {
				totalDistance += manhattanDistance(p, current)
			}

			if totalDistance < maxSize {
				size++
			}
		}
	}

	return size
}

func largestArea(points []point) int {
	gridSize := gridSize(points)
	grid := make(map[point]int)
	infinite := make(map[int]bool)

	for x := 0; x <= gridSize.x; x++ {
		for y := 0; y <= gridSize.y; y++ {
			var id int
			var multipleIds bool

			current := point{x: x, y: y}
			smallestDistance := math.MaxInt64

			for i, p := range points {
				dist := manhattanDistance(p, current)

				if dist < smallestDistance {
					smallestDistance = dist
					id = i
					multipleIds = false
				} else if dist == smallestDistance {
					multipleIds = true
				}
			}

			if multipleIds {
				continue
			}

			grid[current] = id

			if !infinite[id] && (x == 0 || x == gridSize.x || y == 0 || y == gridSize.y) {
				infinite[id] = true
			}
		}
	}

	counts := make(map[int]int)
	var maxCount int

	for _, id := range grid {
		if infinite[id] {
			continue
		}

		counts[id]++

		if counts[id] > maxCount {
			maxCount = counts[id]
		}
	}

	return maxCount
}

func gridSize(points []point) point {
	var size point
	for _, p := range points {
		if p.x > size.x {
			size.x = p.x
		}
		if p.y > size.y {
			size.y = p.y
		}
	}
	return size
}

func manhattanDistance(a, b point) int {
	return int(math.Abs(float64(a.x-b.x)) + math.Abs(float64(a.y-b.y)))
}

func parse(input string) ([]point, error) {
	var points []point
	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		parts := strings.Split(line, ", ")
		if len(parts) != 2 {
			return nil, errors.New("invalid input")
		}
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		points = append(points, point{x: x, y: y})
	}
	return points, nil
}
