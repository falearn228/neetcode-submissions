func spiralOrder(matrix [][]int) []int {
	m,n := len(matrix), len(matrix[0])
	res := make([]int, 0, m*n)

	left, right := 0, len(matrix[0])-1
	top, bot := 0, len(matrix)-1

	for left <= right && top <= bot {
		for i := left; i <= right; i++ {
			res = append(res, matrix[top][i])
		}
		top++

		for i := top; i <= bot; i++ {
			res = append(res, matrix[i][right])
		} 
		right--

		if !(left <= right && top <= bot) {
			break
		}

		for i := right; i >= left; i-- {
			res = append(res, matrix[bot][i])
		}
		bot--

		for i := bot; i >= top; i-- {
			res = append(res, matrix[i][left])
		}
		left++
	}

	return res
}