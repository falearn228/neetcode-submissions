func subsets(nums []int) [][]int {
	res := make([][]int, 0)

	var dfs func(ind int, curr []int)
	dfs = func(ind int, curr []int) {
		tmp := make([]int, len(curr))
		copy(tmp, curr)
		res = append(res, tmp)
			

		for i := ind; i < len(nums); i++ {
			curr = append(curr, nums[i])
			dfs(i+1, curr)
			curr = curr[:len(curr)-1]
		}
	}

	dfs(0, []int{})
	return res
}

