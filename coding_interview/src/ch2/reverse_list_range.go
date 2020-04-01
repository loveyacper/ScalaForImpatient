package ch2list

import (
    "util"
)

// #92
// func RecursiveReverseList(l *util.ListNode) *util.ListNode

// from to >=1, <= len(l)
func ReverseListRange(l *util.ListNode, from, to int) *util.ListNode {
    if l == nil || l.Next == nil {
        return l
    }

    dist := to - from
    if from < 1 || to < 1 || dist <= 0 {
        return l
    }

    // find the to Node
    endNode := l
    for dist > 0 && endNode != nil{
        dist--
        endNode = endNode.Next
    }

    if dist != 0 {
        return l
    }

    var beforeStart *util.ListNode
    startNode := l

    // find the from Node
    for index := 1; index != from && endNode != nil; index++ {
        beforeStart = startNode
        startNode, endNode = startNode.Next, endNode.Next
    }

    var afterEnd *util.ListNode
    if endNode != nil {
        afterEnd = endNode.Next
        endNode.Next = nil
    }

    // reverse [from, to]
    newSubHead := RecursiveReverseList(startNode)
    newSubTail := startNode

    if newSubHead != endNode {
        panic("wrong head")
    }

    if beforeStart == nil {
        l = newSubHead
    } else {
        beforeStart.Next = newSubHead
    }

    newSubTail.Next = afterEnd
    return l
}

