func generateParenthesis(n int) []string {
    res := []string{}
    dfs("", 0, n, &res)
    return res
}
func dfs(p string, left int, tot int, res *[]string) {
    if len(p) == tot * 2 {
        *res = append(*res, p)
        return
    }
    if tot * 2 - len(p) > left  {
        dfs(p + "(", left + 1, tot, res)
    }
    if left > 0 {
        dfs(p + ")", left - 1, tot, res)
    }
}