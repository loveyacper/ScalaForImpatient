package ch2list

import (
    "util"
)

func ReverseList(l *util.ListNode) *util.ListNode {
    if l == nil || l.Next == nil {
        return l
    }

    head := l
    iter := l.Next
    l.Next = nil
    for iter != nil {
        next := iter.Next
        iter.Next = head
        head = iter

        iter = next
    }

    return head
}

func RecursiveReverseList(l *util.ListNode) *util.ListNode {
    if l == nil || l.Next == nil {
        return l
    }

    newList := RecursiveReverseList(l.Next)
    lastNode := l.Next //  第二个节点，变成了tailList的尾节点
    lastNode.Next = l
    l.Next = nil

    return newList
}

