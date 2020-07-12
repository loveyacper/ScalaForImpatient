package ch5string

// 1312

func MinInsertions(s string) int {
    if len(s) <= 1 {
        return 0
    }

    dp := make([][]int, 0, len(s)) // dp[i][j] -> s[i..j]'s min insertions, we need dp[0..len(s)-1]
    for i := 0; i < len(s); i++ {
        dp = append(dp, make([]int, len(s)))
    }

    for step := 1; step < len(s); step++ {
        for i := 0; i + step < len(s); i++ {
            j := i + step
            if s[i] == s[j] {
                dp[i][j] = dp[i+1][j-1]
            } else {
                a := dp[i][j-1]+1
                b := dp[i+1][j]+1
                if a < b {
                    dp[i][j] = a
                } else {
                    dp[i][j] = b
                }
            }
        }
    }

    return dp[0][len(s)-1]
}
