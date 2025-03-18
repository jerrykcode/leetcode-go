type Output struct {
    Row int
    Col int
    Data byte
} 
func convert(s string, numRows int) string {
    if numRows == 1 {
        return s
    }
    outputs := []Output{}
    row, col := 0, 0
    full_col := true
    for _, data := range s {
        outputs = append(outputs, Output{row, col, byte(data)})
        if row == 0 {
            full_col = true
        }
        if full_col {
            if row + 1 < numRows {
                row++
            } else {
                full_col = false
            }
        } 
        if !full_col {
            row--
            col++
        }
    }
    sort.Slice(outputs, func(i int, j int) bool {
        if outputs[i].Row != outputs[j].Row {
            return outputs[i].Row < outputs[j].Row
        } else {
            return outputs[i].Col < outputs[j].Col
        }
    })
    res := ""
    for _, output := range outputs {
        res += string(output.Data)
    }
    return res
}