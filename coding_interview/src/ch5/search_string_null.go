package ch5string

import (
    "fmt"
    "strings"
)


// If multi str found, return the most left one
func SearchStringsWithNull(strs []*string, key string) int {
    pos := -1

    left := 0
    right := len(strs)

    for {
        for left < len(strs) && strs[left] == nil {
            left++
            if left >= right {
                return pos
            }
        }
        for right >= 1 && strs[right-1] == nil {
            right--
            if left >= right {
                return pos
            }
        }

        if left >= right {
            panic(fmt.Sprintf("BUG: left %d >= right %d", left, right))
        }

        mid := (left+right)/2
        if strs[mid] != nil {
            cmp := strings.Compare(*(strs[mid]), key)
            if cmp == 0 {
                pos = mid
                right = mid // keep search left part
            } else if cmp > 0 {
                right = mid
            } else {
                left = mid+1
            }
        } else {
            for i := mid-1; i >= left; i-- {
                if strs[i] != nil {
                    cmp := strings.Compare(*(strs[i]), key)
                    if cmp == 0 {
                        pos = i
                        right = i // keep search left part
                    } else if cmp > 0 {
                        right = i
                    } else {
                        left = mid+1
                    }

                    break
                }
            }
        }
    }

    return pos
}

/*
func stringPtr(s string) *string {
    return &s
}

func main() {
    strs := append([]*string(nil), stringPtr("hello"))
    strs = append(strs, stringPtr("hello"))
    strs = append(strs, nil)
    strs = append(strs, stringPtr("world"))
    strs = append(strs, nil)
    strs = append(strs, stringPtr("world"))

    fmt.Println(SearchStringsWithNull(strs, "world"))
}
*/
