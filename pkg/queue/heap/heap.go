package heap

type heap interface {
	Add(node Node)
	Pop() (any, error)
	GetFirstValue() (any, error)
}

type Node struct {
	key   int
	value any
}

// NewNode creates a simple node structure for use in a heap.
func NewNode(key int, value any) Node {
	return Node{
		key:   key,
		value: value,
	}
}
