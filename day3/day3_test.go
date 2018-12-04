package day3

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestOverlappingSquares(t *testing.T) {
	claims := []claim{
		{Id: 1, Left: 1, Top: 3, Width: 4, Height: 4},
		{Id: 2, Left: 3, Top: 1, Width: 4, Height: 4},
		{Id: 3, Left: 5, Top: 5, Width: 2, Height: 2},
	}

	assert.Equal(t, overlappingSquares(claims), 4)
}

func TestNonOverlappingClaim(t *testing.T) {
	claims := []claim{
		{Id: 1, Left: 1, Top: 3, Width: 4, Height: 4},
		{Id: 2, Left: 3, Top: 1, Width: 4, Height: 4},
		{Id: 3, Left: 5, Top: 5, Width: 2, Height: 2},
	}

	result := nonOverlappingClaim(claims)
	assert.NotNil(t, result)
	assert.Equal(t, &claim{Id: 3, Left: 5, Top: 5, Width: 2, Height: 2}, result)
}

func TestParse(t *testing.T) {
	expected := []claim{
		{Id: 123, Left: 3, Top: 2, Width: 5, Height: 4},
		{Id: 1181, Left: 500, Top: 174, Width: 12, Height: 16},
	}

	input := "#123 @ 3,2: 5x4\n#1181 @ 500,174: 12x16"

	assert.Equal(t, parse(input), expected)
}

func TestParseClaim(t *testing.T) {
	tests := []struct {
		claim    string
		expected claim
	}{
		{claim: "#123 @ 3,2: 5x4", expected: claim{Id: 123, Left: 3, Top: 2, Width: 5, Height: 4}},
		{claim: "#1181 @ 500,174: 12x16", expected: claim{Id: 1181, Left: 500, Top: 174, Width: 12, Height: 16}},
	}
	for _, test := range tests {
		t.Run(test.claim, func(t *testing.T) {
			assert.Equal(t, parseClaim(test.claim), test.expected)
		})
	}
}

func BenchmarkOverlappingSquares(b *testing.B) {
	var claims []claim
	for i := 0; i < 500; i++ {
		claims = append(claims, claim{
			Id:     i,
			Left:   rand.Intn(500),
			Top:    rand.Intn(500),
			Width:  rand.Intn(500),
			Height: rand.Intn(500),
		})
	}

	for i := 0; i < b.N; i++ {
		overlappingSquares(claims)
	}
}
