package ch1stackqueue

import (
   "testing"
)

func TestSubarrayNum(t* testing.T) {
    //
    nums := []int{2,0,1,7,1,1,1,6}
    res := SubarrayNum(nums, 3)

    if res != 14 {
        t.Errorf("TestSubarrayNum failed: encounter %d, but expect 14", res)
    }

    nums = []int{8,1,2,0,9,5}
    res = 10
    if res != 10 {
        t.Errorf("TestSubarrayNum failed: encounter %d, but expect 10", res)
    }
}

