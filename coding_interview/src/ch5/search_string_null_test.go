package ch5string

import (
   "testing"
)

func stringPtr(s string) *string {
    return &s
}

func TestSearchNull(t* testing.T) {
    strs := append([]*string(nil), stringPtr("hello"))
    strs = append(strs, stringPtr("hello"))
    strs = append(strs, nil)
    strs = append(strs, stringPtr("world"))
    strs = append(strs, nil)
    strs = append(strs, stringPtr("world"))

    if i := SearchStringsWithNull(strs, "world"); i != 3 {
        t.Errorf(`should be 3 but %d`, i)
    }
}

func TestSearchNull2(t* testing.T) {
    strs := append([]*string(nil), nil)
    strs = append(strs, stringPtr("hello"))
    strs = append(strs, nil)
    strs = append(strs, nil)
    strs = append(strs, stringPtr("world"))
    strs = append(strs, stringPtr("world"))

    if i := SearchStringsWithNull(strs, "world"); i != 4 {
        t.Errorf(`should be 4 but %d`, i)
    }
}

func TestSearchFail(t* testing.T) {
    strs := append([]*string(nil), nil)
    strs = append(strs, stringPtr("hello"))
    strs = append(strs, nil)
    strs = append(strs, nil)
    strs = append(strs, stringPtr("world"))
    strs = append(strs, stringPtr("world"))

    if i := SearchStringsWithNull(strs, "fuck"); i != -1 {
        t.Errorf(`should be -1 but %d`, i)
    }
}

