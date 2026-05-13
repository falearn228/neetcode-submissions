func groupAnagrams(strs []string) [][]string {
	res := make(map[[26]int][]string)
	var ans [][]string

	for _, s := range strs {
		var count [26]int
		for _, c := range s {
			count[c-'a']++
		}
		res[count] = append(res[count], s)
	}

	for _, group := range res {
		ans = append(ans, group)
	}

	return ans
}
