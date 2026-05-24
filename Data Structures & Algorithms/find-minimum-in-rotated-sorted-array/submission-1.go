
func findMin(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	l, r := 0, len(nums)-1
	for l < r {
		mid := l + (r-l)/2

		if nums[mid] < nums[r] {
			r = mid
		} else {
			l = mid+1
		}
	}
	return nums[l]
}