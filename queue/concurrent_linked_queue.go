package queue

import (
	"github.com/wureny/GTOOL/internal/queue"
	"sync/atomic"
	"unsafe"
)

type ConcurrentLinkedQueue[T any] struct {
	head unsafe.Pointer
	tail unsafe.Pointer
}

func NewConcurrentLinkedQueue[T any]() *ConcurrentLinkedQueue[T] {
	head := &node[T]{}
	return &ConcurrentLinkedQueue[T]{head: unsafe.Pointer(head), tail: unsafe.Pointer(head)}
}

func (c ConcurrentLinkedQueue[T]) Enqueue(t T) error {
	newNode := &node[T]{Val: t}
	newPtr := unsafe.Pointer(newNode)

	for {
		tailPtr := atomic.LoadPointer(&c.tail)
		tail := (*node[T])(tailPtr)
		nextPtr := atomic.LoadPointer(&tail.Next)
		if nextPtr != nil {
			continue
		}
		//更新tail.Next
		if atomic.CompareAndSwapPointer(&tail.Next, nextPtr, newPtr) {
			//若上一步成功，更新tail
			atomic.CompareAndSwapPointer(&c.tail, tailPtr, newPtr)
			return nil
		}
	}
}

func (c ConcurrentLinkedQueue[T]) Dequeue() (T, error) {
	for {
		headPtr := atomic.LoadPointer(&c.head)
		head := (*node[T])(headPtr)
		tailPtr := atomic.LoadPointer(&c.tail)
		tail := (*node[T])(tailPtr)
		if head == tail {
			var t T
			return t, queue.ErrEmptyQueue
		}
		nextPtr := atomic.LoadPointer(&head.Next)
		if atomic.CompareAndSwapPointer(&c.head, headPtr, nextPtr) {
			headNext := (*node[T])(nextPtr)
			return headNext.Val, nil
		}
	}
}
