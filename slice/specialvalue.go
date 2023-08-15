package slice

import (
	"errors"
	"github.com/wureny/GTOOL"
)

func Max[T GTOOL.RealNumber](src []T) (T, error) {
	if len(src) == 0 {
		var tmp T
		return tmp, errors.New("Error: the slice is empty!")
	}
	res := src[0]
	for i := 1; i < len(src); i++ {
		if src[i] > res {
			res = src[i]
		}
	}
	return res, nil
}

func Min[T GTOOL.RealNumber](ts []T) (T, error) {
	if len(ts) == 0 {
		var tmp T
		return tmp, errors.New("Error: the sice is empty!")
	}
	res := ts[0]
	for i := 1; i < len(ts); i++ {
		if ts[i] < res {
			res = ts[i]
		}
	}
	return res, nil
}

// Sum 求和
// 在使用 float32 或者 float64 的时候要小心精度问题
func Sum[T GTOOL.Number](ts []T) T {
	var res T
	for _, n := range ts {
		res += n
	}
	return res
}
