func minEatingSpeed(piles []int, h int) int {
	var k int

	m := 0
	for _, v := range piles {
		m = max(m, v)
	}

	left, right := 1, m
	for left <= right {
		mid := right-(right-left)/2
		finish := 0
		for i := range len(piles) {
			x := piles[i]
			speed := (x+mid-1)/mid
			finish += speed
		}
		if finish > h {
			left = mid+1
		} else {
			right = mid-1
		}
		k = left
	}

	return k
}