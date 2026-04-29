func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

    return max(help(nums[1:]), help(nums[:len(nums)-1]))
}

func help(nums []int) int {
	rob1, rob2 := 0, 0

	for _, num := range nums {
		curr := max(rob1+num, rob2)
		rob1 = rob2
		rob2 = curr
	}

	return max(rob1, rob2)
}