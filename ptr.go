package GTOOL

func Toptr[T any](t T) *T {
	return &t
}
