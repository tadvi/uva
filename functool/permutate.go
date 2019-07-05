package functool

func permutate(slice []int, pos int, fn func(slice []int)) {
	if pos == len(slice)-1 {
		fn(slice)
		return
	}

	for i := pos; i < len(slice); i++ {
		slice[i], slice[pos] = slice[pos], slice[i]
		permutate(slice, pos+1, fn)
		slice[i], slice[pos] = slice[pos], slice[i]
	}

}

func Permutate(slice []int, fn func(slice []int)) {
	permutate(slice, 0, fn)
}
