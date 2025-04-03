func nextPermutation(nums []int)  {
    var i, j int = 0, 0
    for i = len(nums) - 2; i >= 0; i-- {
        if nums[i] < nums[i + 1] {
            break
        }
    }
    if i == -1 {
        for x := 0; x < len(nums) / 2; x++ {
            swap(nums, x, len(nums) - x - 1)
        }
        return
    }
    for j = len(nums) - 1; j > i; j-- {
        if nums[j] > nums[i] {
            break
        }
    }
    swap(nums, i, j)
    sort.Ints(nums[i + 1 : ])
}
func swap(nums []int, i int, j int) {
    nums[i] += nums[j]
    nums[j] = nums[i] - nums[j]
    nums[i] -= nums[j]
}