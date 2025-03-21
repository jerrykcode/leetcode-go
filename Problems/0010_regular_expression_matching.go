func isMatch(s string, p string) bool {
    dp := make([][]bool, len(s) + 1)
    for i := 0; i <= len(s); i++ {
        dp[i] = make([]bool, len(p) + 1)
    }
    dp[0][0] = true
    for i := 1; i <= len(p); i++ {
        dp[0][i] = false
        if i >= 2 && p[i - 1] == '*' {
            dp[0][i] = dp[0][i - 2]
        }
    }
    for i := 1; i <= len(s); i++ {
        dp[i][0] = false
        for j := 1; j <= len(p); j++ {
            dp[i][j] = false
            if s[i - 1] == p[j - 1] || p[j - 1] == '.' {
                dp[i][j] = dp[i - 1][j - 1]
            } else if j >= 2 && p[j - 1] == '*' {
                for k := i; k >= 0; k-- {
                    if dp[k][j - 2] {
                        dp[i][j] = true
                        break
                    }
                    if k > 0 && s[k - 1] != p[j - 2] && p[j - 2] != '.' {
                        break
                    }
                }
            }
        }
    }
    return dp[len(s)][len(p)]
}