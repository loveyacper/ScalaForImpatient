package ch1stackqueue

import (
   "testing"
)

func TestQueue(t* testing.T) {
    const (
        PUSH = iota
        POP
        PEEK
    )

    var tests = []struct {
        action int
        input int
        want int
    } {
        {PUSH, 1, -1},
        {PUSH, 2, -1},
        {PUSH, 3, -1},
        {PEEK, -1, 1},
        {POP, -1, 1},

        {PUSH, 4, -1},
        {PUSH, 5, -1},
        {POP, -1, 2},
        {PEEK, -1, 3},
        {POP, -1, 3},
        {POP, -1, 4},
        {POP, -1, 5},
    }

    q := NewQueue()
    for _, test := range tests {
        if test.action == PUSH {
            q.Push(test.input)
        } else if test.action == POP {
            if got := q.Pop(); got != test.want {
                t.Errorf("Queue pop (%q) = %v, want %v", test.input, got, test.want)
            }
        } else if test.action == PEEK {
            if got := q.Front(); got != test.want {
                t.Errorf("Queue front (%q) = %v, want %v", test.input, got, test.want)
            }
        } else {
            t.Errorf("Queue unknown action (%q)", test.action)
        }
    }
}

