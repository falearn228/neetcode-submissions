func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 1 {
		return []int{nums[0]}
	}

	dq := []int{}
	res := []int{}

	for i, curr := range nums {

		for len(dq) > 0 && dq[0] <= i-k {
			dq = dq[1:]
		}
		
		for len(dq) > 0 && curr > nums[dq[len(dq)-1]] {
			dq = dq[:len(dq)-1]
		}
		dq = append(dq, i)

		if i >= k-1 { 
			res = append(res, nums[dq[0]])
		}
	}
	return res
}
