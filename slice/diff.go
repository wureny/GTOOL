package slice

func DiffSet[T comparable](src, dst []T) []T {
	srcmap := toMap[T](src)
	for _, val := range dst {
		delete(srcmap, val)
	}
	ans := make([]T, 0, len(srcmap))
	for key := range srcmap {
		ans = append(ans, key)
	}
	return ans
}
