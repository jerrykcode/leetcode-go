func countAndSay(n int) string {
	pre := "1"
	res := "1"
	var range_ch rune
	var range_len int
	for i := 0; i < n-1; i++ {
		res = ""
		range_ch = '~'
		range_len = 0
		for _, ch := range pre {
			if range_ch == '~' || range_ch == ch {
				range_ch = ch
				range_len++
			} else {
				res += strconv.Itoa(range_len)
				res += string(range_ch)
				range_ch = ch
				range_len = 1
			}
		}
		if range_len != 0 {
			res += strconv.Itoa(range_len)
			res += string(range_ch)
		}
		pre = res
	}
	return res
}