package main

import (
        "fmt"
)

type fenwickTree struct {
    tree []int
}

func NewFenwickTree (inp []int) fenwickTree {
    ft := fenwickTree{}
    ft.tree = make([]int, 0, len(inp)+1)
    ft.tree = append(ft.tree, 0)
    ft.tree = append(ft.tree, inp...)
    fmt.Println(ft)
    return ft
}

func main () {
    _ = NewFenwickTree([]int{1,2,3,4,5})
}