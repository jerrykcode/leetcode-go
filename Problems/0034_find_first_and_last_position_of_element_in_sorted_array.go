func searchRange(nums []int, target int) []int {
    l_smaller, r_larger, l, r, mid, pos := 0, len(nums) - 1, 0, len(nums) - 1, 0, -1
    for l <= r {
        mid = (l + r) / 2;
        if nums[mid] == target {
            pos = mid
            break
        } else if nums[mid] < target {
            l = mid + 1
            l_smaller = mid
        } else {
            r = mid - 1
            r_larger = mid
        }
    }
    res := []int{-1, -1}
    if pos == -1 {
        return res
    }
    l = l_smaller
    r = pos
    for l <= r {
        mid = (l + r) / 2
        if nums[mid] == target {
            res[0] = mid
            r = mid - 1
        } else {
            l = mid + 1
        }
    }
    l = pos
    r = r_larger
    for l <= r {
        mid = (l + r) / 2
        if nums[mid] == target {
            res[1] = mid
            l = mid + 1
        } else {
            r = mid - 1
        }
    }
    return res
}