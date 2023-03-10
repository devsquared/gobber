package queue

import (
	"fmt"
	"github.com/devsquared/gobber/pkg/queue/test"
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestPriorityQueue_Peek(t *testing.T) {
	type scenario struct {
		name          string
		queue         *PriorityQueue
		expectedValue any
		expectedErr   error
		copiedQueue   *PriorityQueue
	}

	pQueueWithSingleItem := NewPriorityQueue()
	pQueueWithSingleItem.Push(PQItem{value: "hi", priority: 0})

	pQueueWithMultipleItems := NewPriorityQueue()
	pQueueWithMultipleItems.Push(PQItem{value: "hello", priority: 1})
	pQueueWithMultipleItems.Push(PQItem{value: "hiya", priority: 2})

	testScenarios := []scenario{
		{
			name:          "peek on empty queue",
			queue:         NewPriorityQueue(),
			expectedErr:   fmt.Errorf("priority queue: peek called on empty queue"),
			expectedValue: nil,
			copiedQueue:   NewPriorityQueue(),
		},
		{
			name:          "peek on queue with single item",
			queue:         pQueueWithSingleItem,
			expectedErr:   nil,
			expectedValue: "hi",
			copiedQueue:   pQueueWithSingleItem,
		},
		{
			name:          "peek on queue with multiple items",
			queue:         pQueueWithMultipleItems,
			expectedErr:   nil,
			expectedValue: "hiya",
			copiedQueue:   pQueueWithMultipleItems,
		},
	}

	for _, ts := range testScenarios {
		actualValue, actualErr := ts.queue.Peek()

		if actualValue != ts.expectedValue {
			test.ReportTestFailure(ts.name, actualValue, ts.expectedValue)
		}

		test.CheckErrorsAreSame(actualErr, ts.expectedErr)

		// make sure that the queue is unchanged by a peek
		if !cmp.Equal(ts.queue, ts.copiedQueue) {
			test.ReportTestFailure(ts.name, ts.queue, ts.copiedQueue)
		}
	}
}

func TestPriorityQueue_Pop(t *testing.T) {

}

func TestPriorityQueue_Push(t *testing.T) {
}
