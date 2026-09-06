func lengthOfLongestSubstring(s string) int {
    seen := make(map[byte]int)

    left := 0
    maxLength := 0

    for right := 0; right < len(s); right++ {

        index, exists := seen[s[right]]

        if exists && index >= left {
            left = index + 1
        }

        seen[s[right]] = right

        currentLength := right - left + 1

        if currentLength > maxLength {
            maxLength = currentLength
        }
    }

    return maxLength
}