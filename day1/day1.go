package day1

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Solver struct{}

func (s Solver) DayNumber() uint {
	return 1
}

func (s Solver) SolvePart1(input string) (string, error) {
	freqs, err := parse(input)

	if err != nil {
		return "", err
	}

	return fmt.Sprint(sum(freqs)), nil
}

func (s Solver) SolvePart2(input string) (string, error) {
	freqs, err := parse(input)

	if err != nil {
		return "", err
	}

	return fmt.Sprint(findFirstDoubleFrequency(freqs)), nil
}

func sum(v []int) (p int) {
	for _, vv := range v {
		p += vv
	}
	return p
}

func findFirstDoubleFrequency(freqs []int) int {
	seen := map[int]bool{0: true}
	freq := 0

	for {
		for _, f := range freqs {
			freq += f

			if seen[freq] == true {
				return freq
			}

			seen[freq] = true
		}
	}
}

func parse(input string) (freqs []int, err error) {
	var sep string

	if strings.Contains(input, ", ") {
		sep = ", "
	} else if strings.Contains(input, "\n") {
		sep = "\n"
	} else {
		return nil, errors.New("no valid separator found")
	}

	for _, c := range strings.Split(input, sep) {
		var v int
		v, err = strconv.Atoi(c)

		if err != nil {
			return nil, err
		}

		freqs = append(freqs, v)
	}

	return freqs, nil
}
