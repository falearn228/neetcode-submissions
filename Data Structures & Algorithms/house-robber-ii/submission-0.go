func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

    return max(help(nums[1:]), help(nums[:len(nums)-1]))
}

func help(nums []int) int {
    n := len(nums)
	dp := make([]int, n+1)
	dp[1] = nums[0]

	for i := 1; i < n; i++ {
		dp[i+1] = max(dp[i], nums[i]+dp[i-1])
	}

	return dp[n]
}
