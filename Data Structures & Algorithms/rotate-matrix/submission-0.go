func rotate(m [][]int) {
	if len(m) == 0 {
		return
	}

	r := len(m)

	for i, j := 0, r-1; i < j; i, j = i+1, j-1 {
		m[i], m[j] = m[j], m[i]
	}

	for i := range r {
		for j := i + 1; j < r; j++ {
			m[i][j], m[j][i] = m[j][i], m[i][j]
		}
	}
}
