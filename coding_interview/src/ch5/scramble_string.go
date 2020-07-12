package ch5string

// 87

// dp[i][j][len] : s1[i:i+len] and s2[j:j+len] is scramble if val = 1, else val = 2
const (
    Unknown int = iota
    IsCramble
    NotCramble
)

func IsScramble(s1 string, s2 string) bool {
    if len(s1) != len(s2) {
        return false
    }

    // init dp
    strlen := len(s1)
    dp := make([][][]int, strlen)
    for i := 0; i < len(dp); i++ {
        dp[i] = make([][]int, strlen)
        for j := 0; j < len(dp[i]); j++ {
            dp[i][j] = make([]int, strlen+1) // length is 1-based
            for k := 0; k < len(dp[i][j]); k++ {
                dp[i][j][k] = Unknown
            }
        }
    }

    return isScramble(s1, 0, strlen, s2, 0, strlen, dp)
}

func isScramble(s1 string, start1, end1 int, s2 string, start2, end2 int, dp [][][]int) bool {
    if end1 - start1 != end2 - start2 {
        return false
    }
    strlen := end1 - start1
    if strlen == 0 {
        return true
    } else if strlen == 1 {
        return s1[start1] == s2[start2]
    }

    if dp[start1][start2][strlen] != Unknown {
        return dp[start1][start2][strlen] == IsCramble
    }

    // maybe start1 != start2

    for split := start1+1; split < end1; split++ {
        // 以split为分界点，看s1的左半部分能否和s2的左半部分，且s1的右半部分和s2的右半部分是scramble
        // 或者看s1的左半部分能否和s2的右半部分，且s1的右半部分和s2的左半部分是scramble
        // ps: split点划到右半部分
        distance := split - start1
        // 先看s1的左半部分能否和s2的左半部分，且s1的右半部分和s2的右半部分是scramble
        if isScramble(s1, start1, split, s2, start2, start2+distance, dp) &&
           isScramble(s1, split, end1, s2, start2+distance, end2, dp) {
           dp[start1][start2][strlen] = IsCramble
           return true
        }

        // 再试试看s1的左半部分能否和s2的右半部分，且s1的右半部分和s2的左半部分是scramble
        if isScramble(s1, start1, split, s2, end2-distance, end2, dp) &&
           isScramble(s1, split, end1, s2, start2, end2-distance, dp) {
           dp[start1][start2][strlen] = IsCramble
           return true
        }
    }

    dp[start1][start2][strlen] = NotCramble
    return false
}

//  不用dp，节约空间, 但是可能很慢
func IsScramble2(s1 string, s2 string) bool {
    if len(s1) != len(s2) {
        return false
    }

    strlen := len(s1)

    if s1 == s2 {
        return true
    }

    if strlen == 0 {
        return true
    } else if strlen == 1 {
        return false
    }

    stat := [256]int{}
    for i := 0; i < strlen; i++ {
        stat[s1[i]]++
        stat[s2[i]]--
    }

    for i := 0; i < len(stat); i++ {
        if stat[i] != 0 {
            //  字符不完全一致
            return false
        }
    }

    for split := 1; split < strlen; split++ {
        // 以split为分界点，看s1的左半部分能否和s2的左半部分，且s1的右半部分和s2的右半部分是scramble
        // 或者看s1的左半部分能否和s2的右半部分，且s1的右半部分和s2的左半部分是scramble
        // ps: split点划到右半部分

        // 先看s1的左半部分能否和s2的左半部分，且s1的右半部分和s2的右半部分是scramble
        if IsScramble2(s1[:split], s2[:split]) &&
           IsScramble2(s1[split:], s2[split:]) {
           return true
        }

        // 再试试看s1的左半部分能否和s2的右半部分，且s1的右半部分和s2的左半部分是scramble
        if IsScramble2(s1[:split], s2[strlen-split:]) &&
           IsScramble2(s1[split:], s2[:strlen-split]) {
           return true
        }
    }

    return false
}
/*
Given a string s1, we may represent it as a binary tree by partitioning it to two non-empty substrings recursively.

Below is one possible representation of s1 = "great":

    great
   /    \
  gr    eat
 / \    /  \
g   r  e   at
           / \
          a   t
To scramble the string, we may choose any non-leaf node and swap its two children.

For example, if we choose the node "gr" and swap its two children, it produces a scrambled string "rgeat".

    rgeat
   /    \
  rg    eat
 / \    /  \
r   g  e   at
           / \
          a   t
We say that "rgeat" is a scrambled string of "great".

Similarly, if we continue to swap the children of nodes "eat" and "at", it produces a scrambled string "rgtae".

    rgtae
   /    \
  rg    tae
 / \    /  \
r   g  ta  e
       / \
      t   a
We say that "rgtae" is a scrambled string of "great".

Given two strings s1 and s2 of the same length, determine if s2 is a scrambled string of s1.

Example 1:

Input: s1 = "great", s2 = "rgeat"
Output: true
Example 2:

Input: s1 = "abcde", s2 = "caebd"
Output: false
*/
