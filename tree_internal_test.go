package binarylifting

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTree(t *testing.T) {
	data := []*struct {
		root    rune
		adjList [][]rune
		n       int
		log2n   int
		levels  []int
		parents [][]rune
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
			n:     128,
			log2n: 7,
			parents: [][]rune{
				'b': {'a', -1, -1, -1, -1, -1, -1},
				'c': {'a', -1, -1, -1, -1, -1, -1},
				'd': {'a', -1, -1, -1, -1, -1, -1},
				'e': {'b', 'a', -1, -1, -1, -1, -1},
				'f': {'b', 'a', -1, -1, -1, -1, -1},
				'g': {'c', 'a', -1, -1, -1, -1, -1},
				'h': {'c', 'a', -1, -1, -1, -1, -1},
				'i': {'d', 'a', -1, -1, -1, -1, -1},
				'j': {'d', 'a', -1, -1, -1, -1, -1},
				'k': {'e', 'b', -1, -1, -1, -1, -1},
				'l': {'e', 'b', -1, -1, -1, -1, -1},
				'm': {'f', 'b', -1, -1, -1, -1, -1},
				'n': {'k', 'e', 'a', -1, -1, -1, -1},
				'o': {'k', 'e', 'a', -1, -1, -1, -1},
				'p': {'m', 'f', 'a', -1, -1, -1, -1},
				'q': {'m', 'f', 'a', -1, -1, -1, -1},
				'r': {'p', 'm', 'b', -1, -1, -1, -1},
				's': {'r', 'p', 'f', -1, -1, -1, -1},

				// Length of 128
				127: nil,
			},
			levels: []int{
				'a': 0,
				'b': 1, 'c': 1, 'd': 1,
				'e': 2, 'f': 2, 'g': 2, 'h': 2, 'i': 2, 'j': 2,
				'k': 3, 'l': 3, 'm': 3,
				'n': 4, 'o': 4, 'p': 4, 'q': 4,
				'r': 5,
				's': 6,

				// Length of 128
				127: 0,
			},
		},
	}

	for _, el := range data {
		el := el

		emptyParents := make([]rune, el.log2n)
		for i := 0; i < len(emptyParents); i++ {
			emptyParents[i] = -1
		}

		for i := 0; i < el.n; i++ {
			if el.parents[i] == nil {
				el.parents[i] = emptyParents
			}
		}

		tr := NewTree[rune](el.root, el.adjList, el.n)

		assert.Equal(t, el.n, tr.n)
		assert.Equal(t, el.log2n, tr.log2n)
		assert.Equal(t, el.levels, tr.level)
		assert.Equal(t, el.parents, tr.parents)
		assert.Equal(t, 0, tr.level[el.root])
	}
}
