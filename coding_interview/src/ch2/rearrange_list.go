package ch2list

import (
    "util"
)

// 1,2,3,4--> 1,3,2,4
// 1,2,3,4,5--> 1,3,2,4,5
func RearrangeList(l *util.ListNode) *util.ListNode {
    if l == nil || l.Next == nil {
        return l
    }

    fast := l
    slow := l
    prevSlow := l

    for {
        if fast != nil && fast.Next != nil {
            fast = fast.Next.Next
            prevSlow = slow
            slow = slow.Next
        } else {
            break
        }
    }

    // now slow point to the 2nd half
    prevSlow.Next = nil
    head := l
    for node1, node2 := head, slow; node1 != nil && node2 != nil; {
        next1 := node1.Next
        next2 := node2.Next

        node1.Next = node2
        node2.Next = next1

        // first part maybe shorter
        if next1 != nil {
            node1 = next1
        } else {
            node1.Next.Next = next2
            break
        }

        node2 = next2
    }

    return head

}

