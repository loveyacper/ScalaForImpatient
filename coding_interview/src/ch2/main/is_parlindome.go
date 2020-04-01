package main

import (
    "ch2"
    "util"
    "fmt"
)

func main() {
    l := util.MakeList(1,3,5,7,9)
    fmt.Println(ch2list.IsParlindomeList(l))
    l.Print()

    l = util.MakeList(1,3,1)
    fmt.Println(ch2list.IsParlindomeList(l))
    l.Print()

    l = util.MakeList(1,3,3,1)
    fmt.Println(ch2list.IsParlindomeList(l))
    l.Print()
}

