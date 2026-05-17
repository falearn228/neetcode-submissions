func checkInclusion(s1 string, s2 string) bool {
	str := [26]int{}
	for _, v := range s1 {
		str[v-'a']++
	}

	i, j := 0, 0
	w := len(s1)
	cpy := [26]int{}
	for ; i < len(s2)-w+1; i++ {
		for ; j-w < i && j < len(s2); j++ {
			cpy[s2[j]-'a']++
		}
		if cpy == str {
			return true
		}
		cpy[s2[i]-'a']--
	}

	return false
}