package math

// Factor ...
func Factor(i int) int {
	if i <= 1 {
		return 1
	}
	return i * Factor(i-1)
}
