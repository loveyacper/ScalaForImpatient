package main

import (
    "ch2"
    "util"
)

func main() {
    l := util.MakeList(1,3,5,7,9)
    l = ch2list.RemoveLastK(l, 4)
    l.Print()
}
