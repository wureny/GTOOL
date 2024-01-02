package queue

import (
	"github.com/wureny/GTOOL"
	"github.com/wureny/GTOOL/internal/queue"
	"sync"
)

type ConcurrentPriorityQueue[T any] struct {
	pq queue.PriorityQueue[T]
	m  sync.RWMutex
}

func (c *ConcurrentPriorityQueue[T]) Enqueue(t T) error {
	c.m.Lock()
	defer c.m.Unlock()
	return c.pq.Enqueue(t)
}

func (c *ConcurrentPriorityQueue[T]) Dequeue() (T, error) {
	c.m.Lock()
	defer c.m.Unlock()
	return c.pq.Dequeue()
}

func (c *ConcurrentPriorityQueue[T]) Peek() (T, error) {
	c.m.Lock()
	defer c.m.Unlock()
	return c.pq.Peek()
}

func NewConcurrentPriorityQueue[T any](capacity int, compare GTOOL.Comparator[T]) *ConcurrentPriorityQueue[T] {
	return &ConcurrentPriorityQueue[T]{pq: *queue.NewPriorityQueue[T](capacity, compare)}
}

func (c *ConcurrentPriorityQueue[T]) Len() int {
	c.m.RLock()
	defer c.m.RUnlock()
	return c.pq.Length()
}
func (c *ConcurrentPriorityQueue[T]) Cap() int {
	c.m.RLock()
	defer c.m.RUnlock()
	return c.pq.Capacity()
}
