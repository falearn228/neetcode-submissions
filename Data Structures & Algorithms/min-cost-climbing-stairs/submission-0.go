func minCostClimbingStairs(cost []int) int {
	dp := make([]int, len(cost)+1)
	
	for i := 1; i < len(dp)-1; i++ {
		dp[i+1] = min(dp[i]+cost[i], dp[i-1]+cost[i-1])
	}

	n := len(cost)
	return dp[n]
}