package slice

func NewReserve[T comparable](src []T) []T {
	ans := make([]T, 0, len(src))
	for i := len(src) - 1; i >= 0; i-- {
		ans = append(ans, src[i])
	}
	return ans
}

func Reserve[T comparable](src []T) {
	for i, j := 0, len(src)-1; i < j; i, j = i+1, j-1 {
		src[i], src[j] = src[j], src[i]
	}
}
