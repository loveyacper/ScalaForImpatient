package main

import (
    "ch2"
    "util"
)

func main() {
    l := util.MakeList(1,2,3,4)
    l = ch2list.SelectSortList(l)
    l.Print()

    l = util.MakeList(2,5,9,7)
    l = ch2list.SelectSortList(l)
    l.Print()

    l = util.MakeList(2,0,1,7,1,1,1,6)
    l = ch2list.SelectSortList(l)
    l.Print()
}
