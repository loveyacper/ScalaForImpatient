package ch5string

import (
   "testing"
)

func TestMinInsertions(t* testing.T) {
    var tests = []struct {
        input string
        want int
    } {
        {"zzazz", 0},
        {"mbadm", 2},
        {"leetcode", 5},
        {"g", 0},
        {"no", 1},
    }

    for _, test := range tests {
        if got := MinInsertions(test.input); got != test.want {
            t.Errorf("input %v, got %v, want %v", test.input, got, test.want)
        }
    }
}


