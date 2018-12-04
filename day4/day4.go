package day4

import (
	"fmt"
	"math"
	"time"
)

type Solver struct{}

func (s Solver) DayNumber() uint {
	return 4
}

func (s Solver) SolvePart1(input string) (string, error) {
	return solve(input, strategy1)
}

func (s Solver) SolvePart2(input string) (string, error) {
	return solve(input, strategy2)
}

func solve(input string, s strategy) (string, error) {
	records, err := Parse(input)
	if err != nil {
		return "", err
	}
	return fmt.Sprint(s(records)), nil
}

type strategy func(records []*record) int

func strategy1(records []*record) int {
	guardMap := buildMap(records)

	var id, minute, totalSleep int

	for g, ma := range guardMap {
		total := ma.TotalTime()

		if total > totalSleep {
			totalSleep = total
			id = g
			_, minute = ma.MostUsedMinute()
		}
	}

	return id * minute
}

func strategy2(records []*record) int {
	guardMap := buildMap(records)

	var id, minute, timesUsed int

	for i, ma := range guardMap {
		t, m := ma.MostUsedMinute()

		if t > timesUsed {
			timesUsed = t
			minute = m
			id = i
		}
	}

	return id * minute
}

func buildMap(records []*record) map[int]minutesAsleep {
	guardMap := make(map[int]minutesAsleep)

	var guardId int
	var sleepStart time.Time

	for _, r := range records {
		switch r.Type {
		case typeBegin:
			guardId = r.Id
			break

		case typeFallAsleep:
			sleepStart = r.Date
			break

		case typeWakeUp:
			if guardMap[guardId] == nil {
				guardMap[guardId] = make(minutesAsleep)
			}

			guardMap[guardId].AddSleep(sleepStart, r.Date)
			break
		}
	}

	return guardMap
}

type minutesAsleep map[int]int

func (m minutesAsleep) AddSleep(start time.Time, end time.Time) {
	for i := start; i.Sub(end) <= 0; i = i.Add(time.Minute) {
		m[i.Minute()]++
	}
}

func (m minutesAsleep) TotalTime() (total int) {
	for _, t := range m {
		total += t
	}
	return total
}

func (ma minutesAsleep) MostUsedMinute() (timesUsed int, minute int) {
	timesUsed = 0
	minute = int(math.Inf(1))

	for m, t := range ma {
		if t > timesUsed || (t == timesUsed && m < minute) {
			timesUsed = t
			minute = m
		}
	}

	return timesUsed, minute
}
