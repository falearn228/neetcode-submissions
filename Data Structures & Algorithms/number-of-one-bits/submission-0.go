func hammingWeight(n int) int {
	// 0001001
	// 1110110
	// 
	var r int
	for n > 0 {
		if n&1 != 0 {
			r++
		}
		n = n >> 1
	}

	return r
}