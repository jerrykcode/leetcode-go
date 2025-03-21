func romanToInt(s string) int {
    table := make(map[byte]int)
    table['I'] = 1
    table['V'] = 5
    table['X'] = 10
    table['L'] = 50
    table['C'] = 100
    table['D'] = 500
    table['M'] = 1000
    res := 0
    for i, _ := range s {
        if i == len(s) - 1 || table[s[i + 1]] <= table[s[i]] {
            res += table[s[i]]
        } else {
            res -= table[s[i]]
        }
    }
    return res
}