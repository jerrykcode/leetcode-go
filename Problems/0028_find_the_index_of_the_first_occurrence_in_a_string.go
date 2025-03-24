func strStr(haystack string, needle string) int {
    s := needle + "#" + haystack
    pi := make([]int, len(s))
    for i := 1; i < len(s); i++ {
        j := pi[i - 1]
        for j > 0 && s[j] != s[i] {
            j = pi[j - 1]
        }
        if s[i] == s[j] {
            j++
        }
        pi[i] = j
        if pi[i] == len(needle) && i >= len(needle) {
            return i - 2 * len(needle)
        }
    }
    return -1
}