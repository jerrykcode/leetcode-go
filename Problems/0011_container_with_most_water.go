func maxArea(height []int) int {
    var l, r int = 0, len(height) - 1
    var res, tmp int = 0, 0
    for l <= r {
        tmp = (r - l) * min(height[l], height[r])
        if tmp > res {
            res = tmp
        }
        if height[l] < height[r] {
            l++
        } else {
            r--
        }
    }
    return res
}