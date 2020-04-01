package main

import (
    "ch2"
    "util"
)

func main() {
    l := util.MakeList(1,2,3,3,4,4,2,1,1)
    l = ch2list.RemoveDuplicates(l)
    l.Print()

    l = util.MakeList(1,1,1,1,1,1)
    l = ch2list.RemoveDuplicates(l)
    l.Print()

    l = util.MakeList(1,2,3)
    l = ch2list.RemoveDuplicates(l)
    l.Print()

    l = util.MakeList(3,3,1,1,3)
    l = ch2list.RemoveDuplicates(l)
    l.Print()
}
