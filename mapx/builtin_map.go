package mapx

// 封装一个map，用作装饰器模式中被装饰的对象
type builtinMap[K comparable, V any] struct {
	data map[K]V
}

func (b *builtinMap[K, V]) Put(k K, v V) error {
	b.data[k] = v
	return nil
}

func (b *builtinMap[K, V]) Get(key K) (V, bool) {
	ans, ok := b.data[key]
	if !ok {
		var s any
		return s, false
	}
	return ans, true
}

func (b *builtinMap[K, V]) Delete(k K) (V, bool) {
	ans, ok := b.data[k]
	if !ok {
		var s any
		return s, false
	}
	delete(b.data, k)
	return ans, true
}

func (b *builtinMap[K, V]) Keys() []K {
	/*ans := make([]K, len(b.data))
	for s := range b.data {
		ans = append(ans, s)
	}
	return ans
	*/
	return keys[K, V](b.data)
}

func (b *builtinMap[K, V]) Values() []V {
	/*
		ans := make([]V, len(b.data))
		for _, val := range b.data {
			ans = append(ans, val)
		}
		return ans
	*/
	return values[K, V](b.data)
}
