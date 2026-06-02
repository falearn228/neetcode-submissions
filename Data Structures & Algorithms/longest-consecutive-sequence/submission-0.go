func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	res := 0
	store := make(map[int]struct{})
	for _, num := range nums {
		store[num] = struct{}{}
	}

	for _, num := range nums {
		streak, curr := 0, num
		var ok bool

		for _, ok = store[curr]; ok; _, ok = store[curr] {
			curr++
			streak++
		}
		
		res = max(res, streak)
	}

	return res
}