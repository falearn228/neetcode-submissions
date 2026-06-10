func search(nums []int, target int) int {
	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		}
		return -1
	}
	l, r := 0, len(nums)-1

	var mmin func(int, int) int
	mmin = func(l, r int) int {
		for l < r {
			mid := l + (r-l)/2

			if nums[mid] < nums[r] {
				r = mid
			} else {
				l = mid + 1
			}
		}
		return l
	}

	split := mmin(l, r)

	lnums, rnums := nums[:split], nums[split:]

	if len(lnums) > 0 && lnums[0] <= target && lnums[len(lnums)-1] >= target {
		r = split-1 
	} else if len(rnums) > 0 && rnums[0] <= target && rnums[len(rnums)-1] >= target {
		l = split
	} else {
		return -1
	}

	for l < r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			return mid
		}

		if nums[mid] > target {
			r = mid
		} else {
			l = mid + 1
		}
	}
	if nums[l] == target {
		return l
	}

	return -1
}