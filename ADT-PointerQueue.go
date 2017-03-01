package leetcode

import (
	"log"
	"unsafe"
)

const (
	queuesize = 100
	arraysize = queuesize + 1
)

// PointerQueue is a array queue of pointer
type PointerQueue struct {
	base        [arraysize]unsafe.Pointer
	read, write int
}

// Enque add a pointer at PointerQueue[write],
// on condition that q is not Full
func (q *PointerQueue) Enque(ptr unsafe.Pointer) {
	// check full
	if q.Full() {
		log.Println("Enque to a full queue.")
		return
	}
	q.base[q.write] = ptr
	q.write = (q.write + 1) % arraysize
}

// Deque return the pointer at PointerQueue[read],
// on condition that q is not empty
func (q *PointerQueue) Deque() unsafe.Pointer {
	// check empty
	if q.Empty() {
		log.Println("Deque form an empty queue")
		return nil
	}
	ptr := q.base[q.read]
	q.read = (q.read + 1) % arraysize
	return ptr
}

func (q PointerQueue) Front() unsafe.Pointer {
	if q.Empty() {
		return nil
	}
	return q.base[q.read]
}

func (q PointerQueue) Empty() bool {
	return q.read == q.write
}

func (q PointerQueue) Full() bool {
	return q.read == (q.write+1)%arraysize
}
