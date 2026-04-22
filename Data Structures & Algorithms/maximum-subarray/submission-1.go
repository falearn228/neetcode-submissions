func maxSubArray(nums []int) int {
    var res, maxx = -1_000_000_000, -1_000_000_000
    for i := range nums {
        res = max(nums[i], res+nums[i])
        maxx = max(maxx, res)
    }

    return maxx
}
