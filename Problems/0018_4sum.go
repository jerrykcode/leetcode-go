func fourSum(nums []int, target int) [][]int {
    sort.Ints(nums)
    _nums := []struct{
        data int
        cnt int
    }{}
    appear_idx := make(map[int]int)
    cnt, n := 1, 0
    for i, val := range nums {
        if i == len(nums) - 1 || nums[i + 1] != val {
            _nums = append(_nums, struct{
                data int
                cnt int
            }{val, cnt})
            appear_idx[val] = n
            n++
            cnt = 1
        } else {
            cnt++
        }
    }
    res := [][]int{}
    for i := 0; i < n; i++ {
        j := i
        if _nums[i].cnt == 1 {
            j = i + 1
        }
        for ; j < n; j++ {
            k := j
            if _nums[j].cnt == 1 || (j == i && _nums[j].cnt <= 2) {
                k = j + 1
            }
            for ; k < n; k++ {
                need := target - _nums[i].data - _nums[j].data - _nums[k].data
                idx, exists := appear_idx[need]
                if exists {
                    ok := false
                    if idx > k {
                        ok = true
                    } else if idx == k {
                        if k > j && _nums[k].cnt > 1 {
                            ok = true
                        }
                        if k == j && j > i && _nums[k].cnt > 2 {
                            ok = true
                        }
                        if k == j && j == i && _nums[k].cnt > 3 {
                            ok = true
                        }
                    }
                    if ok {
                        res = append(res, []int{_nums[i].data, _nums[j].data, _nums[k].data, need})
                    }
                }
            }
        }
    }
    return res
}