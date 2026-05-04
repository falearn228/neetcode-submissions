func hammingWeight(n int) int {
	// 0001001
	// 1110110
	// 
	var r int
	for n > 0 {
        n = n & (n-1) // remove rightest 1
        r++
	}

	return r
}