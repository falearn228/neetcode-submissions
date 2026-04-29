func rob(nums []int) int {
	// n := len(nums)
	var rob1, rob2 int

	for i := range nums {
		temp := max(rob1+nums[i], rob2)
		rob1 = rob2
		rob2 = temp
	}

	return max(rob2, rob1)
}
