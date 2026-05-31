import "slices"

func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}

	res := make([][]int, 0)
	curr := make([]int, 0, len(nums))
	used := make([]bool, len(nums))

	var dfs func()
	dfs = func() {
		if len(curr) == len(nums) {
			res = append(res, slices.Clone(curr))
		}

		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}
			used[i] = true
			curr = append(curr, nums[i])
			dfs()

			used[i] = false
			curr = curr[:len(curr)-1]
		}
	}
	dfs()
	return res
}