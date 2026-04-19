func isAnagram(s string, t string) bool {
    n, m := len(s), len(t)
    if n != m {
        return false
    }

    freq := make(map[rune]int)

    for _, v := range s {
        freq[v]++
    }

    for _, v := range t {
        freq[v]--
    }

    for _, val := range freq {
        if val != 0 {
            return false
        }
    }
    return true
}
