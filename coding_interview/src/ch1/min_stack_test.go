package ch1stackqueue

import (
   "testing"
)

func TestMinStack(t* testing.T) {
    const (
        PUSH = iota
        POP
        MIN
    )

    var tests = []struct {
        action int
        input int
        want int
    } {
        {PUSH, 3, -1},
        {PUSH, 7, -1},
        {MIN, -1, 3},

        {PUSH, 8, -1},
        {PUSH, 1, -1},
        {MIN, -1, 1},

        {PUSH, 1, -1},
        {POP, -1, 1},
        {MIN, -1, 1},
        {POP, -1, 1},
        {MIN, -1, 3},
    }

    s := &MinStack{}
    for _, test := range tests {
        if test.action == PUSH {
            s.Push(test.input)
        } else if test.action == POP {
            if got := s.Pop(); got != test.want {
                t.Errorf("MinStack pop (%q) = %v, want %v", test.input, got, test.want)
            }
        } else if test.action == MIN {
            if got := s.GetMin(); got != test.want {
                t.Errorf("MinStack get min (%q) = %v, want %v", test.input, got, test.want)
            }
        } else {
            t.Errorf("MinStack unknown action (%q)", test.action)
        }
    }
}

