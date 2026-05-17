func checkInclusion(s1 string, s2 string) bool {
	w := len(s1)
	if w > len(s2) {
		return false
	}

	var need [26]int
	var window [26]int

	for i := 0; i < w; i++ {
		need[s1[i]-'a']++
		window[s2[i]-'a']++
	}

	if need == window {
		return true
	}

	for r := w; r < len(s2); r++ {
		window[s2[r]-'a']++       // добавили новый правый символ
		window[s2[r-w]-'a']--     // удалили левый старый символ

		if window == need {
			return true
		}
	}

	return false
}