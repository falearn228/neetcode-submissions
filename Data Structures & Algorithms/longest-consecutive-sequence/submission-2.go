func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	res := 0
	store := make(map[int]struct{})
	for _, num := range nums {
		store[num] = struct{}{}
	}

	for num := range store {
		if _, ok := store[num-1]; !ok {
			leng := 1
			for {
				if _, ok := store[num+leng]; ok {
					leng++
				} else {
					break
				}
			}
			res = max(res, leng)
		}
	}

	return res
}