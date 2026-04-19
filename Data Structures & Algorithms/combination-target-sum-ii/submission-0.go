func combinationSum2(nums []int, target int) [][]int {
	res := make([][]int, 0)
	sort.Ints(nums)

	var dfs func(ind, sum int, curr []int)
	dfs = func(ind, sum int, curr []int) {
		if sum > target {
			return
		}
		if sum == target {
			tmp := make([]int, len(curr))
			copy(tmp, curr)
			res = append(res, tmp)
			return
		}

		for i := ind; i < len(nums); i++ {
			if i > ind && nums[i] == nums[i-1] {
				continue
			}
			curr = append(curr, nums[i])
			dfs(i+1, sum+nums[i], curr)
			curr = curr[:len(curr)-1]
		}
	}

	dfs(0, 0, []int{})
	return res
}