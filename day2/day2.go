package day2

import (
	"fmt"
	"strings"
)

type Solver struct{}

func (s Solver) DayNumber() uint {
	return 2
}

func (s Solver) SolvePart1(input string) (string, error) {
	boxIds := strings.Split(input, "\n")

	return fmt.Sprint(calcChecksum(boxIds)), nil
}

func (s Solver) SolvePart2(input string) (string, error) {
	boxIds := strings.Split(input, "\n")

	return findCorrectBoxIds(boxIds), nil
}

func findCorrectBoxIds(boxIds []string) string {
	for i, idA := range boxIds {
		for _, idB := range boxIds[i+1:] {
			common := commonChars(idA, idB)

			if len(common) == len(idA)-1 {
				return common
			}
		}
	}

	return ""
}

func commonChars(a string, b string) string {
	var result []byte
	for i := 0; i < len(a); i++ {
		if a[i] == b[i] {
			result = append(result, a[i])
		}
	}
	return string(result)
}

func calcChecksum(boxIds []string) int {
	twice := 0
	thrice := 0

	for _, i := range boxIds {
		tw, th := countChars(i)
		twice += tw
		thrice += th
	}

	return twice * thrice
}

func countChars(boxId string) (twice, thrice int) {
	counts := make(map[int32]int)

	for _, c := range boxId {
		counts[c]++
	}

	for _, c := range counts {
		if c == 2 {
			twice = 1
		} else if c == 3 {
			thrice = 1
		}
	}

	return twice, thrice
}
