package main

import (
    "fmt"
)

type segmentTree struct {
    tree []int
}

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
    for i:=n-1;i>=0;i--{
        lc := 2*i+1
        rc := 2*i+2
        fmt.Println("Idx: Parent:", i, " LC:", lc, " RC:", rc)
        if rc > len(st.tree)-1 {
            fmt.Println("No rc")
            st.tree[i] = st.tree[lc]
            continue
        }
        if st.tree[lc] > st.tree[rc] {
            st.tree[i] = st.tree[rc]
        } else {
            st.tree[i] = st.tree[lc]
        }
        fmt.Println("Parent:", st.tree[i], " LC:", st.tree[lc], " RC:", st.tree[rc])
    }
}

func main () {
    fmt.Println("In Segment Tree")
    st := NewSegmentTree([]int{2, 6, 10, 1, 2})
    fmt.Println("St:", st)
}