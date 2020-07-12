package ch5string

import (
    "errors"
    "strconv"
    "strings"
    "unicode/utf8"
)


// #443
// "aaabba" -> a_3_b_2_a_1

func StatString(str string) string {
    var c rune
    start := -1
    var builder strings.Builder

    for i, v := range str {
        if v != c || start == -1 {
            end := i
            if start != -1 {
                builder.WriteRune(c)
                nRunes := (end - start) / utf8.RuneLen(v)
                builder.WriteString("_" + strconv.Itoa(nRunes) + "_")
            }

            c = v
            start = i
        }
    }

    if start >= 0 && start < len(str) {
        builder.WriteRune(c)
        nRunes := (len(str) - start) / utf8.RuneLen(c)
        builder.WriteString("_" + strconv.Itoa(nRunes))
    }

    return builder.String()
}

type item struct {
    upperIndex int
    c rune
}

type IndexArray []item

// a_3_b_2 -> Get(3) = b
func NewIndexArray(stat string) IndexArray {
    arr := make([]item, 0)
    off := 0

    var inNum = false
    var c rune
    var builder strings.Builder

    for _, v := range stat {
        if v == '_' {
            if inNum {
                idx, _ := strconv.Atoi(builder.String())
                off += idx
                builder.Reset()

                arr = append(arr, item{upperIndex: off, c : c})
                inNum = false
            }
        } else if v >= '0' && v <= '9' {
            inNum = true
            builder.WriteRune(v)
        } else {
            inNum = false
            c = v
        }
    }

    if inNum {
        idx, _ := strconv.Atoi(builder.String())
        off += idx
        arr = append(arr, item{upperIndex: off, c : c})
    }

    return arr
}

func(arr IndexArray)Get(index int) (rune, error) {
    if index < 0 {
        return 0, errors.New("Out of range")
    }

    left := 0
    right := len(arr)

    pos := -1

    for left < right {
        mid := (left + right)/2
        if arr[mid].upperIndex <= index {
            left = mid+1
        } else {
            pos = mid
            right = mid
        }
    }

    if pos == -1 {
        return 0, errors.New("Out of range")
    }

    return arr[pos].c, nil
}

