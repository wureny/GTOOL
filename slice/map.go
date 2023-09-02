package slice

func toMap[T comparable](src []T) map[T]struct{} {
	var dataMap = make(map[T]struct{}, len(src))
	for _, v := range src {
		// 使用空结构体,减少内存消耗
		dataMap[v] = struct{}{}
	}
	return dataMap
}

func deduplicate[T comparable](data []T) []T {
	ansMap := toMap(data)
	ans := make([]T, 0, len(ansMap))
	for key := range ansMap {
		ans = append(ans, key)
	}
	return ans
}
