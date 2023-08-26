package slice

func Contains[T comparable](src []T, dst T) ([]int, bool) {
	return ContainsFunc(src, func(src T) bool {
		return src == dst
	})
}

func ContainsFunc[T any](src []T, equal func(src T) bool) ([]int, bool) {
	arr := make([]int, 0, 1)
	for idx, val := range src {
		if equal(val) {
			arr = append(arr, idx)
		}
	}
	if len(arr) == 0 {
		var tmp []int
		return tmp, false
	}
	return arr, true
}

func ContainsAny[T comparable](src []T, dst []T) bool {
	srcMap := toMap[T](src)
	for _, val := range dst {
		if _, exist := srcMap[val]; exist {
			return true
		}
	}
	return false
}

func ContainsAll[T comparable](src, dst []T) bool {
	srcMap := toMap[T](src)
	for _, v := range dst {
		if _, exist := srcMap[v]; !exist {
			return false
		}
	}
	return true
}
