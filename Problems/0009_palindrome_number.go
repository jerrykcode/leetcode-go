func isPalindrome(x int) bool {
    if x < 0 {
        return false
    }
    digits := [10]int{}
    var i int
    for i = 0; x > 0; i++ {
        digits[i] = x % 10
        x /= 10
    }
    for j := 0; j < i / 2; j++ {
        if digits[j] != digits[i - 1 - j] {
            return false
        }
    }
    return true
}