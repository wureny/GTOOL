package slice

import "GTOOL/internal/errs"

func Delete[T any](src []T, index int) ([]T, T, error) {
	le := len(src)
	if index < 0 || index >= le {
		var z T
		return nil, z, errs.ErrIndexOutOfRange(le, index)
	}
	ele := src[index]
	copy(src[index:le-1], src[index+1:])
	src = src[:le-1]
	return src, ele, nil
}
