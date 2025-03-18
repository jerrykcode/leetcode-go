func twoSum(nums []int, target int) []int {
    var val2idxs = make(map[int][]int)
    for idx, val := range nums {
        val2idxs[val] = append(val2idxs[val], idx)
    }
    for idx, val := range nums {
        need := target - val
        var idxs = val2idxs[need]
        for _, idx2 := range idxs {
            if idx2 != idx {
                return []int{idx, idx2}
            }
        }
    }
    return []int{}
}