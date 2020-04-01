package main

import (
    "ch2"
    "util"
)

func main() {
    // has 3 parts
    l := util.MakeList(1,2,3,4)
    l = ch2list.PartitionList(l, 3)
    l.Print()

    // has smaller and bigger parts
    l = util.MakeList(9,5,2,7)
    l = ch2list.PartitionList(l, 3)
    l.Print()

    // has smaller parts
    l = util.MakeList(9,5,2,7)
    l = ch2list.PartitionList(l, 10)
    l.Print()

    // has bigger parts
    l = util.MakeList(9,5,2,7)
    l = ch2list.PartitionList(l, 1)
    l.Print()
}
