func threeSumClosest(nums []int, target int) int {
    sort.Ints(nums)
    n := len(nums)
    var res int
    var has bool = false
    for i, val := range nums {
        for j := i + 1; j < n; j++ {
            var left, right, mid, sum int = j + 1, n - 1, 0, 0
            for left <= right {
                mid = (left + right) / 2
                sum = val + nums[j] + nums[mid]
                if (!has) || abs(sum - target) < abs(res - target) {
                    res = sum
                    has = true
                }
                if sum == target {
                    break
                }
                if sum < target {
                    left = mid + 1
                } else {
                    right = mid - 1
                }
            }
        }
    }
    return res
}
func abs(a int) int {
    if a >= 0 {
        return a
    } else {
        return 0 - a
    }
}
// -4 -1 1 2