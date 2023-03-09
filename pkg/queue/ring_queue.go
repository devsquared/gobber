package queue

import "fmt"

// Basic ring queue implementation with a basic slice of interface{}.
// With this implementation of a ring queue, we utilize bit-masking to speed up the processes.
// This means that it is important to keep the size of the buffer a power of 2.

// minRingQueueSize starts at 16 and must be a power of 2.
const minRingQueueSize = 16

// RingQueue represents the ring buffer queue.
type RingQueue struct {
	buffer []any
	head   int // marker of the head in the slice
	tail   int // marker of the tail in the slice
	count  int // length of the queues contents
}

// NewRingQueue constructs a new RingQueue instance.
func NewRingQueue() *RingQueue {
	return &RingQueue{
		buffer: make([]any, minRingQueueSize), //create a buffer the minimum size to begin
	}
}

// Length gets the length of the queue currently.
func (q *RingQueue) Length() int {
	return q.count
}

// resize handles resizing the queue whenever it is needed. This will either double its length if space is needed
// or it will shrink the size if the queue is less than half full.
func (q *RingQueue) resize() {
	//start by doubling the size
	newBuffer := make([]any, q.count<<1)

	//now appropriately copy the contents
	if q.tail > q.head {
		copy(newBuffer, q.buffer[q.head:q.tail])
	} else {
		n := copy(newBuffer, q.buffer[q.head:])
		copy(newBuffer[n:], q.buffer[:q.tail])
	}

	//now reset the values in the newly resized queue instance
	q.head = 0
	q.tail = q.count
	q.buffer = newBuffer
}

// Push enqueues a new element on to the end of the queue.
func (q *RingQueue) Push(element any) {
	// if we have run out of room, let's resize
	if q.count == len(q.buffer) {
		q.resize()
	}

	q.buffer[q.tail] = element
	q.tail = (q.tail + 1) & (len(q.buffer) - 1) //bitwise modulus using AND
	q.count++
}

// Peek provides utility to see the front of the queue. Returns an error whenever the queue is empty.
func (q *RingQueue) Peek() (any, error) {
	// if the queue is empty, error
	if q.count <= 0 {
		return nil, fmt.Errorf("peek attempted on empty queue")
	}
	return q.buffer[q.head], nil
}

// Pop dequeues the element from the front of the queue and returns it. If the queue is empty, an error is returned.
func (q *RingQueue) Pop() (any, error) {
	// if the queue is empty, error
	if q.count <= 0 {
		return nil, fmt.Errorf("pop attempted on empty queue")
	}
	result := q.buffer[q.head]                  // get the result
	q.buffer[q.head] = nil                      // remove result from queue
	q.head = (q.head + 1) & (len(q.buffer) - 1) // bitwise modulus using AND
	q.count--

	// if buffer is bigger than minimum size and 1/4 full, resize
	if len(q.buffer) > minRingQueueSize && (q.count<<2) == len(q.buffer) {
		q.resize()
	}

	return result, nil
}
