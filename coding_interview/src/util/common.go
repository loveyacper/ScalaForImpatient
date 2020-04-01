package util

import (
    "fmt"
)

type ListNode struct {
    Val int
    Next *ListNode
}

func NewListNode(v int) *ListNode{
    n := &ListNode{}
    n.Val = v

    return n
}

func MakeList(vals ...int) *ListNode {
    var head *ListNode
    var tail *ListNode

    for _, v := range vals {
        n := NewListNode(v)
        if head == nil {
            head = n
            tail = n
        } else {
            tail.Next = n
            tail = n
        }
    }

    return head
}

func (l *ListNode) Print() {
    for iter := l; iter != nil; iter = iter.Next {
        fmt.Printf("%d ", iter.Val)
    }

    if l != nil {
        fmt.Printf("\n")
    }
}

func (l *ListNode) Len() int {
    length := 0
    for iter := l; iter != nil; iter = iter.Next {
        length++
    }

    return length
}

func (l *ListNode) Equal(l2 *ListNode) bool {
    iter1 := l
    iter2 := l2
    for ; iter1 != nil && iter2 != nil; iter1, iter2 = iter1.Next, iter2.Next {
        if iter1.Val != iter2.Val {
            return false
        }
    }

    return iter1 == nil && iter2 == nil
}

