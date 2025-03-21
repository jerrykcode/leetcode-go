func myAtoi(s string) int {
    var res, n, i int = 0, len(s), 0
    for i < n && s[i] == ' ' {
        i++
    }
    var positive bool = true
    if i == n {
        return 0
    }
    if s[i] == '-' {
        positive = false
        i++
    } else if s[i] == '+' {
        i++
    }
    for i < n && s[i] >= '0' && s[i] <= '9' {
        res = res * 10 + int(s[i] - '0')
        if res > 2147483647 && positive {
            break
        } 
        if res > 2147483648 {
            break
        }
        i++
    }
    if !positive {
        res *= -1
    }
    if res > 2147483647 {
        res = 2147483647
    } else if res < -2147483648 {
        res = -2147483648
    }
    return res
}