package list

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
	panic("a")
}
