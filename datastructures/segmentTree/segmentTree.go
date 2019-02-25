package main

import (
    "fmt"
)

type segmentTree struct {
    tree []int
}

const INFINITY = 4294967295

func NewSegmentTree(inp []int) segmentTree {
    st := segmentTree{}
    st.createTree(inp)
    return st
}

func (st *segmentTree) createTree (inp []int) {
    n := len(inp)
    st.tree = make([]int, 2*n)
    for i:=0;i<n;i++ {
        st.tree[n+i]=inp[i]
    }
    for i:=n-1;i>=1;i--{
        lc := 2*i
        rc := 2*i+1
        if st.tree[lc] > st.tree[rc] {
            st.tree[i] = st.tree[rc]
        } else {
            st.tree[i] = st.tree[lc]
        }
    }
}

func (st *segmentTree) rangeMin (start, end int) int {
    n := len(st.tree)/2
    start += n
    end += n

    min := INFINITY
    for ; start < end ; start, end = start/2, end/2 {
        if start % 2 == 1 {
            if min > st.tree[start] {
                min = st.tree[start]
            }
            start++
        }
        if end % 2 == 1 {
            end --
            if min > st.tree[end] {
                min = st.tree[end]
            }
        }
    }
    return min
}

func main () {
    fmt.Println("In Segment Tree")
    st := NewSegmentTree([]int{2, 6, 10, 4, 7, 28, 9, 11, 6, 33})
    fmt.Println("St:", st)
    fmt.Println(st.rangeMin(0,5))
    fmt.Println(st.rangeMin(2,6))
}