package ch2list

import (
    "util"
)

func RemoveDuplicates(l *util.ListNode) *util.ListNode {
    if l == nil || l.Next == nil {
        return l
    }

    for node := l; node != nil; node = node.Next {
        v := node.Val
        prev := node
        for iter := node.Next; iter != nil; {
            next := iter.Next
            if iter.Val == v {
                prev.Next = next
                iter.Next = nil
            } else {
                prev = iter
            }

            iter = next
        }
    }

    return l
}

