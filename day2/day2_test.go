package day2

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestCountChars(t *testing.T) {
	tests := []struct {
		input  string
		twice  int
		thrice int
	}{
		{input: "abcdef", twice: 0, thrice: 0},
		{input: "bababc", twice: 1, thrice: 1},
		{input: "abbcde", twice: 1, thrice: 0},
		{input: "abcccd", twice: 0, thrice: 1},
		{input: "aabcdd", twice: 1, thrice: 0},
		{input: "abcdee", twice: 1, thrice: 0},
		{input: "ababab", twice: 0, thrice: 1},
	}
	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			twice, thrice := countChars(test.input)
			assert.Equal(t, twice, test.twice)
			assert.Equal(t, thrice, test.thrice)
		})
	}
}

func TestCalcChecksum(t *testing.T) {
	input := []string{
		"abcdef",
		"bababc",
		"abbcde",
		"abcccd",
		"aabcdd",
		"abcdee",
		"ababab",
	}

	assert.Equal(t, calcChecksum(input), 12)
}

func TestCommonChars(t *testing.T) {
	tests := []struct {
		a      string
		b      string
		output string
	}{
		{a: "abcde", b: "fghij", output: ""},
		{a: "abcde", b: "axcye", output: "ace"},
		{a: "fghij", b: "fguij", output: "fgij"},
	}
	for _, test := range tests {
		t.Run(test.a+" -> "+test.b, func(t *testing.T) {
			assert.Equal(t, commonChars(test.a, test.b), test.output)
		})
	}
}

func TestFindCorrectBoxIds(t *testing.T) {
	input := []string{
		"abcde",
		"fghij",
		"klmno",
		"pqrst",
		"fguij",
		"axcye",
		"wvxyz",
	}

	assert.Equal(t, findCorrectBoxIds(input), "fgij")
}
