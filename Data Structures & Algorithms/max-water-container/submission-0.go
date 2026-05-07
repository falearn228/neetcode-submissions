func maxArea(heights []int) int {
	if len(heights) == 0 {
		return 0
	}

	ans := 0
	
	l, r := 0, len(heights)-1
	for l < r {
		ans = max(ans, (r-l)*min(heights[l], heights[r]))
		if heights[l] < heights[r] {
			l++
		} else {
			r--
		}
	}

	return ans
}

