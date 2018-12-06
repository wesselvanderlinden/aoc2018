package day6

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSafeAreaSize(t *testing.T) {
	points := []point{
		{x: 1, y: 1},
		{x: 1, y: 6},
		{x: 8, y: 3},
		{x: 3, y: 4},
		{x: 5, y: 5},
		{x: 8, y: 9},
	}

	assert.Equal(t, 16, safeAreaSize(points, 32))
}

func TestLargestArea(t *testing.T) {
	points := []point{
		{x: 1, y: 1},
		{x: 1, y: 6},
		{x: 8, y: 3},
		{x: 3, y: 4},
		{x: 5, y: 5},
		{x: 8, y: 9},
	}

	assert.Equal(t, 17, largestArea(points))
}

func TestManhattanDistance(t *testing.T) {
	tests := []struct {
		a        point
		b        point
		expected int
	}{
		{
			a:        point{x: 1, y: 1},
			b:        point{x: 1, y: 1},
			expected: 0,
		},
		{
			a:        point{x: 1, y: 1},
			b:        point{x: 5, y: 5},
			expected: 8,
		},
		{
			a:        point{x: 3, y: 4},
			b:        point{x: 8, y: 3},
			expected: 6,
		},
	}
	for i, test := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			assert.Equal(t, test.expected, manhattanDistance(test.a, test.b))
		})
	}
}

func TestParse(t *testing.T) {
	input := `
1, 1
1, 6
8, 3
3, 4
5, 5
8, 9
`
	expected := []point{
		{x: 1, y: 1},
		{x: 1, y: 6},
		{x: 8, y: 3},
		{x: 3, y: 4},
		{x: 5, y: 5},
		{x: 8, y: 9},
	}

	r, err := parse(input)

	assert.NoError(t, err)
	assert.Equal(t, expected, r)
}

func BenchmarkSafeAreaSize(b *testing.B) {
	points := []point{
		{x: 1, y: 1},
		{x: 1, y: 6},
		{x: 8, y: 3},
		{x: 3, y: 4},
		{x: 5, y: 5},
		{x: 8, y: 9},
	}

	for i := 0; i < b.N; i++ {
		safeAreaSize(points, 32)
	}
}

func BenchmarkLargestArea(b *testing.B) {
	points := []point{
		{x: 1, y: 1},
		{x: 1, y: 6},
		{x: 8, y: 3},
		{x: 3, y: 4},
		{x: 5, y: 5},
		{x: 8, y: 9},
	}

	for i := 0; i < b.N; i++ {
		largestArea(points)
	}
}

func BenchmarkManhattanDistance(b *testing.B) {
	for i := 0; i < b.N; i++ {
		manhattanDistance(point{x: 3, y: 4}, point{x: 8, y: 3})
	}
}
