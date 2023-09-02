package slice

func UnionSet[T comparable](src, dst []T) []T {
	srcMap, dstMap := toMap(src), toMap(dst)
	for key := range srcMap {
		dstMap[key] = struct{}{}
	}
	ans := make([]T, 0, len(dstMap))
	for key := range dstMap {
		ans = append(ans, key)
	}
	return ans
}
