package mapx

// import "github.com/wureny/GTOOL/syncx"
type node[T Hashable, ValType any] struct {
	key  Hashable
	val  ValType
	next *node[T, ValType]
}

type Hashable interface {
	// Code 返回该元素的哈希值
	// 注意：哈希值应该尽可能的均匀以避免冲突
	Code() uint64
	//Equals 用于比较两个元素是否相等
	Equals(key any) bool
}
