func jump(nums []int) int {
	var res, l, r int

	for r < len(nums)-1 {
		far := 0
		for i := l; i <= r; i++ {
			j := nums[i]
			far = max(i+j, far)
		}

		l = r+1
		r = far
		res++
	}
	return res
}
