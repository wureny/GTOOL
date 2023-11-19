package queue

import (
	"errors"
	"github.com/wureny/GTOOL"
)

var (
	ErrOutOfCapacity = errors.New("GTOOL: 超出最大容量限制")
	ErrEmptyQueue    = errors.New("GTOOL: 队列为空")
)

type PriorityQueue[T any] struct {
	// 比较器
	comp GTOOL.Comparator[T]
	// 队列容量
	capacity int
	//队列中的元素，为便于计算父子节点的index，0位置留空，根节点从1开始
	data []T
}

// NewPriorityQueue 创建优先队列 capacity <= 0 时，为无界队列，否则有有界队列
func NewPriorityQueue[T any](capacity int, compare GTOOL.Comparator[T]) *PriorityQueue[T] {
	sliceCap := capacity + 1
	if capacity < 1 {
		capacity = 0
		sliceCap = 64
	}
	return &PriorityQueue[T]{
		capacity: capacity,
		data:     make([]T, 1, sliceCap),
		comp:     compare,
	}
}

func (p *PriorityQueue[T]) Length() int {
	return len(p.data) - 1
}

func (p *PriorityQueue[T]) Capacity() int {
	return p.capacity
}

func (p *PriorityQueue[T]) IsBoundless() bool {
	return p.capacity <= 0
}

func (p *PriorityQueue[T]) isFull() bool {
	return p.capacity > 0 && len(p.data)-1 == p.capacity
}
func (p *PriorityQueue[T]) isEmpty() bool {
	return len(p.data) == 1
}

func (p *PriorityQueue[T]) Peek() (T, error) {
	if p.isEmpty() {
		var t T
		return t, ErrEmptyQueue
	}
	return p.data[1], nil
}

// Enqueue 入队
func (p *PriorityQueue[T]) Enqueue(t T) error {
	if p.isFull() {
		return ErrOutOfCapacity
	}
	p.data = append(p.data, t)
	node, parent := len(p.data)-1, len(p.data)/2
	for parent > 0 && p.comp(p.data[node], p.data[parent]) < 0 {
		p.data[node], p.data[parent] = p.data[parent], p.data[node]
		node, parent = parent, parent/2
	}
	return nil
}

// Dequeue 出队
func (p *PriorityQueue[T]) Dequeue() (T, error) {
	if p.isEmpty() {
		var t T
		return t, ErrEmptyQueue
	}

	pop := p.data[1]
	p.data[1] = p.data[len(p.data)-1]
	p.data = p.data[:len(p.data)-1]
	p.shrinkIfNecessary()
	p.heapify(p.data, len(p.data)-1, 1)
	return pop, nil
}

// shrinkIfNecessary 如果队列中的元素个数小于容量的1/4，则缩容
func (p *PriorityQueue[T]) shrinkIfNecessary() {
	if p.IsBoundless() {
		//TODO 无界队列暂不支持缩容
		//p.data = slice.Shrink[T](p.data)
	}
}

// heapify 从上往下堆化
func (p *PriorityQueue[T]) heapify(data []T, n, i int) {
	minPos := i
	for {
		if left := i * 2; left <= n && p.comp(data[left], data[minPos]) < 0 {
			minPos = left
		}
		if right := i*2 + 1; right <= n && p.comp(data[right], data[minPos]) < 0 {
			minPos = right
		}
		if minPos == i {
			break
		}
		data[i], data[minPos] = data[minPos], data[i]
		i = minPos
	}
}
