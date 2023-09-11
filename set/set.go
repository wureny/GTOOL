package set

type Set[T comparable] interface {
	Add(key T)
	Delete(key T)
	// Exist 返回是否存在这个元素
	Exist(key T) bool
	Keys() []T
}

type MapSet[T comparable] struct {
	m map[T]struct{}
}

func NewMapSet[T comparable](size int) *MapSet[T] {
	return &MapSet[T]{
		m: make(map[T]struct{}, size),
	}
}

func (s *MapSet[T]) Add(key T) {
	s.m[key] = struct{}{}
}

func (s *MapSet[T]) Delete(key T) {
	delete(s.m, key)
}

func (s *MapSet[T]) Exist(key T) bool {
	_, ok := s.m[key]
	return ok
}

// 注意：keys方法返回的元素顺序不确定
func (s *MapSet[T]) Keys() []T {
	ans := make([]T, len(s.m))
	for key := range s.m {
		ans = append(ans, key)
	}
	return ans
}
