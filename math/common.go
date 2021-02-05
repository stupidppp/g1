package math

func factor(i int) int {
	return i * factor(i-1)
}
