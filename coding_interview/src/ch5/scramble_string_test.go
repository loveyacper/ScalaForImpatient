package ch5string

import (
   "testing"
)

func TestScramble(t* testing.T) {
    if ok := IsScramble("ccabcbabcbabbbbcbb", "bbbbabccccbbbabcba"); ok {
        t.Errorf(`should be false`)
    }

    /*
    if ok := IsScramble("abbbbbccccbbbabcba", "bbbbabccccbbbabcba"); !ok {
        t.Errorf(`should be true`)
    }
    */
}

