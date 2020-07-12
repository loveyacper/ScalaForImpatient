package ch5string

// 003

func NonRepeatSubstring(s string) int {
    lastVisit := [256]int{}
    for i := 0; i < len(lastVisit); i++ {
        lastVisit[i] = -1
    }

    maxLen := 0
    start := 0
    for i := 0; i < len(s); i++ {
        if lastVisit[int(s[i])] != -1 {
            // s[i] is repeated, update start pos
            start = lastVisit[int(s[i])] + 1
        }

        lastVisit[int(s[i])] = i

        if maxLen < i - start + 1 {
            maxLen = i - start + 1
        }
    }

    return maxLen
}

