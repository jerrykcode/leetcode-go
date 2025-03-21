func intToRoman(num int) string {
    var table = [...][2]string {
        {"I", "V"},
        {"X", "L"},
        {"C", "D"},
        {"M", "input exceed ..."},
    }
    const _1, _5 int = 0, 1
    res := ""
    for i := 0; num > 0; i++ {
        v := num % 10
        if v == 4 {
            res = table[i][_1] + table[i][_5] + res
        } else if v == 9 {
            res = table[i][_1] + table[i + 1][_1] + res
        } else if v < 4 {
            for j := 0; j < v; j++ {
                res = table[i][_1] + res
            }
        } else {
            for j := 5; j < v; j++ {
                res = table[i][_1] + res
            }
            res = table[i][_5] + res
        }
        num /= 10
    }
    return res
}