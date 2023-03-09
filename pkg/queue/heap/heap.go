package heap

type heap interface {
	Add(node node)
	Pop() (any, error)
	GetFirstValue() (any, error)
}

type node struct {
	key   int
	value any
}

// NewNode creates a simple node structure for use in a heap.
func NewNode(key int, value any) node {
	return node{
		key:   key,
		value: value,
	}
}
