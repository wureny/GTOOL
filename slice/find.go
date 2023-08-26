package slice

func Find[T any](src []T, match matchFunc[T]) (T, bool) {
	for _, val := range src {
		if match(val) {
			return val, true
		}
	}
	var s T
	return s, false
}

func FindAll[T any](src []T, match matchFunc[T]) []T {
	//注意其中的扩容机制
	ans := make([]T, 0, len(src)>>3+1)
	for _, val := range src {
		if match(val) {
			ans = append(ans, val)
		}
	}
	return ans
}
