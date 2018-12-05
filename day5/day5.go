package day5

import (
	"fmt"
	"math"
	"strings"
)

type Solver struct{}

func (s Solver) DayNumber() uint {
	return 5
}

func (s Solver) SolvePart1(input string) (string, error) {
	result := len(reactPolymer(input))
	return fmt.Sprint(result), nil
}

func (s Solver) SolvePart2(input string) (string, error) {
	length := math.MaxInt64

	for i := 65; i <= 90; i++ {
		s := strings.Replace(input, string(i), "", -1)
		s = strings.Replace(s, string(i+32), "", -1)
		l := len(reactPolymer(s))

		if l < length {
			length = l
		}
	}

	return fmt.Sprint(length), nil
}

func reactPolymer(input string) string {
	var reaction []rune

	for _, c := range input {
		if len(reaction) > 0 && isReactive(reaction[len(reaction)-1], c) {
			reaction = reaction[:len(reaction)-1]
		} else {
			reaction = append(reaction, c)
		}
	}

	return string(reaction)
}

func isReactive(a, b rune) bool {
	return math.Abs(float64(a-b)) == 32
}
