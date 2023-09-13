package mapx

type mapi[K comparable, V any] interface {
	Put(K, V) error
	Get(key K) (V, bool)
	//第一个返回值是删除元素的值，第二个返回值代表
	//是否成功删除
	Delete(k K) (V, bool)
	//返回所有键的值，注意，返回值的顺序取决于具体的实现；
	Keys() []K
	//返回所有值，注意，返回值的顺序取决于具体的实现；
	Values() []V
}
