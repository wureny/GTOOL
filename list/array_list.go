package list

import "github.com/wureny/GTOOL/internal/errs"

var (
	_ List[int] = (*ArrayList[int])(nil)
)

type ArrayList[T any] struct {
	data []T
}

// NewArrayList 初始化一个len为0，cap为cap的ArrayList
func NewArrayList[T any](cap int) *ArrayList[T] {
	return &ArrayList[T]{data: make([]T, 0, cap)}
}

// NewArrayListOf 直接使用 ts，而不会执行复制
func NewArrayListOf[T any](ts []T) *ArrayList[T] {
	return &ArrayList[T]{
		data: ts,
	}
}

func (a ArrayList[T]) Get(index int) (T, error) {
	if index < 0 || index >= len(a.data) {
		var s T
		return s, errs.ErrIndexOutOfRange(len(a.data), index)
	}
	return a.data[index], nil
}

func (a ArrayList[T]) Append(ems ...T) error {
	//TODO implement me
	a.data = append(a.data, ems...)
	return nil
}

func (a ArrayList[T]) Add(index int, em T) error {
	//TODO implement me
	if index < 0 || index > len(a.data) {
		return errs.ErrIndexOutOfRange(len(a.data), index)
	}
	a.data = append(a.data, em)
	copy(a.data[index+1:], a.data[index:])
	a.data[index] = em
	return nil
}

func (a ArrayList[T]) Set(index int, t T) error {
	length := len(a.data)
	if index >= length || index < 0 {
		return errs.ErrIndexOutOfRange(length, index)
	}
	a.data[index] = t
	return nil
}

func (a ArrayList[T]) Delete(index int) (T, error) {
	//TODO implement me
	panic("implement me")
}

func (a ArrayList[T]) Len() int {
	return len(a.data)
}

func (a ArrayList[T]) Cap() int {
	return cap(a.data)
}

func (a ArrayList[T]) Range(index int, fn func(index int, t T) error) error {
	for key, value := range a.data {
		e := fn(key, value)
		if e != nil {
			return e
		}
	}
	return nil
}

func (a ArrayList[T]) AsSlice() []T {
	res := make([]T, len(a.data))
	copy(res, a.data)
	return res

}
