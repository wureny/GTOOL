package slice

import "github.com/wureny/GTOOL/internal/errs"

func Add[T any](src []T, ele T, index int) ([]T, error) {
	le := len(src)
	if index < 0 || index >= le {
		return nil, errs.ErrIndexOutOfRange(le, index)
	}
	var aele T
	src = append(src, aele)
	copy(src[index+1:], src[index:len(src)-1])
	src[index] = ele
	return src, nil
}
