// Link https://contest.cs.cmu.edu/295/s20/tutorials/lca.mark
package binarylifting

import "math"

type Node interface {
	int | int8 | int16 | int32 | int64
}

type Tree[T Node] struct {
	n     int // Number of nodes.
	log2n int // log2(n)

	// Level[i] is the height of node i in the tree. Root node has level 0.
	level []int

	// 2^j parent of i.
	parents [][]T
}

func (r *Tree[T]) LCA(u, v T) T {
	// Suppose "u" is always deeper than "v" without loss of generality
	if r.level[v] > r.level[u] {
		u, v = v, u
	}

	// Step 1:
	// Move the lower of the two vertices up the tree so that it is on the same level as the other one.
	// We can do this in O(log(n)) steps using the binary ancestors by noticing the following useful fact.
	// We want to move it level[u]−level[v] steps up. We can do this since every number can be written
	// as a sum of powers of two (this is possible because we can write the number in binary).
	// Therefore, we can simply greedily move u up by the largest power of two that does not overshoot
	// until we reach the desired level. This takes at most log(n) time.
	vLevel := r.level[v]
	for i := r.log2n - 1; i >= 0; i-- {
		if r.level[u]-pow(2, i) >= vLevel {
			u = r.parents[u][i]
		}
	}

	if u == v {
		return u
	}

	// Step 2:
	// Move both u and v up the tree as far as possible without them touching (which would be at their LCA).
	// Suppose that their LCA is on level j and they are on level i. Then we want to move them both up i - j levels.
	// We can do this using the same trick as above, by greedily moving up by powers of two as long as they do not overshoot.
	// The key observation that makes this work is that we can do this even without knowing the value of i − j,
	// because we can always check whether a move would be an overshoot by checking whether the corresponding
	// ancestors are equal, and moving up only when they are not equal.
	for i := r.log2n - 1; i >= 0; i-- {
		if r.parents[u][i] != r.parents[v][i] {
			u = r.parents[u][i]
			v = r.parents[v][i]
		}
	}

	return r.parents[u][0]
}

func NewTree[T Node](root T, adjList [][]T, n int) *Tree[T] {
	log2n := int(math.Log2(float64(n)))

	parents := make([][]T, n)

	for i := 0; i < n; i++ {
		initV := make([]T, log2n)
		for j := 0; j < len(initV); j++ {
			initV[j] = -1
		}

		parents[i] = initV
	}

	levels := make([]int, n)

	var dfs func(node T, level int)

	dfs = func(node T, level int) {
		levels[node] = level

		for i := 1; i < log2n; i++ {
			if parents[node][i-1] == -1 {
				break
			}

			parents[node][i] = parents[parents[node][i-1]][i-1]
		}

		for _, child := range adjList[node] {
			parents[child][0] = node

			dfs(child, level+1)
		}
	}

	// Step: Preprocess and compute binary ancestors and levels
	// We can compute the binary ancestors and levels in our preprocessing step with a depth-first search.
	// The key observation that makes binary ancestors easy to compute is the following:
	// parents[node][i] = parents[parents[node][i-1]][i-1]
	// We can use this fact to precompute all of the binary ancestors of all vertices in O(log(n)) time.
	dfs(root, 0)

	return &Tree[T]{
		n:       n,
		log2n:   log2n,
		level:   levels,
		parents: parents,
	}
}

func pow(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}
