package heap

import (
	"fmt"
	"reflect"
	"testing"
)

// When testing the heap, it is important to understand how the slice implementation works. To test that the
// heap is in a good state at any given moment, we can directly check the underlying slice.

func TestBinaryHeap_Add(t *testing.T) {
	type scenario struct {
		name         string
		startingHeap *MaxHeap
		input        Node
		expected     []Node
	}

	maxHeapWithLow := NewMaxHeap()
	maxHeapWithLow.Add(NewNode(0, "testing!"))

	maxHeapWithHigh := NewMaxHeap()
	maxHeapWithHigh.Add(NewNode(99, "testing it all!"))

	maxHeapWithSpread := NewMaxHeap()
	maxHeapWithSpread.Add(NewNode(0, "yee"))
	maxHeapWithSpread.Add(NewNode(99, "haw"))

	testScenarios := []scenario{
		{
			name:         "add a node to an empty heap",
			startingHeap: NewMaxHeap(),
			input:        NewNode(1, "hello!"),
			expected:     []Node{NewNode(1, "hello!")},
		},
		{
			name:         "add a node to a heap with key lower than rest",
			startingHeap: maxHeapWithHigh,
			input:        NewNode(0, "konichiwa!"),
			expected:     []Node{NewNode(99, "testing it all!"), NewNode(0, "konichiwa!")},
		},
		{
			name:         "add a node to a heap with key higher than rest",
			startingHeap: maxHeapWithLow,
			input:        NewNode(999999, "woah"),
			expected:     []Node{NewNode(999999, "woah"), NewNode(0, "testing!")},
		},
		{
			name:         "add a node to a heap with a key in between the rest",
			startingHeap: maxHeapWithSpread,
			input:        NewNode(50, "middle"),
			expected:     []Node{NewNode(99, "haw"), NewNode(0, "yee"), NewNode(50, "middle")},
		},
	}

	for _, ts := range testScenarios {
		// add the new input node
		ts.startingHeap.Add(ts.input)

		actualHeap := ts.startingHeap.heap
		if !reflect.DeepEqual(actualHeap, ts.expected) {
			t.Fatalf(reportTestFailure(ts.name, actualHeap, ts.expected))
		}
	}
}

func TestMaxHeap_Pop(t *testing.T) {
	type scenario struct {
		name                  string
		startingHeap          *MaxHeap
		expectedValue         any
		expectedErr           error
		expectedRemainingHeap []Node
	}

	maxHeapWithLow := NewMaxHeap()
	maxHeapWithLow.Add(NewNode(0, "testing!"))

	maxHeapWithHigh := NewMaxHeap()
	maxHeapWithHigh.Add(NewNode(99, "testing it all!"))

	maxHeapWithSpread := NewMaxHeap()
	maxHeapWithSpread.Add(NewNode(0, "yee"))
	maxHeapWithSpread.Add(NewNode(99, "haw"))

	testScenarios := []scenario{
		{
			name:                  "pop on an empty heap",
			startingHeap:          NewMaxHeap(),
			expectedErr:           fmt.Errorf("max heap: pop called on empty heap"),
			expectedRemainingHeap: []Node{},
		},
		{
			name:                  "pop on a single node queue",
			startingHeap:          maxHeapWithLow,
			expectedValue:         "testing!",
			expectedRemainingHeap: []Node{},
		},
		{
			name:                  "pop with multiple possible nodes",
			startingHeap:          maxHeapWithSpread,
			expectedValue:         "haw",
			expectedRemainingHeap: []Node{NewNode(0, "yee")},
		},
	}

	for _, ts := range testScenarios {
		// add the new input node
		actualValue, actualErr := ts.startingHeap.Pop()

		if actualErr != nil && ts.expectedErr != nil {
			if actualErr.Error() != ts.expectedErr.Error() {
				t.Fatalf(reportTestFailure(ts.name, actualErr, ts.expectedErr))
			}
		}

		if actualValue != ts.expectedValue {
			t.Fatalf(reportTestFailure(ts.name, actualValue, ts.expectedValue))
		}

		actualHeap := ts.startingHeap.heap
		if !reflect.DeepEqual(actualHeap, ts.expectedRemainingHeap) {
			t.Fatalf(reportTestFailure(ts.name, actualHeap, ts.expectedRemainingHeap))
		}
	}
}

// likely to move out to a test package if we continue testing this way
func reportTestFailure(scenarioName string, got, wanted any) string {
	return fmt.Sprintf("scenario: %s \n\t got: %v, wanted: %v", scenarioName, got, wanted)
}
