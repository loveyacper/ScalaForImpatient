package ch1stackqueue

import (
   "testing"
)

func TestMaxSlideWindow(t* testing.T) {
    nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
    k := 3
    expect := []int{3, 3, 5, 5, 6, 7}

    res := MaxSlidingWindow(nums, k)

    if len(res) != len(expect) {
        t.Errorf("TestMaxSlideWindow failed without equal length")
    }

    for i := range(expect) {
        if expect[i] != res[i] {
            t.Errorf("TestMaxSlideWindow failed at [%d]: encounter %d, but expect %d", i, res[i], expect[i])
        }
    }
}

