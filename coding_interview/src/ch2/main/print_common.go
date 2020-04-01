package main

import (
    "ch2"
    "util"
    "fmt"
)

func main() {
    fmt.Println("Hello test list")
    l1 := util.MakeList(1,3,5,7,9)
    l2 := util.MakeList(2,4,5,6,7,8,9)
    l3 := ch2list.CommonList(l1, l2)
    l3.Print()
}
