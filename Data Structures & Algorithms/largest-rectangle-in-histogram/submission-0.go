func largestRectangleArea(heights []int) int {
	if len(heights) == 1 {
		return heights[0]
	}

	n := len(heights)
	sck := make([]int, 0, n)
	var maxArea int

	for i := range n + 1 {
		for len(sck) > 0 && (i == n || heights[i] < heights[sck[len(sck)-1]]) {
			height := heights[sck[len(sck)-1]]
			sck = sck[:len(sck)-1]

			width := i
			if len(sck) > 0 {
				width = i - sck[len(sck)-1] - 1
			}

			area := height * width
			maxArea = max(maxArea, area)

		}
		if i < n {
			sck = append(sck, i)
		}
	}

	return maxArea
}