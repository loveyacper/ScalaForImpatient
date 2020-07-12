package ch5string

import (
   "testing"
)

func TestReverseWords(t* testing.T) {
    if s1, s2 := ReverseWords("I love u."), ReverseWords("  I  love  u.  "); s1 != s2 {
        t.Errorf(`should equal: [%v] and [%v]`, s1, s2)
    }

    if s := ReverseWords("    "); len(s) != 0 {
        t.Errorf(`should be 0 : len %v, [%v]`, len(s), s)
    }

    if s := ReverseWords(""); len(s) != 0 {
        t.Errorf(`should be 0 : len %v, [%v]`, len(s), s)
    }

    if s := ReverseWords("abc"); s != "abc" {
        t.Errorf(`should equal to abc : [%v]`, s)
    }
}

