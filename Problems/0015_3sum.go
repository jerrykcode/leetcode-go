func threeSum(nums []int) [][]int {
    n := len(nums)
    vals := make([]int, n)
    cnts := make([]int, n)
    sort.Ints(nums)
    var cnt, m int = 1, 0
    for i := 0; i < n; i++ {
        if i + 1 == n || nums[i] != nums[i + 1] {
            vals[m] = nums[i]
            cnts[m] = cnt
            m++
            cnt = 1
        } else {
            cnt++
        }
    }
    var res = [][]int{}
    for i := 0; i < m; i++ {
        var next int = i
        if cnts[i] == 1 {
            next = i + 1
        }
        for j := next; j < m; j++ {
            target := 0 - vals[i] - vals[j]
            left, right := j + 1, m - 1
            if cnts[j] > 2 || (i != j && cnts[j] > 1) {
                left = j
            }
            for left <= right {
                mid := (left + right) / 2
                if vals[mid] == target {
                    res = append(res, []int{vals[i], vals[j], target})
                    break
                } else if vals[mid] > target {
                    right = mid - 1
                } else {
                    left = mid + 1
                }
            }
        }
    }
    return res
}