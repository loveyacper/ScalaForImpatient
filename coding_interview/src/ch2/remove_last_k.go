package ch2list

import (
    "util"
)
// #19
func RemoveLastK(l *util.ListNode, k int) *util.ListNode {
    if k < 1 || l == nil {
        return l
    }

    probe := l
    for k > 0 && probe != nil {
        k--
        probe = probe.Next
    }

    if k == 0 {
        toDel := l
        pre := toDel
        for probe != nil {
            pre = toDel
            toDel = toDel.Next
            probe = probe.Next
        }

        if toDel == l {
            l = l.Next
        } else {
            pre.Next = toDel.Next
            toDel.Next = nil
        }
    }

    return l
}

