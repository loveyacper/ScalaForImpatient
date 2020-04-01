package ch2list

import (
    "util"
)

// #25

func ReverseListKGroup(l *util.ListNode, k int) *util.ListNode {
    if l == nil || l.Next == nil || k <= 1 {
        return l
    }

    var newHead, newTail *util.ListNode

    subHead := l
    count := 0

    for node := l; node != nil; {
        count++
        if count == k {
            count = 0

            next := node.Next
            node.Next = nil

            newSubHead := RecursiveReverseList(subHead)
            newSubTail := subHead

            if newHead == nil {
                newHead = newSubHead
                newTail = newSubTail
            } else {
                newTail.Next = newSubHead
                newTail = newSubTail
            }

            subHead = next
            node = next
        } else {
            node = node.Next
        }
    }

    if newHead == nil {
        return l
    }

    newTail.Next = subHead
    return newHead
}

