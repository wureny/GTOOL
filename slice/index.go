package slice

func IndexFunc[T any](src []T, match matchFunc[T]) int {
	for idx, val := range src {
		if match(val) {
			return idx
		}
	}
	return -1
}

func IndexFuncAll[T any](src []T, match matchFunc[T]) []int {
	arr := make([]int, 0, len(src)>>3+1)
	for idx, val := range src {
		if match(val) {
			arr = append(arr, idx)
		}
	}
	return arr
}

func Index[T comparable](src []T, dst T) int {
	for idx, val := range src {
		if val == dst {
			return idx
		}
	}
	return -1
}

func IndexAll[T comparable](src []T, dst T) []int {
	ans := make([]int, 0, len(src)>>3+1)
	for idx, val := range src {
		if val == dst {
			ans = append(ans, idx)
		}
	}
	return ans
}
