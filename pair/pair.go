package pair

type Pair[K any, V any] struct {
	Key   K
	Value V
}

func NewPair[K any, V any](key K, value V) *Pair[K, V] {
	return &Pair[K, V]{Key: key, Value: value}
}
