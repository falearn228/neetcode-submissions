func insert(it [][]int, ni []int) [][]int {
	res := make([][]int, 0, len(it)+1)
	if len(it) == 0 {
		return [][]int{ni}
	}

	n := len(it)
	for i := range n {
		if ni[1] < it[i][0] {
			res = append(res, ni)
			return append(res, it[i:]...)
		} else if ni[0] > it[i][1] {
			res = append(res, it[i])
		} else {
			ni[0] = min(ni[0], it[i][0])
			ni[1] = max(ni[1], it[i][1])
		}
	}

	res = append(res, ni)

	return res
}