package ch2list

import (
    "util"
)
// #234
func IsParlindomeList(l *util.ListNode) bool {
    if l == nil || l.Next == nil {
        return true
    }

    // slow will point to the second half
    var middle *util.ListNode
    slow, fast := l, l
    for fast != nil {
        middle = slow
        slow = slow.Next
        fast = fast.Next
        if fast != nil {
            fast = fast.Next
        }
    }

    middle.Next = nil
    subHead := ReverseList(slow)
    isParlindome := true
    for iter1, iter2 := l, subHead; iter2 != nil; iter1, iter2 = iter1.Next, iter2.Next {
        if iter1.Val != iter2.Val {
            isParlindome = false
            break
        }
    }

    // restore list
    middle.Next = ReverseList(subHead)

    return isParlindome
}

