package list

import (
	"github.com/wureny/GTOOL/internal/errs"
)

type node[T any] struct {
	prev *node[T]
	next *node[T]
	val  T
}

type LinkedList[T any] struct {
	head   *node[T]
	tail   *node[T]
	length int
}

func NewLinkedList[T any]() *LinkedList[T] {
	head := &node[T]{}
	tail := &node[T]{next: head, prev: head}
	head.next, head.prev = tail, tail
	ans := &LinkedList[T]{
		head: head,
		tail: tail,
		//ekit中没有下面一行代码 ？？？
		//length: 2,
	}
	return ans
}

func (l *LinkedList[T]) checkIndex(index int) bool {
	return 0 <= index && index < l.Len()
}

func (l *LinkedList[T]) Len() int {
	return l.length
}

// findNode 使用之前必须确保index有效
func (l *LinkedList[T]) findNode(index int) *node[T] {
	var cur *node[T]
	if index < l.Len()/2 {
		cur = l.head
		for i := -1; i < index; i++ {
			cur = cur.next
		}
	} else {
		cur = l.tail
		for i := l.Len(); i > index; i-- {
			cur = cur.prev
		}
	}
	return cur
}

func (l *LinkedList[T]) Get(index int) (T, error) {
	if !l.checkIndex(index) {
		var tmp T
		return tmp, errs.ErrIndexOutOfRange(l.length, index)
	}
	return l.findNode(index).val, nil
}

func (l *LinkedList[T]) Set(index int, t T) error {
	if !l.checkIndex(index) {
		return errs.ErrIndexOutOfRange(l.Len(), index)
	}
	node := l.findNode(index)
	node.val = t
	return nil
}

func (l *LinkedList[T]) Delete(index int) (T, error) {
	if !l.checkIndex(index) {
		var zeroValue T
		return zeroValue, errs.ErrIndexOutOfRange(l.Len(), index)
	}
	node := l.findNode(index)
	node.prev.next = node.next
	node.next.prev = node.prev
	node.prev = nil
	node.next = nil
	l.length--
	return node.val, nil
}

func (l *LinkedList[T]) Append(ts ...T) error {
	for _, val := range ts {
		node := &node[T]{
			prev: l.tail.prev,
			next: l.tail,
			val:  val,
		}
		node.prev.next, node.next.prev = node, node
		l.length++
	}
	return nil
}

func (l *LinkedList[T]) Add(index int, ts T) error {
	if index < 0 || index > l.Len() {
		return errs.ErrIndexOutOfRange(l.length, index)
	}
	if index == l.length {
		return l.Append(ts)
	}
	n := l.findNode(index)
	node := &node[T]{
		prev: n.prev,
		next: n,
		val:  ts,
	}
	n.prev.next = node
	n.prev = node
	l.length++
	return nil
}

func (l *LinkedList[T]) Range(index int, fn func(index int, t T) error) error {
	if !l.checkIndex(index) {
		return errs.ErrIndexOutOfRange(l.length, index)
	}
	node := l.head.next
	for i := 1; i <= index; i++ {
		node = node.next
	}
	for cur, i := node, index; i < l.length; i++ {
		err := fn(i, cur.val)
		if err != nil {
			return err
		}
		cur = cur.next
	}
	return nil
}

func (l *LinkedList[T]) AsSlice() []T {
	ans := make([]T, 0, l.length)
	if l.length == 0 {
		return ans
	}
	for node, i := l.head.next, 0; i < l.length; i++ {
		ans = append(ans, node.val)
		node = node.next
	}
	return ans
}
