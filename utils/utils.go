package utils

func Ternary[T any](condition bool, result1 T, result2 T) T {
	if condition {
		return result1
	} else {
		return result2
	}
}
