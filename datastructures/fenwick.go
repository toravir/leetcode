package main

import (
	"fmt"
)

type fenwickTree struct {
	tree []int
}

func getLsb(i int) int {
	return (i & -i)
}
func (ft *fenwickTree) create(inp []int) {
	ft.tree = make([]int, 1, len(inp)+1)
	ft.tree = append(ft.tree, inp...)

	for i := 1; i < len(ft.tree); i++ {
		par := i + getLsb(i)
		if par < len(ft.tree) {
			ft.tree[par] += ft.tree[i]
		}
	}
}

func (ft *fenwickTree) getCumSum(i int) int {
	sum := 0
	for ; i != 0; i = i & ^getLsb(i) {
		sum += ft.tree[i]
	}
	return sum
}

func (ft *fenwickTree) getRangeSum(start, end int) int {
	return ft.getCumSum(end) - ft.getCumSum(start-1)
}

func (ft *fenwickTree) updateArr(index, delta int) {
	i := index
	for ; i != 0; i = i & ^getLsb(i) {
		ft.tree[i] += delta
	}
}
func NewFenwickTree(inp []int) fenwickTree {
	ft := fenwickTree{}
	ft.create(inp)
	return ft
}

func main() {
	ft := NewFenwickTree([]int{3, 4, -2, 7, 3, 11, 5, -8, -9, 2, 4, -8})
	fmt.Println(ft.getRangeSum(3, 6))
    ft.updateArr(4, -3)
	fmt.Println(ft.getRangeSum(3, 6))
}
