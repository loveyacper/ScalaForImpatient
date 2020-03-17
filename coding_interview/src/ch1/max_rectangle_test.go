package ch1stackqueue

import (
   "testing"
)

func TestMaxRectangle(t* testing.T) {
    //
    nums := []int{3,2,3,0}
    res := MaxRectangle(nums)

    if res != 6 {
        t.Errorf("TestMaxRectangle failed: encounter %d, but expect 6", res)
    }

    //
    nums = []int{3,7,8,1,2,9,5}
    res = MaxRectangle(nums)

    if res != 14 {
        t.Errorf("TestMaxRectangle failed: encounter %d, but expect 14", res)
    }

    //
    nums = []int{2,3}
    res = MaxRectangle(nums)

    if res != 4 {
        t.Errorf("TestMaxRectangle failed: encounter %d, but expect 4", res)
    }
}

