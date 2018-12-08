package day8

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMetadataSumComplicated(t *testing.T) {
	n := node{
		Metadata: []int{1, 1, 2},
		Children: []node{
			{Metadata: []int{10, 11, 12}},
			{
				Metadata: []int{2},
				Children: []node{
					{Metadata: []int{99}},
				},
			},
		},
	}

	assert.Equal(t, 66, metadataAsReferenceSum(n))
}

func TestMetadataSum(t *testing.T) {
	n := node{
		Metadata: []int{1, 1, 2},
		Children: []node{
			{Metadata: []int{10, 11, 12}},
			{
				Metadata: []int{2},
				Children: []node{
					{Metadata: []int{99}},
				},
			},
		},
	}

	assert.Equal(t, 138, metadataSum(n))
}

func TestParseTree(t *testing.T) {
	expected := node{
		Metadata: []int{1, 1, 2},
		Children: []node{
			{Metadata: []int{10, 11, 12}},
			{
				Metadata: []int{2},
				Children: []node{
					{Metadata: []int{99}},
				},
			},
		},
	}

	res, err := parseTree("2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2")

	assert.NoError(t, err)
	assert.Equal(t, expected, res)
}
