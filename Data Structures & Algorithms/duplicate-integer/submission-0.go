func hasDuplicate(nums []int) bool {
    freq := make(map[int]struct{})

    for i := range nums {
        if _, ok := freq[nums[i]]; ok {
            return true
        }
        freq[nums[i]] = struct{}{}
    }
    return false
}
