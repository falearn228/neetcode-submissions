func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	var ans = 0
	last := make(map[byte]int, len(s))
	left := 0

	for right := 0; right < len(s); right++ {
		if prev, ok := last[s[right]]; ok && prev >= left {
			left = prev + 1
		}
		last[s[right]] = right
		ans = max(ans, right-left+1)
	}
	
	return ans
}
