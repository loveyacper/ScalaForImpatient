package ch5string

// 151
import (
    "strings"
)

func ReverseWords(s string) string {
    if len(s) == 0 {
        return s
    }

    var builder strings.Builder

    end := -1
    first := true
    for j := len(s)-1; j >= 0; j-- {
        if s[j] == ' ' {
            if end != -1 {
                if first {
                    builder.WriteString(s[j+1:end])
                    first = false
                } else {
                    builder.WriteString(" " + s[j+1:end])
                }
                end = -1
            }
        } else {
            if end == -1 {
                end = j+1
            }
        }
    }

    if end != -1 {
        if first {
            builder.WriteString(s[0:end])
        } else {
            builder.WriteString(" " + s[0:end])
        }
    }

    return builder.String()
}

