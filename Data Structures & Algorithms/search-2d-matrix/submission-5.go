func searchMatrix(nums [][]int, target int) bool {
	if len(nums) == 0 {
		return false
	}

	left, right := 0, len(nums)-1

	for left < right {
		mid := right-(right-left)/2

		if nums[mid][0] < target {
			left = mid+1
		} else if nums[mid][0] > target {
			right = mid-1
		} else {
			return true
		}
	}

	for l := left-1; l <= right+1; l++ {
		if l < 0 || l >= len(nums) {
			continue
		}
		nleft, nright := 0, len(nums[0])-1
		for nleft <= nright {
			mid := nright-(nright-nleft)/2

			if nums[l][mid] < target {
				nleft = mid+1 
			} else if nums[l][mid] > target {
				nright = mid-1
			} else {
				return true
			}
		}
	}

	return false
}
