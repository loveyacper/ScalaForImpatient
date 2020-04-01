package ch2list

import (
    "util"
)
// #86
func PartitionList(l *util.ListNode, pivot int) *util.ListNode {
    if l == nil || l.Next == nil {
        return l
    }

    var (
        smaller, stail *util.ListNode
        equal, etail *util.ListNode
        bigger, btail *util.ListNode
    )

    for iter := l; iter != nil; iter = iter.Next {
        if iter.Val < pivot {
            if smaller == nil {
                smaller = iter
            } else {
                stail.Next = iter
            }
            stail = iter
        } else if iter.Val == pivot {
            if equal == nil {
                equal = iter
            } else {
                etail.Next = iter
            }
            etail = iter
        } else {
            if bigger == nil {
                bigger = iter
            } else {
                btail.Next = iter
            }
            btail = iter
        }
    }

    head := smaller
    if smaller != nil {
        if equal != nil {
            stail.Next = equal
            etail.Next = bigger
        } else {
            stail.Next = bigger
        }
    } else if equal != nil {
        head = equal
        etail.Next = bigger
    } else {
        head = bigger
    }

    return head
}

