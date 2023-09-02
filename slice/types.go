package slice

// 判断元素是否满足条件
type matchFunc[T any] func(src T) bool

// 判断两个元素是否相等
type equalFunc[T any] func(src, dst T) bool
