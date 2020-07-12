package ch5string

// # 316
// https://leetcode.com/problems/remove-duplicate-letters/
// 删除多余的字符，使得字符串的每个字符只出现一次，且字典序最小
//"cbacdcbc" -> Output: "acdb"

func RemoveDuplicateLetters(s string) string {
    // 统计字符频率
    frequency := [256]int{}
    for _, ch := range s {
        frequency[int(ch)]++
    }

    // used数组判断字符是否已经在result
    used := [256]bool{}
    result := make([]rune, 0)
    for _, ch := range s {
        frequency[int(ch)]--
        if used[int(ch)] {
            continue
        }

        for len(result) > 0 {
            back := result[len(result)-1]

            // back != ch
            if back < ch {
                break
            } else {
                if frequency[int(back)] > 0 {
                    // 已经选择的这个back字符比ch大，且后面还有back字符
                    // 先弹出back，让ch进去，字典序更小
                    result = result[:len(result)-1]
                    used[int(back)] = false
                } else {
                    break
                }
            }
        }

        result = append(result, ch)
        used[int(ch)] = true
    }

    return string(result)
}

