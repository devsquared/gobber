package heap

import (
	"fmt"
	"testing"
)

func TestBinaryHeap_Add(t *testing.T) {
	binaryHeap := NewMaxHeap()

	binaryHeap.Add(node{key: 3, value: "yee"})
	binaryHeap.Add(node{key: 1, value: "yeet"})
	binaryHeap.Add(node{key: 5, value: "haw"})
	binaryHeap.Add(node{key: 2, value: "alsdf"})
	binaryHeap.Add(node{key: 8, value: "ok"})
	fmt.Println(binaryHeap)

	for len(binaryHeap.heap) > 0 {
		remove, _ := binaryHeap.Pop()

		fmt.Println(remove)
	}
}
