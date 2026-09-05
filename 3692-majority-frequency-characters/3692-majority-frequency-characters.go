func majorityFrequencyGroup(s string) string {
    freq := [26]int{}
    for _, r := range s {
        freq[r-'a']++
    }
    freqnum := [101]int{}
    for _, f := range freq {
        if f != 0 {
            freqnum[f]++
        }
    }
    m, num := 0, 0
    for i := 100; i >= 0; i-- {
        if m < freqnum[i] {
            m = freqnum[i]
            num = i
        }
    }
    res := ""
    for i, f := range freq {
        if f == num {
            res += string(i + 'a')
        }
    }
    return res
}