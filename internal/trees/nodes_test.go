package trees

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDepthFirstTraversal(t *testing.T) {
	leaf1 := Leaf[int](1)
	leaf2 := Leaf[int](2)
	leaf3 := Leaf[int](10)
	parent := Node[int]{val: 3, children: []Node[int]{leaf1, leaf2, leaf3}}
	root := Node[int]{val: 0, children: []Node[int]{parent}}
	var xs []int
	root.Dfs(func(node *Node[int]) {
		xs = append(xs, node.val)
	})
	//         0
	//       3
	//  1  2  10
	assert.Equal(t, []int{1, 2, 10, 3, 0}, xs)
}
