func isPalindrome(s string) bool {
	str := strings.ReplaceAll(s, " ","")
	str = strings.ToLower(str)
	n := len(str)
	for i, j := 0, n-1; i < j;  {
		if !isAlphaBet(str[i]) {
			i += 1
			continue
		}
		if !isAlphaBet(str[j]) {
			j -= 1
			continue
		}
		if str[i] != str[j] {
			return false
		}

		i += 1
		j -= 1
	} 

	return true
}

func isAlphaBet(s byte) bool {
	if (s >= 'a' && s <= 'z') || (s >= '0' && s <= '9') {
		return true
	}
	return false
}