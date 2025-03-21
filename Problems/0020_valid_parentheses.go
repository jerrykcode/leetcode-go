func isValid(s string) bool {
    stack := make([]byte, len(s))
    var cnt int = 0
    for i := 0; i < len(s); i++ {
        if s[i] == '(' || s[i] == '[' || s[i] == '{' {
            stack[cnt] = s[i]
            cnt++
        } else if s[i] == ')' {
            if cnt >= 1 && stack[cnt - 1] == '(' {
                cnt--
            } else {
                return false
            }
        } else if s[i] == ']' {
            if cnt >= 1 && stack[cnt - 1] == '[' {
                cnt--
            } else {
                return false
            }
        } else if s[i] == '}' {
            if cnt >= 1 && stack[cnt - 1] == '{' {
                cnt--
            } else {
                return false
            }
        }
    }
    return cnt == 0
}