func countBits(n int) []int {
	r := make([]int, 0, n+1)
	for i := range n+1 {
		var v = uint(i)
		r = append(r, bits.OnesCount(v))
	}

	return r
}
