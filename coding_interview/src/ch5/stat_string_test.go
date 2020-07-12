package ch5string

import (
   "testing"
)

func TestStatString(t* testing.T) {
    var tests = []struct {
        input string
        want string
    } {
        {"", ""},
        {"aaa", "a_3"},
        {"aaabbaccc", "a_3_b_2_a_1_c_3"},
        {"abc", "a_1_b_1_c_1"},
        {"你好", "你_1_好_1"},
    }

    overflowIndex := 9999999
    underflowIndex := -1
    for _, test := range tests {
        if got := StatString(test.input); got != test.want {
            t.Errorf("input %s, got %s, want %s", test.input, got, test.want)
        } else {
            indices := NewIndexArray(got)
            rindex := 0
            for _, v := range(test.input) {
                if ch, err := indices.Get(rindex); err != nil || ch != v {
                    t.Errorf("input %s, index %d, got %v, want %v", test.input, rindex, ch, v)
                }
                rindex++
            }

            if _, err := indices.Get(overflowIndex); err == nil {
                t.Errorf("input %s, expect out of range error", test.input)
            }
            if _, err := indices.Get(underflowIndex); err == nil {
                t.Errorf("input %s, expect out of range error", test.input)
            }
        }
    }
}

