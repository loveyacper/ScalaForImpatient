package ch2list

import (
    "util"
)

func SelectSortList(l *util.ListNode) *util.ListNode {
    if l == nil || l.Next == nil {
        return l
    }

    head := l
    prevCurr := l
    for curr := l; curr != nil; {
        min := curr
        next := curr.Next

        prevMin := curr
        for iter, prevIter := next, curr; iter != nil; iter = iter.Next {
            if iter.Val < min.Val {
                prevMin = prevIter
                min = iter
            }
            prevIter = iter
        }
        if min != curr {
            // move min before curr
            prevMin.Next = min.Next
            min.Next = curr
            if curr == head {
                head = min
            } else {
                prevCurr.Next = min
            }
        }

        prevCurr = min
        curr = min.Next
    }

    return head
}

