package slice

func SymmetricDiffSet[T comparable](src, dst []T) []T {
	srcMap, dstMap := toMap(src), toMap(dst)
	for key := range dstMap {
		if _, exit := srcMap[key]; exit {
			delete(srcMap, key)
			delete(dstMap, key)
		}
	}
	for k, v := range dstMap {
		srcMap[k] = v
	}
	ans := make([]T, 0, len(srcMap))
	for key := range srcMap {
		ans = append(ans, key)
	}
	return ans
}
