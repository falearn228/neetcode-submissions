func characterReplacement(s string, k int) int {
	if len(s) == 1 {
		return 1
	}

	ans := 0
	freq := [26]int{}
	maxFreq := 0
	left := 0

	for right := 0; right < len(s); right++ {
		idx := s[right] - 'A'
		freq[idx]++
		maxFreq = max(maxFreq, freq[idx])

		windowLen := right - left + 1
		for windowLen > maxFreq+k {
			freq[s[left]-'A']--
			left++
			windowLen = right-left+1
		}
		ans = max(ans, windowLen)
	}
	
	return ans
}
