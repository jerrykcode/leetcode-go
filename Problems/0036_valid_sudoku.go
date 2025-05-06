func isValidSudoku(board [][]byte) bool {
    used := make([]bool, 10)
    for i := 0; i < 9; i++ {
        initialize(used)
        for j := 0; j < 9; j++ {
            if check(used, board[i][j]) {
                return false
            }
        }
    }
    for i := 0; i < 9; i++ {
        initialize(used)
        for j := 0; j < 9; j++ {
            if check(used, board[j][i]) {
                return false
            }
        }
    }
    for i := 0; i <= 6; i += 3 {
        for j := 0; j <= 6; j += 3 {
            initialize(used)
            for a := 0; a < 3; a++ {
                for b := 0; b < 3; b++ {
                    if check(used, board[i + a][j + b]) {
                        return false
                    }
                }
            }
        }
    }
    return true
}
func initialize(used []bool)  {
    for i, _ := range used {
        used[i] = false
    }
}
func check(used []bool, v byte) bool {
    if v == '.' {
        return false
    }
    v -= '0'
    res := used[v]
    used[v] = true
    return res
}