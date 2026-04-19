func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	var ans = 1
	set := make(map[byte]bool, len(s))
	left, right := 0, 1
	for right < len(s) {
		set[s[left]] = true
		if ok := set[s[right]]; ok {
			for left < right && set[s[right]] {
				set[s[left]] = false
				left++
			}
		}
		ans = max(ans, right-left+1)
		set[s[right]] = true
		right++
	}

	return ans
}
