package graphs

type Node[T any] struct {
	val      T
	children []Node[T]
}

func Leaf[T any](val T) Node[T] {
	return Node[T]{val: val}
}

func (node *Node[T]) Dfs(f func(*Node[T])) {
	for idx := range node.children {
		var child *Node[T] = &node.children[idx]
		child.Dfs(f)
	}
	f(node)
}
