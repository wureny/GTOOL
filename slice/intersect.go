package slice

func IntersectSet[T comparable](src, dst []T) []T {
	tmpLen := 0
	if len(src) > len(dst) {
		tmpLen = len(dst)
	} else {
		tmpLen = len(src)
	}
	ans := make([]T, 0, tmpLen)
	srcMap := toMap(src)
	for _, val := range dst {
		if _, exit := srcMap[val]; exit {
			ans = append(ans, val)
		}
	}
	return deduplicate[T](ans)
}
