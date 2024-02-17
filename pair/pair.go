package pair

type Pair[K any, V any] struct {
	Key   K
	Value V
}

func NewPair[K any, V any](key K, value V) *Pair[K, V] {
	return &Pair[K, V]{Key: key, Value: value}
}

func (p *Pair[K, V]) GetKey() K {
	return p.Key
}

func (p *Pair[K, V]) GetValue() V {
	return p.Value
}

func (p *Pair[K, V]) SetKey(key K) {
	p.Key = key
}

func (p *Pair[K, V]) SetValue(value V) {
	p.Value = value
}

func (p *Pair[K, V]) Set(key K, value V) {
	p.Key = key
	p.Value = value
}

func (p *Pair[K, V]) Clone() *Pair[K, V] {
	return &Pair[K, V]{Key: p.Key, Value: p.Value}
}

func (p *Pair[K, V]) Split(pair *Pair[K, V]) (K, V) {
	return pair.Key, pair.Value
}
