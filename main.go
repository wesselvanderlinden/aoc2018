package main

import (
	"aoc2018/day1"
	"aoc2018/day2"
	"aoc2018/day3"
	"aoc2018/day4"
	"aoc2018/day5"
	"fmt"
	"io/ioutil"
	"time"
)

func main() {
	solve(day1.Solver{})
	solve(day2.Solver{})
	solve(day3.Solver{})
	solve(day4.Solver{})
	solve(day5.Solver{})
}

func solve(solver solver) {
	inputFile := fmt.Sprint("input/day", solver.DayNumber(), ".txt")
	input, err := ioutil.ReadFile(inputFile)

	if err != nil {
		panic(err)
	}

	fmt.Println("------ Day", solver.DayNumber(), "------")

	fmt.Println("Part 1")
	solvePart(solver.SolvePart1, string(input))
	fmt.Println()
	fmt.Println("Part 2")
	solvePart(solver.SolvePart2, string(input))
	fmt.Println()
}

func solvePart(f solverFunc, input string) {
	s := time.Now()
	answer, err := f(input)

	fmt.Println("\tDuration:", time.Since(s))

	if err != nil {
		fmt.Println("\tError:", err)
	} else {
		fmt.Println("\tAnswer:", answer)
	}
}

type solver interface {
	DayNumber() uint
	SolvePart1(input string) (string, error)
	SolvePart2(input string) (string, error)
}

type solverFunc func(input string) (string, error)
