package day1

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestSum(t *testing.T) {
	tests := []struct {
		input  []int
		output int
	}{
		{input: []int{1, 1, 1}, output: 3},
		{input: []int{1, 1, -2}, output: 0},
		{input: []int{-1, -2, -3}, output: -6},
	}

	for _, test := range tests {
		assert.Equal(t, sum(test.input), test.output)
	}
}

func TestFindFirstDoubleFrequency(t *testing.T) {
	tests := []struct {
		input  []int
		output int
	}{
		{input: []int{1, -1}, output: 0},
		{input: []int{3, 3, 4, -2, -4}, output: 10},
		{input: []int{-6, 3, 8, 5, -6}, output: 5},
		{input: []int{7, 7, -2, -7, -4}, output: 14},
	}

	for _, test := range tests {
		assert.Equal(t, findFirstDoubleFrequency(test.input), test.output)
	}
}

func TestParse(t *testing.T) {
	tests := []struct {
		input string
		freqs []int
		err   error
	}{
		{input: "+1, +1, +1", freqs: []int{1, 1, 1}},
		{input: "+1, +1, -2", freqs: []int{1, 1, -2}},
		{input: "-1\n-2\n-3", freqs: []int{-1, -2, -3}},
	}

	for _, test := range tests {
		freqs, err := parse(test.input)
		if err != nil {
			t.Error(err)
		} else {
			assert.Equal(t, freqs, test.freqs)
		}
	}
}
