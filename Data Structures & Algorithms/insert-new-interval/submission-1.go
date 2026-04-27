func insert(intervals [][]int, newInterval []int) [][]int {
	res := make([][]int, 0, len(intervals)+1)
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}

	n := len(intervals)
	i := 0
	for i < n && intervals[i][1] < newInterval[0] {
		res = append(res, intervals[i])
		i++
	}

	for i < n && intervals[i][0] <= newInterval[1]{
		newInterval[0] = min(newInterval[0], intervals[i][0])
		newInterval[1] = max(newInterval[1], intervals[i][1])
		i++
	}
	res = append(res, newInterval)
	for ; i < n; i++ {
		res = append(res, intervals[i])
	}

	return res
}
