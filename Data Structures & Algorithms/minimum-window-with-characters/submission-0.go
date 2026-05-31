func minWindow(s string, t string) string {
    if len(s) < len(t) || t == "" {
		return ""
	}

	countT := make(map[byte]int)
	for i := range t {
		countT[t[i]]++
	}

	have, need := 0, len(countT)
	res := []int{-1, -1}
	reslen := math.MaxInt
	l := 0
	window := make(map[byte]int)

	for r := 0; r < len(s); r++ {
		c := s[r]
		window[c]++

		if _, ok := countT[c]; ok && window[c] == countT[c] {
			have++
		}

		for have == need {
			if r-l+1 < reslen {
				res = []int{l, r}
				reslen = r-l+1
			}

			window[s[l]]--
			if _, ok := countT[s[l]]; ok && window[s[l]] < countT[s[l]] {
				have--
			}
			l++
		}
	}

	if res[0] == -1 {
		return ""
	}
	return s[res[0]:res[1]+1]
}
