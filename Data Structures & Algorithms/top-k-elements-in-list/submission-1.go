func topKFrequent(nums []int, k int) []int {
	count := make(map[int]int)
	bucket := make([][]int, len(nums)+1)

	for _, v := range nums {
		count[v]++
	}

	for num, count := range count {
		bucket[count] = append(bucket[count], num)
	}

	var res []int
	for i := len(bucket)-1; i > 0; i-- {
		for _, num := range bucket[i] {
			res = append(res, num)
			if len(res) == k {
				return res
			}
		}
	}
	return []int{}
}
