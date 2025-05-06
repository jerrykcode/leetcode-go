func solveSudoku(board [][]byte)  {
    if dfs(board, 0, 0) {
       
    }
}
func dfs(board [][]byte, r int, c int) bool {
    if r == 9 {
        return true
    }
    next_r, next_c := r, c + 1
    for ; next_r < 9; next_r++ {
        for ; next_c < 9; next_c++ {
            if board[next_r][next_c] == '.' {
                break
            }
        }
        if next_c < 9 {
            break
        }
        next_c = 0
    }
    if (board[r][c] != '.') {
        return dfs(board, next_r, next_c)
    }
    var i byte
    for i = '1'; i <= '9'; i++ {
        var j, k int
        for j = 0; j < 9; j++ {
            if board[r][j] == i {
                break
            }
        }
        if j < 9 {
            continue
        }
        for j = 0; j < 9; j++ {
            if board[j][c] == i {
                break
            }
        }
        if j < 9 {
            continue
        }
        for j = 0; j < 3; j++ {
            for k = 0; k < 3; k++ {
                if board[(r/3)*3 + j][(c/3)*3 + k] == i {
                    break
                }
            }
            if k < 3 {
                break
            }
        }
        if j == 3 && k == 3 {
            board[r][c] = i
            if dfs(board, next_r, next_c) {
                return true
            }
            board[r][c] = '.'
        }
    }
    return false
}