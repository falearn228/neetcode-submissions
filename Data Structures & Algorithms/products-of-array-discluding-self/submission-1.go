func productExceptSelf(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	res := make([]int, len(nums))

	prefix := 1
	postfix := 1

	for i := 0; i < len(nums); i++ {
		res[i] = prefix
		prefix *= nums[i]
	}

	for i := len(nums)-1; i >= 0; i-- {
		res[i] *= postfix
		postfix *= nums[i]
	}

	return res
}