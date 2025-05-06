func searchInsert(nums []int, target int) int {
    l, r, mid, res := 0, len(nums) - 1, 0, len(nums)
    for l <= r {
        mid = (l + r) / 2
        if nums[mid] > target {
            res = mid
            r = mid - 1
        } else if nums[mid] < target {
            l = mid + 1
        } else {
            return mid
        }
    }
    return res
}