package main

import (
    "fmt"
    "github.com/toravir/leetcode/funcs"
)


func main () {
    var nodes [5]funcs.ListNode
    nodes[0].Val = 1
    nodes[0].Next = &nodes[1]
    nodes[1].Val = 2
    nodes[1].Next = &nodes[2]
    nodes[2].Val = 3
    nodes[2].Next = &nodes[3]
    nodes[3].Val = 4
    nodes[3].Next = &nodes[4]
    nodes[4].Val = 5
    nodes[4].Next = nil

    nw := funcs.ReverseBetween(&nodes[0], 1, 5)
    for nw != nil {
        fmt.Println(nw)
        nw = nw.Next
    }
}
