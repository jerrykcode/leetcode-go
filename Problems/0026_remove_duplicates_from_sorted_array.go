func removeDuplicates(nums []int) int {
    var delete_start, delete_end int = -1, -1
    for i := 0; i < len(nums); i++ {
        j := binary(nums, i)
        if j == i {
            continue
        }
        if delete_start == -1 {
            delete_start = i + 1
            delete_end = j
        } else {
            move(nums, delete_end + 1, delete_start, i - delete_end)
            delete_start += i - delete_end 
            delete_end = j
        }
        i = j
    }
    if delete_start == -1 {
        return len(nums)
    }
    if delete_end < len(nums) - 1 {
        move(nums, delete_end + 1, delete_start, len(nums) - 1 - delete_end)
        delete_start += len(nums) - 1 - delete_end
    }
    return delete_start
}
func binary(nums []int, i int) int {
    left, right, mid, res := i, len(nums) - 1, 0, i
    for left <= right {
        mid = (left + right) / 2
        if nums[mid] == nums[i] {
            res = mid
            left = mid + 1
        } else if nums[mid] > nums[i] {
            right = mid - 1
        } else {
            // can not reach here
            fmt.Print("ERROR!")
        }
    }
    return res
}
func move(nums []int, src int, des int, sz int) {
    for ; sz > 0; sz-- {
        nums[des] = nums[src]
        src++
        des++
    }
}