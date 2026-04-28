func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
    
    var prev2 = 2
    var prev1 = 1
    
    for i := 3; i <= n; i++ {
        var current = prev1 + prev2
        prev1 = prev2
        prev2 = current
    }
    
    return prev2
}
