package main

import (
    "ch2"
    "util"
)

func main() {
    l := util.MakeList(1,3)
    l = ch2list.ReverseList(l)
    l.Print()

    l = ch2list.RecursiveReverseList(l)
    l.Print()
}
