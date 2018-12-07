package main

import (
	"aoc2018/day1"
	"aoc2018/day2"
	"aoc2018/day3"
	"aoc2018/day4"
	"aoc2018/day5"
	"aoc2018/day6"
	"aoc2018/day7"
	"fmt"
	"io/ioutil"
	"os"
	"text/tabwriter"
	"time"
)

func main() {
	solveAll([]solver{
		day1.Solver{},
		day2.Solver{},
		day3.Solver{},
		day4.Solver{},
		day5.Solver{},
		day6.Solver{},
		day7.Solver{},
	})
}

func solveAll(solvers []solver) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)

	fmt.Fprintln(w, "Day\tPart\tDuration\tAnswer")
	fmt.Fprintln(w, "---\t----\t--------\t------")

	for _, s := range solvers {
		solve(s, w)
	}

	w.Flush()
}

func solve(solver solver, w *tabwriter.Writer) {
	inputFile := fmt.Sprint("input/day", solver.DayNumber(), ".txt")
	input, err := ioutil.ReadFile(inputFile)

	if err != nil {
		panic(err)
	}

	fmt.Fprint(w, solver.DayNumber())
	fmt.Fprint(w, "\t", 1)
	solvePart(solver.SolvePart1, string(input), w)
	fmt.Fprint(w, "\t", 2)
	solvePart(solver.SolvePart2, string(input), w)
}

func solvePart(f solverFunc, input string, w *tabwriter.Writer) {
	s := time.Now()
	answer, err := f(input)

	fmt.Fprint(w, "\t", time.Since(s))

	if err != nil {
		fmt.Fprint(w, "\t\033[31m", err, "\033[0m")
	} else {
		fmt.Fprint(w, "\t", answer)
	}

	fmt.Fprintln(w)
}

type solver interface {
	DayNumber() uint
	SolvePart1(input string) (string, error)
	SolvePart2(input string) (string, error)
}

type solverFunc func(input string) (string, error)
