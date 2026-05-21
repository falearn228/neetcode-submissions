func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	if len(temperatures) <= 1 {
		return res
	}

	stack := []int{}
	for i, v := range temperatures {
		for len(stack) > 0 && v > temperatures[stack[len(stack)-1]] {
			ind := stack[len(stack)-1]
			res[ind] = i - ind
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}

	return res
}