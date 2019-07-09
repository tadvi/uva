package functool

// N returns slice for range loop iteration.
func N(n int) []struct{} {
	return make([]struct{}, n)
}
