package main

import (
    "ch2"
    "util"
)

func main() {
    l := util.MakeList(1)
    l = ch2list.RearrangeList(l)
    l.Print()

    l = util.MakeList(1,2)
    l = ch2list.RearrangeList(l)
    l.Print()

    l = util.MakeList(1,2,3)
    l = ch2list.RearrangeList(l)
    l.Print()

    l = util.MakeList(1,2,3,4)
    l = ch2list.RearrangeList(l)
    l.Print()

    l = util.MakeList(1,2,3,4,5)
    l = ch2list.RearrangeList(l)
    l.Print()

    l = util.MakeList(1,2,3,4,5,6)
    l = ch2list.RearrangeList(l)
    l.Print()
}

