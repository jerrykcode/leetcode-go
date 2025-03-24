func removeElement(nums []int, val int) int {
    var del_start, del_end int = -1, -1
    for i := 0; i < len(nums); i++ {
        if nums[i] == val {
            if del_start == -1 {
                del_start = i
                del_end = i
            } else {
                move(nums, del_start, del_end + 1, i - del_end - 1)
                del_start += i - del_end - 1
                del_end = i
            }
        }
    }
    if del_start == -1 {
        return len(nums)
    }
    if del_end < len(nums) - 1 {
        move(nums, del_start, del_end + 1, len(nums) - del_end - 1)
        del_start += len(nums) - del_end - 1
    }
    return del_start
}
func move(nums []int, des int, src int, cnt int) {
    if des < src {
        for i := 0; i < cnt; i++ {
            nums[des + i] = nums[src + i]
        }
    }
}