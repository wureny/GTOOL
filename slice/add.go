package slice

import "github.com/wureny/GTOOL/internal/slice"

func Add[T any](src []T, elm T, index int) ([]T, error) {
	res, err := slice.Add(src, elm, index)
	return res, err
}
