package day7

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToposort(t *testing.T) {
	instructions := []instruction{
		{Step: 'C', FinishedBefore: 'A'},
		{Step: 'C', FinishedBefore: 'F'},
		{Step: 'A', FinishedBefore: 'B'},
		{Step: 'A', FinishedBefore: 'D'},
		{Step: 'B', FinishedBefore: 'E'},
		{Step: 'D', FinishedBefore: 'E'},
		{Step: 'F', FinishedBefore: 'E'},
	}

	r, err := toposort(instructions)
	assert.NoError(t, err)
	assert.Equal(t, "CABDFE", r)
}

func TestParse(t *testing.T) {
	input := `
Step C must be finished before step A can begin.
Step C must be finished before step F can begin.
Step A must be finished before step B can begin.
Step A must be finished before step D can begin.
Step B must be finished before step E can begin.
Step D must be finished before step E can begin.
Step F must be finished before step E can begin.
`
	expected := []instruction{
		{Step: 'C', FinishedBefore: 'A'},
		{Step: 'C', FinishedBefore: 'F'},
		{Step: 'A', FinishedBefore: 'B'},
		{Step: 'A', FinishedBefore: 'D'},
		{Step: 'B', FinishedBefore: 'E'},
		{Step: 'D', FinishedBefore: 'E'},
		{Step: 'F', FinishedBefore: 'E'},
	}

	instr, err := parse(input)

	assert.NoError(t, err)
	assert.Equal(t, expected, instr)
}

func BenchmarkToposort(b *testing.B) {
	instructions := []instruction{
		{Step: 'C', FinishedBefore: 'A'},
		{Step: 'C', FinishedBefore: 'F'},
		{Step: 'A', FinishedBefore: 'B'},
		{Step: 'A', FinishedBefore: 'D'},
		{Step: 'B', FinishedBefore: 'E'},
		{Step: 'D', FinishedBefore: 'E'},
		{Step: 'F', FinishedBefore: 'E'},
	}
	for i := 0; i < b.N; i++ {
		toposort(instructions)
	}
}
