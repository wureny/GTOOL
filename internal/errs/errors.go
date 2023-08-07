package errs

import "fmt"

func ErrIndexOutOfRange(length int, index int) error {
	return fmt.Errorf("GTOOL: index out of range: length:%d,index: %d", length, index)
}

func ErrInvalidType(want, got string) error {
	return fmt.Errorf("GTOOL:type change failed! want : %s, got : %s", want, got)
}
