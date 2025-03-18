func lengthOfLongestSubstring(s string) int {
    var last_appear_idx = [256]int{}
    for i := 0; i < 256; i++ {
        last_appear_idx[i] = -1
    }
    var n, i, j, res int = len(s), 0, 0, 0
    for j < n {
        ch := s[j]
        last := last_appear_idx[ch]
        if last >= i {
            // [i, j - 1]
            if j - i > res {
                res = j - i
            }
            i = last + 1
        }
        last_appear_idx[ch] = j
        j++
    }
    if n - i > res {
        res = n - i
    }
    return res
}