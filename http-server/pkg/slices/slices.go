package slices

// Equal indicates equality of two num slices.
func Equal[T byte | int | float32 | float64](x []T, y []T) bool {
	if len(x) != len(y) {
		return false
	}

	for i, xi := range x {
		if xi != y[i] {
			return false
		}
	}

	return true
}
