package ch1stackqueue

import (
   "testing"
)

func TestStackSort(t* testing.T) {
    s := NewStack()
    s.Push(3)
    s.Push(7)
    s.Push(8)
    s.Push(1)
    s.Push(2)
    s.Push(0)
    s.Push(9)
    s.Push(5)

    s.Sort()
    if !s.IsSorted() {
        t.Errorf("Not sorted!")
    }
}

