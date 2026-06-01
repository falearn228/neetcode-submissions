func productExceptSelf(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	res := make([]int, len(nums))

	prefLeft, prefRight := make([]int, len(nums)), make([]int, len(nums))
	prefLeft[0], prefRight[len(nums)-1] = 1, 1

	for i := 1; i < len(nums); i++ {
		prefLeft[i] = nums[i-1] * prefLeft[i-1]
	}

	for i := len(nums)-2; i >= 0; i-- {
		prefRight[i] = nums[i+1] * prefRight[i+1]
	}

	
	for i := range res {
		res[i] = prefLeft[i] * prefRight[i]
	}

	// fmt.Println(res)

	return res
}