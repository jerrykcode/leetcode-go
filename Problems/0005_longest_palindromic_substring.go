func longestPalindrome(s string) string {
    var n int = len(s)
    var dp = make([][]bool, n)
    for i := 0; i < n; i++ {
        dp[i] = make([]bool, n)
        dp[i][i] = true
    }
    var res, res_i, res_j int = 1, 0, 0
    for l := 2; l <= n; l++ {
        for i, j := 0, l - 1; j < n; {
            if s[i] == s[j] && (i + 1 == j || dp[i + 1][j - 1]) {
                if l > res {
                    res = l
                    res_i = i
                    res_j = j
                }
                dp[i][j] = true
            } else {
                dp[i][j] = false
            }
            i++
            j++
        } 
    }
    return s[res_i : res_j + 1]
}