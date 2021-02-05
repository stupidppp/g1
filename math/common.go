package math

// Factor ...
func Factor(i int) int {
	return i * Factor(i-1)
}
