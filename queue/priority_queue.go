package queue

import (
	"github.com/wureny/GTOOL"
	"github.com/wureny/GTOOL/internal/queue"
)

type PriorityQueue[T any] struct {
	priorityQueue *queue.PriorityQueue[T]
}

func NewPriorityQueue[T any](capacity int, compare GTOOL.Comparator[T]) *PriorityQueue[T] {
	pq := &PriorityQueue[T]{}
	pq.priorityQueue = queue.NewPriorityQueue[T](capacity, compare)
	return pq
}

func (p *PriorityQueue[T]) Length() int {
	return p.priorityQueue.Length()
}

func (p *PriorityQueue[T]) Capacity() int {
	return p.priorityQueue.Capacity()
}

func (p *PriorityQueue[T]) IsBoundless() bool {
	return p.priorityQueue.IsBoundless()
}

func (p *PriorityQueue[T]) Enqueue(val T) error {
	return p.priorityQueue.Enqueue(val)
}

func (p *PriorityQueue[T]) Dequeue() (T, error) {
	return p.priorityQueue.Dequeue()
}

func (p *PriorityQueue[T]) Peek() (T, error) {
	return p.priorityQueue.Peek()
}
