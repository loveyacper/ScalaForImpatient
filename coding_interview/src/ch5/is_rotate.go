package ch5string

// 796
import "strings"

func IsRotate(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}

    s := str2 + str2
    return strings.Contains(s, str1)
}
