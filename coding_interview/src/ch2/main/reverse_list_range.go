package main

import (
    "ch2"
    "util"
)

func main() {
    l := util.MakeList(1,3,5,7)
    l = ch2list.ReverseListRange(l, 1, 3)
    l.Print()

    l = ch2list.ReverseListRange(l, 1, 3)
    l.Print()
}
