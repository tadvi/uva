package filter

func Ints(slice []int, fn func(v int) bool) []int {
	b := slice[:0]
	for _, v := range slice {
		if fn(v) {
			b = append(b, v)
		}
	}
	return b
}

func Float64s(slice []float64, fn func(v float64) bool) []float64 {
	b := slice[:0]
	for _, v := range slice {
		if fn(v) {
			b = append(b, v)
		}
	}
	return b
}

func Strings(slice []string, fn func(v string) bool) []string {
	b := slice[:0]
	for _, v := range slice {
		if fn(v) {
			b = append(b, v)
		}
	}
	return b
}
