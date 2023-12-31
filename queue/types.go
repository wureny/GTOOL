package queue

import (
	"context"
	"unsafe"
)

// BlockingQueue blocking queue
// Reference Queue ordinary queue
// 一个阻塞队列是否遵循 FIFO 取决于具体实现
type BlockingQueue[T any] interface {
	// Enqueue puts elements into the queue. If there is a free position in the queue before ctx times out, the element will be put into the queue;
	// Otherwise return error.
	// In the case of timeout or caller's initiative to cancel, all implementations must return ctx.
	// The caller can check whether error is context.DeadlineExceeded
	// Or context.Canceled to determine the reason for failure to join the queue
	// Note that the caller must use errors.Is to determine and cannot directly use ==
	Enqueue(ctx context.Context, t T) error
	// Dequeue gets an element from the head of the queue
	// If there are elements in the queue before ctx times out, the element at the head of the queue will be returned, otherwise an error will be returned.
	// In the case of timeout or caller's initiative to cancel, all implementations must return ctx.
	// The caller can check whether error is context.DeadlineExceeded
	// Or context.Canceled to determine the reason for failure to join the queue
	// Note that the caller must use errors.Is to determine and cannot directly use ==
	Dequeue(ctx context.Context) (T, error)
}

// Queue ordinary queue
// Refer to BlockingQueue blocking queue
// Whether a queue follows FIFO depends on the implementation
type Queue[T any] interface {
	// Enqueue puts elements into the queue. If the queue is full at this time, an error is returned.
	Enqueue(t T) error
	// Dequeue gets an element from the head of the queue
	// If there are no elements in the queue at this time, then return an error
	Dequeue() (T, error)
}

type node[T any] struct {
	Val  T
	Next unsafe.Pointer
}
