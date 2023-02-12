package binarylifting_test

import (
	"testing"

	"github.com/ilyadt/binarylifting"

	"github.com/stretchr/testify/assert"
)

func TestTree_LCA(t *testing.T) {
	data := []*struct {
		root    rune
		adjList [][]rune
		n       int
		lca     [][3]rune
	}{
		{
			root: 'a',
			adjList: [][]rune{
				'a': {'b', 'c', 'd'},
				'b': {'e', 'f'},
				'c': {'g', 'h'},
				'd': {'i', 'j'},
				'e': {'k', 'l'},
				'f': {'m'},
				'k': {'n', 'o'},
				'm': {'p', 'q'},
				'p': {'r'},
				'r': {'s'},

				// Length of 128
				127: nil,
			},
			n: 128,
			lca: [][3]rune{
				{'a', 'a', 'a'},
				{'a', 'b', 'a'},
				{'b', 'b', 'b'},
				{'n', 'o', 'k'},
				{'s', 'n', 'b'},
				{'r', 'i', 'a'},
			},
		},
	}

	for _, el := range data {
		tr := binarylifting.NewTree[rune](el.root, el.adjList, el.n)

		for _, tc := range el.lca {
			assert.Equal(t, tc[2], tr.LCA(tc[0], tc[1]))
		}
	}
}
