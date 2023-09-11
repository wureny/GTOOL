package list

type List[T any] interface {
	Get(index int) (T, error)
	Append(ems ...T) error
	Add(index int, em T) error
	Set(index int, t T) error
	Delete(index int) (T, error)
	Len() int
	Cap() int
	Range(index int, fn func(index int, t T) error) error
	// AsSlice 将 List 转化为一个切片
	// 不允许返回nil，在没有元素的情况下，
	// 必须返回一个长度和容量都为 0 的切片
	// AsSlice 每次调用都必须返回一个全新的切片
	AsSlice() []T
}
