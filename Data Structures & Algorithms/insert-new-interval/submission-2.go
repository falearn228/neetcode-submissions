func insert(intervals [][]int, newInterval []int) [][]int {
	res := make([][]int, 0, len(intervals)+1)
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}

	n := len(intervals)
	target := newInterval[0]
	l, r := 0, n-1

	for l <= r {
		mid := (l+r) /2
		if intervals[mid][0] < target {
			l = mid+1
		} else {
			r = mid-1
		}
	}

	intervals = append(intervals[:l], append([][]int{newInterval}, intervals[l:]...)...)

	for _, it := range intervals {
		if len(res) == 0 || res[len(res)-1][1] < it[0] {
			res = append(res, it)
		} else {
			res[len(res)-1][1] = max(res[len(res)-1][1], it[1])
		}
	}

	return res
}