package GTOOL

// Comparator 用于比较两个元素的大小
// src<dst 返回-1
// src=dst 返回0
// src>dst 返回1
type Comparator[T any] func(src T, dst T) int

func ComparatorRealNumber[T RealNumber](src T, dst T) int {
	if src < dst {
		return -1
	} else if src == dst {
		return 0
	} else {
		return 1
	}
}
