package slice

import "github.com/wureny/GTOOL/internal/slice"

// 删除特定下标的元素
func Delete[T any](src []T, index int) ([]T, error) {
	res, _, err := slice.Delete(src, index)
	return res, err
}

// 删除满足符合条件的元素
func FilterDelete[T any](src []T, f func(index int, src T) bool) []T {
	Pos := 0
	for idx, _ := range src {
		if f(idx, src[idx]) {
			continue
		}
		src[Pos] = src[idx]
		Pos++
	}
	return src[:Pos]
}
