import "slices"

func topKFrequent(nums []int, k int) []int {
	mp := make(map[int]int)
	for _, v := range nums {
		mp[v] += 1
	}

	arr := make([][2]int, 0, k)

	for num, cnt := range mp {
		arr = append(arr, [2]int{num, cnt})
	}

	slices.SortFunc(arr, func(i, j [2]int) int {
		if i[1] >= j[1] {
			return -1
		}
		return 1
	})

	ans := make([]int, 0, k)
	for i := range k {
		ans = append(ans, arr[i][0])
	}
	return ans
}
