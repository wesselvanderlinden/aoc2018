package day7

import (
	"bytes"
	"fmt"
	"github.com/pkg/errors"
	"regexp"
	"sort"
	"strings"
)

type Solver struct{}

func (s Solver) DayNumber() uint { return 7 }

func (s Solver) SolvePart1(input string) (string, error) {
	instr, err := parse(input)
	if err != nil {
		return "", err
	}
	return toposort(instr)
}

func (s Solver) SolvePart2(input string) (string, error) { return "not implemented yet", nil }

func toposort(instructions []instruction) (string, error) {
	graph := make(map[byte][]byte)

	// map instructions into a dependency graph
	for _, instr := range instructions {
		graph[instr.FinishedBefore] = append(graph[instr.FinishedBefore], instr.Step)
		if graph[instr.Step] == nil {
			graph[instr.Step] = []byte{}
		}
	}

	var result []byte

	for len(graph) > 0 {
		var steps []byte

		// find steps with 0 dependencies
		for step, deps := range graph {
			if len(deps) == 0 {
				steps = append(steps, step)
			}
		}

		if len(steps) == 0 {
			return "", errors.New("circular dependency found")
		}

		// sort the steps alphabetically
		sort.Slice(steps, func(i, j int) bool {
			return steps[i] < steps[j]
		})

		result = append(result, steps[0])

		// remove step from graph
		delete(graph, steps[0])

		for node := range graph {
			i := bytes.IndexByte(graph[node], steps[0])

			if i >= 0 {
				graph[node] = append(graph[node][:i], graph[node][i+1:]...)
			}
		}
	}

	return string(result), nil
}

func parse(input string) ([]instruction, error) {
	reg, err := regexp.Compile(`^Step ([A-Z]) must be finished before step ([A-Z]) can begin.$`)

	if err != nil {
		return nil, err
	}

	var instructions []instruction

	for i, l := range strings.Split(strings.TrimSpace(input), "\n") {
		match := reg.FindSubmatch([]byte(l))

		if len(match) != 3 {
			return nil, errors.New(fmt.Sprint("invalid input given at line ", i))
		}

		instructions = append(instructions, instruction{
			Step:           match[1][0],
			FinishedBefore: match[2][0],
		})
	}

	return instructions, nil
}

type instruction struct {
	Step           byte
	FinishedBefore byte
}
