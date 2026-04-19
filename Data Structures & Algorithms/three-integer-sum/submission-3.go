func threeSum(nums []int) [][]int {
    var res [][]int
    sort.Ints(nums)

    for i := range nums {
        val := nums[i]

        if i > 0 && val == nums[i-1] {
            continue
        }

        l, r := i+1, len(nums)-1
        for l < r {
            sum := val + nums[l] + nums[r]
            if sum > 0 {
                r--
            } else if sum < 0 {
                l++
            } else {
                res = append(res, []int{val, nums[l], nums[r]})
                l,r = l+1, r-1
                for l < r && nums[l] == nums[l-1] {
                    l++
                } 
            }
        }
    }

    return res
}
