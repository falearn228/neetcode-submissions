func maxProfit(prices []int) int {
	minBuy, maxSell := math.MaxInt, math.MinInt
	ans := 0
	for i := range prices {
		maxSell = prices[i]
		minBuy = min(minBuy, prices[i])
		ans = max(ans, maxSell-minBuy)
	}

	return ans
}
