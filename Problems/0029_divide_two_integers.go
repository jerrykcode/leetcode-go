func divide(dividend int, divisor int) int {
    if dividend == 0 {
        return 0
    }
    pos := true
    if dividend < 0 {
        dividend = 0 - dividend
        pos = !pos
    }
    if divisor < 0 {
        divisor = 0 - divisor
        pos = !pos
    }
    res := 0
    for x := dividend; x >= divisor; {
        y, i := divisor, 1
        for x >= y + y {
            y += y
            i += i
        }
        res += i
        x -= y
    }
    if !pos {
        res = 0 - res
    }
    if res < -2147483648 {
        res = -2147483648
    }
    if res > 2147483647 {
        res = 2147483647
    }
    return res
}