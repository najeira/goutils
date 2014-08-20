// based on https://github.com/eapache/queue
// Copyright (c) 2014 Evan Huus
package queue

import (
	"sync"
)

const minQueueLen = 16

// Queue represents a single instance of the queue data structure.
type Queue struct {
	buf               []interface{}
	head, tail, count int
	mu                sync.Mutex
}

// New constructs and returns a new Queue.
func New() *Queue {
	return &Queue{buf: make([]interface{}, minQueueLen)}
}

// Length returns the number of elements currently stored in the queue.
func (q *Queue) Length() int {
	q.mu.Lock()
	defer q.mu.Unlock()
	return q.count
}

func (q *Queue) resize() {
	newBuf := make([]interface{}, q.count*2)
	if q.tail > q.head {
		copy(newBuf, q.buf[q.head:q.tail])
	} else {
		copy(newBuf, q.buf[q.head:len(q.buf)])
		copy(newBuf[len(q.buf)-q.head:], q.buf[:q.tail])
	}
	q.head = 0
	q.tail = q.count
	q.buf = newBuf
}

// Add puts an element on the end of the queue.
func (q *Queue) Add(elem interface{}) {
	q.mu.Lock()
	defer q.mu.Unlock()
	if q.count == len(q.buf) {
		q.resize()
	}
	q.buf[q.tail] = elem
	q.tail = (q.tail + 1) % len(q.buf)
	q.count++
}

// Peek returns the element at the head of the queue. This call panics
// if the queue is empty.
func (q *Queue) Peek() interface{} {
	q.mu.Lock()
	defer q.mu.Unlock()
	if q.count <= 0 {
		panic("queue: empty queue")
	}
	return q.buf[q.head]
}

// Pop returns the element at the head of the queue and removes the element
// from the front of the queue. If you actually want the element, call Peek
// first. This call panics if the queue is empty.
func (q *Queue) Pop() interface{} {
	q.mu.Lock()
	defer q.mu.Unlock()
	if q.count <= 0 {
		panic("queue: empty queue")
	}
	ret := q.buf[q.head]
	q.buf[q.head] = nil
	q.head = (q.head + 1) % len(q.buf)
	q.count--
	if len(q.buf) > minQueueLen && q.count*4 <= len(q.buf) {
		q.resize()
	}
	return ret
}
