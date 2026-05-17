func trap(height []int) int {
	var res int
	if len(height) <= 1 {
		return res
	}

	l, r := 0, len(height)-1
	maxl, maxr := height[l], height[r]
	for l < r {
		lbar, rbar := height[l], height[r]

		if lbar <= rbar {
			lbar = height[l]
			if lbar >= maxl {
				maxl = lbar
			} else {
				res += maxl - lbar
			}
			l++
		} else {
			rbar = height[r]
			if rbar >= maxr {
				maxr = rbar
			} else {
				res += maxr - rbar
			}
			r--
		}
	}

	return res
}