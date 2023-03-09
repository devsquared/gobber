package heap

import "fmt"

// For ease of implementation, we will use a simple slice or array implementation for a tree.
// This means that we follow these common rules for indices:
// - *Parent Index*: (i - 1) / 2
// - *Children Indices*
//		- Left Child: 2 * i + 1
// 		- Right Child: 2 * i + 2

// MaxHeap represents a heap with the max values towards the top
type MaxHeap struct {
	heap []Node
}

// NewMaxHeap is a simple constructor to get a max heap.
func NewMaxHeap() *MaxHeap {
	return &MaxHeap{
		heap: []Node{},
	}
}

// Add inserts a new node into the MaxHeap. After adding, the MaxHeap fixes the remaining nodes ordering.
func (b *MaxHeap) Add(node Node) {
	b.heap = append(b.heap, node)
	b.bubbleUp(len(b.heap) - 1) // bubble up the new node
}

// Pop removes the max keyed node from the MaxHeap. After removing, the MaxHeap fixes the remaining nodes ordering.
func (b *MaxHeap) Pop() (any, error) {
	if len(b.heap) <= 0 {
		return nil, fmt.Errorf("max heap: pop called on empty heap")
	}

	removed := b.heap[0]
	heapSize := len(b.heap)

	if heapSize > 1 {
		b.heap[0] = b.heap[len(b.heap)-1]
	}

	b.heap = b.heap[:len(b.heap)-1]
	b.bubbleDown(0)

	return removed.value, nil
}

// GetFirstValue returns the value from the max node of the heap. This does not remove this node from the heap.
// Similar to a peek in a queue.
func (b *MaxHeap) GetFirstValue() (any, error) {
	if len(b.heap) <= 0 {
		return nil, fmt.Errorf("max heap: get first value called on empty heap")
	}

	return b.heap[0].value, nil
}

func (b *MaxHeap) bubbleUp(index int) {
	for index > 0 {
		parentIndex := getParentIndex(index)

		if b.heap[parentIndex].key > b.heap[index].key {
			// the node is now in the correct place and is bubbled up; we are done
			return
		}

		b.heap[parentIndex], b.heap[index] = b.heap[index], b.heap[parentIndex] //swap the nodes
		index = parentIndex
	}
}

func (b *MaxHeap) bubbleDown(index int) {
	for 2*index+1 < len(b.heap) {
		minChildIndex := b.maxChildIndex(index)

		if b.heap[minChildIndex].key < b.heap[index].key {
			// the node is now in the correct place and is bubbled down; we are done
			return
		}

		b.heap[minChildIndex], b.heap[index] = b.heap[index], b.heap[minChildIndex]
		index = minChildIndex
	}
}

func (b *MaxHeap) maxChildIndex(index int) int {
	if getRightIndex(index) >= len(b.heap) {
		return getLeftIndex(index)
	}

	if b.heap[getRightIndex(index)].key > b.heap[getLeftIndex(index)].key {
		return getRightIndex(index)
	}

	return getLeftIndex(index)
}

func getParentIndex(index int) int {
	return (index - 1) / 2
}

func getLeftIndex(index int) int {
	return 2*index + 1
}

func getRightIndex(index int) int {
	return 2*index + 2
}
