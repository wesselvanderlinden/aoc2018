package day4

import (
	"errors"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

func Parse(input string) ([]*record, error) {
	p := &parser{}
	p.Parse(input)
	return p.records, p.err
}

type parser struct {
	records []*record
	err     error
}

func (p *parser) Parse(input string) {
	p.parseRecords(strings.Trim(input, " \n"))
	p.sortRecords()
}

func (p *parser) parseRecords(input string) {
	if p.err != nil {
		return
	}
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		p.parseRecord(line)
	}
}

func (p *parser) parseRecord(line string) {
	if p.err != nil {
		return
	}

	var r record

	reg, err := regexp.Compile(`^\[(.*)\] (Guard #(\d+) begins shift|falls asleep|wakes up)`)
	if err != nil {
		p.err = err
		return
	}

	matches := reg.FindStringSubmatch(line)
	if len(matches) < 3 && len(matches) > 4 {
		p.err = err
		return
	}

	r.Date, err = time.Parse("2006-01-02 15:04", matches[1])
	if err != nil {
		p.err = err
		return
	}

	switch matches[2] {
	case "falls asleep":
		r.Type = typeFallAsleep
		break

	case "wakes up":
		r.Type = typeWakeUp
		break

	default:
		if len(matches) != 4 {
			p.err = errors.New("no guard id found")
			return
		}

		r.Type = typeBegin
		r.Id, err = strconv.Atoi(matches[3])

		if err != nil {
			p.err = err
			return
		}
		break
	}

	p.records = append(p.records, &r)
}

func (p *parser) sortRecords() {
	if p.err != nil {
		return
	}
	sort.Slice(p.records, func(i, j int) bool {
		return p.records[i].Date.Before(p.records[j].Date)
	})
}

const (
	typeBegin = iota
	typeFallAsleep
	typeWakeUp
)

type record struct {
	Id   int
	Date time.Time
	Type int
}
