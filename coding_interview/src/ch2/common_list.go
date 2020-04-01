package ch2list

import (
    "util"
)

func CommonList(l1, l2 *util.ListNode) *util.ListNode {
    var head *util.ListNode
    var tail *util.ListNode
    for l1 != nil && l2 != nil {
        if l1.Val < l2.Val {
            l1 = l1.Next
        } else if l1.Val > l2.Val {
            l2 = l2.Next
        } else {
            n := util.NewListNode(l1.Val)
            if head == nil {
                head = n
            } else {
                tail.Next = n
            }
            tail = n
            l1, l2 = l1.Next, l2.Next
        }
    }

    return head
}

