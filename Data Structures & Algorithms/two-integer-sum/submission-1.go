func twoSum(nums []int, target int) []int {
    freq := make(map[int]int)

    for i := range nums {
        if index, ok := freq[target-nums[i]]; ok {
            return []int{index, i}
        }
        freq[nums[i]] = i
    }

    return []int{}
}
