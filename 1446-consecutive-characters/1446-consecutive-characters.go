func maxPower(s string) int {
    left := 0
    right:= left + 1
    maxPower := 1
    curPower := 0
    for right < len(s) {
        if s[left] == s[right] {
            curPower = right -left + 1
            right += 1
        } else {
            curPower = 0
            left += 1
            right = left + 1
        }

        if curPower > maxPower {
            maxPower = curPower
        }
    }
    return maxPower
}