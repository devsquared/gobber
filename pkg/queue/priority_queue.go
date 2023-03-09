package queue

import (
	"fmt"
	"github.com/devsquared/gobber/pkg/queue/heap"
	"reflect"
)

type PQItem struct {
	value    any
	priority int
}

type PriorityQueue struct {
	heap  *heap.MaxHeap
	count int
}

func NewPriorityQueue() *PriorityQueue {
	return &PriorityQueue{
		heap:  heap.NewMaxHeap(),
		count: 0,
	}
}

func (q *PriorityQueue) Pop() (any, error) {
	// return error if the queue is empty
	if q.count <= 0 {
		return nil, fmt.Errorf("priority queue: pop called on empty queue")
	}

	poppedValue, err := q.heap.Pop()
	if err != nil {
		return nil, fmt.Errorf("priority queue: error in pop: %e", err)
	}

	return poppedValue, nil
}

// Push enqueues an element onto the PriorityQueue. If an element is given that is not a PQItem, a priority of 0 is given.
func (q *PriorityQueue) Push(element any) {
	var item *PQItem
	if reflect.ValueOf(element).Kind() != reflect.ValueOf(PQItem{}).Kind() {
		item = &PQItem{value: element, priority: 0}
	} else {
		item = element.(*PQItem)
	}

	newHeapNode := heap.NewNode(item.priority, item.value)
	q.heap.Add(newHeapNode)
	q.count++
}

func (q *PriorityQueue) Peek() (any, error) {
	// return error if the queue is empty
	if q.count <= 0 {
		return nil, fmt.Errorf("priority queue: peek called on empty queue")
	}

	peekValue, err := q.heap.GetFirstValue()
	if err != nil {
		return nil, fmt.Errorf("priority queue: error in peek: %e", err)
	}

	return peekValue, nil
}
