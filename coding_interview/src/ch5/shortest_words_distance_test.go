package ch5string

import (
   "testing"
)

func TestDiffWords(t* testing.T) {
    arr := []string{"a", "c", "b", "d", "x", "a", "b", "y"}
    if d := ShortestWordDistance(arr, "a", "b"); d != 1 {
        t.Errorf(`should be 1 but : [%v]`, d)
    }

    if d := ShortestWordDistance(arr, "y", "a"); d != 2 {
        t.Errorf(`should be 2 but : [%v]`, d)
    }

    if d := ShortestWordDistance(arr, "no", "b"); d != -1 {
        t.Errorf(`should be -1 but : [%v]`, d)
    }
}

func TestSameWords(t* testing.T) {
    arr := []string{"m", "c", "pe", "c", "m", "a", "c"}
    if d := ShortestWordDistance(arr, "m", "m"); d != 4 {
        t.Errorf(`should be 4 but : [%v]`, d)
    }

    if d := ShortestWordDistance(arr, "c", "c"); d != 2 {
        t.Errorf(`should be 2 but : [%v]`, d)
    }
}

