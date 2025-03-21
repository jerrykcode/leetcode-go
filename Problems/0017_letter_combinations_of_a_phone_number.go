var d2a = [8][4] byte { // digit to alphabet [2, 9]
    {'a', 'b', 'c', ' '},
    {'d', 'e', 'f', ' '},
    {'g', 'h', 'i', ' '},
    {'j', 'k', 'l', ' '},
    {'m', 'n', 'o', ' '},
    {'p', 'q', 'r', 's'},
    {'t', 'u', 'v', ' '},
    {'w', 'x', 'y', 'z'},
}
var final_res []string
func letterCombinations(digits string) []string {
    if len(digits) == 0 {
        return []string{}
    }
    res := make([]byte, len(digits))
    final_res = []string{}
    dfs(digits, res, 0)
    return final_res
}

func dfs(digits string, res []byte, idx int) {
    n := len(digits)
    if idx == n {
        final_res = append(final_res, string(res))
        return
    }
    for _, letter := range d2a[digits[idx] - '0' - 2] {
        if letter == ' ' {
            break
        }
        res[idx] = letter
        dfs(digits, res, idx + 1)   
    }
}