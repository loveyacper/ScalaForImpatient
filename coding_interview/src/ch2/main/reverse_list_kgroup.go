package main

import (
    "ch2"
    "util"
)

func main() {
    l := util.MakeList(1,2,3,4,5,6,7)
    l = ch2list.ReverseListKGroup(l, 3)
    l.Print()

    l = ch2list.ReverseListKGroup(l, 3)
    l.Print()

    l = util.MakeList(1,2,3)
    l = ch2list.ReverseListKGroup(l, 3)
    l.Print()

    l = util.MakeList(1,2)
    l = ch2list.ReverseListKGroup(l, 3)
    l.Print()
}
