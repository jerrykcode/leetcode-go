func longestValidParentheses(s string) int {
    if len(s) == 0 {
        return 0
    }
    dp := make([]int, len(s))
    dp[0] = 0
    res := 0
    for i := 1; i < len(s); i++ {
        if s[i] == '(' {
            dp[i] = 0
        } else {  // )
            if s[i - 1] == '(' {
                dp[i] = 2
                if i - 2 > 0 {
                    dp[i] += dp[i - 2]
                }
            } else {
                if dp[i - 1] > 0 && i - dp[i - 1] - 1 >= 0 && s[i - dp[i - 1] - 1] == '(' {
                    dp[i] = dp[i - 1] + 2
                    if i - dp[i - 1] - 2 > 0 {
                        dp[i] += dp[i - dp[i - 1] - 2]
                    }
                } else {
                    dp[i] = 0
                }
            }
        }
        if dp[i] > res {
            res = dp[i]
        }
    }
    return res
}