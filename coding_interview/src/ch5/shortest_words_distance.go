package ch5string

// 245

// word1和word2可以相同，认为是两个不同下标实体
// ["p", "m", "pe", "c", "m"], word1 = word2 = "m", return 3
func ShortestWordDistance(arr []string, word1, word2 string) int {
    same := (word1 == word2)

    pos1, pos2 := -len(arr), -len(arr)
    minDist := len(arr)

    for i, s := range arr {
        if s == word1 {
            if same && pos2 < pos1 {
                pos2 = i
            } else {
                pos1 = i
            }
            if myAbs(pos1 - pos2) < minDist {
                minDist = myAbs(pos1 - pos2)
            }
        } else if s == word2 {
            if same && pos1 < pos2 {
                pos1 = i
            } else {
                pos2 = i
            }
            if myAbs(pos1 - pos2) < minDist {
                minDist = myAbs(pos1 - pos2)
            }
        }
    }

    if minDist >= len(arr) {
        return -1
    }

	return minDist
}

func myAbs(n int) int {
    if n < 0 {
        return -n
    }
    return n
}

