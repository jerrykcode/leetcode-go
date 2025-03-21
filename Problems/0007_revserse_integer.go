func reverse(x int) int {
    if x < 0 {
        return -1 * reverse(-x)
    }
    var digits [10]int
    var i int = 0
    for x > 0 {
        digits[i] = x % 10
        x /= 10
        i++
    }
    var res int = 0
    for j := 0; j < i; j++ {
        res = res * 10 + digits[j]
        if res > 
    }
    return res
}