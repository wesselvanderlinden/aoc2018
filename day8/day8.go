package day8

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Solver struct{}

func (s Solver) DayNumber() uint { return 8 }

func (s Solver) SolvePart1(input string) (string, error) {
	n, err := parseTree(input)
	if err != nil {
		return "", err
	}
	return fmt.Sprint(metadataSum(n)), nil
}

func (s Solver) SolvePart2(input string) (string, error) {
	n, err := parseTree(input)
	if err != nil {
		return "", err
	}
	return fmt.Sprint(metadataAsReferenceSum(n)), nil
}

func metadataAsReferenceSum(n node) (sum int) {
	if len(n.Children) == 0 {
		for _, m := range n.Metadata {
			sum += m
		}
		return sum
	}

	for _, m := range n.Metadata {
		idx := m - 1

		if idx < 0 || idx >= len(n.Children) {
			continue
		}

		sum += metadataAsReferenceSum(n.Children[idx])
	}

	return sum
}

func metadataSum(n node) (sum int) {
	for _, m := range n.Metadata {
		sum += m
	}
	for _, c := range n.Children {
		sum += metadataSum(c)
	}
	return sum
}

type node struct {
	Metadata []int
	Children []node
}

type treeParser struct {
	parts []string
	pos   int
	err   error
}

func (t *treeParser) parse(input string) (node, error) {
	t.parts = strings.Split(strings.TrimSpace(input), " ")
	n := t.parseNode()
	if t.err != nil {
		return n, t.err
	}
	if t.pos < len(t.parts) {
		return n, errors.New("more than 1 root node discovered")
	}
	return n, nil
}

func (t *treeParser) parseNode() (n node) {
	if t.err != nil {
		return n
	}

	numChildren := t.next()
	numMeta := t.next()

	for i := 0; i < numChildren; i++ {
		n.Children = append(n.Children, t.parseNode())
	}

	for i := 0; i < numMeta; i++ {
		n.Metadata = append(n.Metadata, t.next())
	}

	return n
}

func (t *treeParser) next() (v int) {
	if t.err != nil {
		return v
	}
	if t.pos >= len(t.parts) {
		t.err = errors.New(fmt.Sprint("invalid position: ", t.pos))
		return v
	}
	v, t.err = strconv.Atoi(t.parts[t.pos])
	t.pos++
	return v
}

func parseTree(input string) (node, error) {
	t := &treeParser{}
	return t.parse(input)
}
