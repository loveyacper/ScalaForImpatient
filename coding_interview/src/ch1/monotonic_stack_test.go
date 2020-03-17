package ch1stackqueue

import (
   "testing"
)

func TestMonotonicStackUnique(t* testing.T) {
    nums := []int{3,4,1,5,6,2,7}
    expect := []IntPair{{-1, 2}, {0, 2}, {-1, -1}, {2, 5}, {3, 5}, {2, -1}, {5, -1}}

    /*
-1 5
0 5
0 5
2 5
2 5
-1 -1
-1 -1
*/
    res := LeftRightMinValues(nums)

    if len(res) != len(expect) {
        t.Errorf("TestMonotonicStack not equal length")
    }

    for i := range(expect) {
        if expect[i] != res[i] {
            t.Errorf("TestMonotonicStack failed at [%d]: encounter %d, but expect %d", i, res[i], expect[i])
        }
    }
}

func TestMonotonicStackRepeat(t* testing.T) {
    nums := []int{3,4,4,5,5,2,2}
    expect := []IntPair{{-1, 5}, {0, 5}, {0, 5}, {2, 5}, {2, 5}, {-1, -1}, {-1, -1}}

    res := LeftRightMinValues(nums)

    if len(res) != len(expect) {
        t.Errorf("TestMonotonicStack not equal length")
    }

    for i := range(expect) {
        if expect[i] != res[i] {
            t.Errorf("TestMonotonicStack failed at [%d]: encounter %d, but expect %d", i, res[i], expect[i])
        }
    }
}

