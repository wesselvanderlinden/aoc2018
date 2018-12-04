package day3

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Solver struct{}

func (s Solver) DayNumber() uint {
	return 3
}

func (s Solver) SolvePart1(input string) (string, error) {
	claims := parse(input)
	return fmt.Sprint(overlappingSquares(claims)), nil
}

func (s Solver) SolvePart2(input string) (string, error) {
	claims := parse(input)
	nonOverlapping := nonOverlappingClaim(claims)
	if nonOverlapping == nil {
		return "", errors.New("no non overlapping claim found")
	}
	return fmt.Sprint(nonOverlapping.Id), nil
}

type point struct {
	x int
	y int
}

type claim struct {
	Id     int
	Left   int
	Top    int
	Width  int
	Height int
}

func overlappingSquares(claims []claim) (overlapping int) {
	var grid [1000][1000]int

	for _, cl := range claims {
		for x := 0; x < cl.Width; x++ {
			for y := 0; y < cl.Height; y++ {
				rx := cl.Left + x
				by := cl.Top + y
				if grid[rx][by] == 1 {
					overlapping++
				}
				grid[rx][by]++
			}
		}
	}

	return overlapping
}

func nonOverlappingClaim(claims []claim) *claim {
outer:
	for _, c := range claims {
		for _, o := range claims {
			if c.Id == o.Id {
				continue
			}

			if overlaps(c, o) {
				continue outer
			}
		}

		return &c
	}

	return nil
}

func overlaps(a, b claim) bool {
	return a.Left < b.Left+b.Width &&
		a.Left+a.Width > b.Left &&
		a.Top < b.Top+b.Height &&
		a.Top+a.Height > b.Top
}

func parse(input string) (claims []claim) {
	for _, s := range strings.Split(input, "\n") {
		claims = append(claims, parseClaim(s))
	}
	return claims
}

func parseClaim(c string) claim {
	r := regexp.MustCompile(`#(\d+) @ (\d+),(\d+): (\d+)x(\d+)`)
	matches := r.FindStringSubmatch(c)

	if len(matches) != 6 {
		panic("invalid")
	}

	toInt := func(v string) int {
		i, _ := strconv.Atoi(v)
		return i
	}

	return claim{
		Id:     toInt(matches[1]),
		Left:   toInt(matches[2]),
		Top:    toInt(matches[3]),
		Width:  toInt(matches[4]),
		Height: toInt(matches[5]),
	}
}
