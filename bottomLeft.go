/*
https://leetcode.com/problems/find-bottom-left-tree-value

513. Find Bottom Left Tree Value

Given a binary tree, find the leftmost value in the last row of the tree.

*/

package main

import "fmt"

type TreeNode struct {
  Val int
  Left *TreeNode
  Right *TreeNode
}

func findBottomLeftValue(root *TreeNode) int {
  ret := 0
  depth := 0
  doDFS(root, &ret, &depth, 0)
  return ret
}

func doDFS(root *TreeNode, ret *int, maxDepth *int, curDepth int) {
  if (root == nil) {
    return
  }
  if (curDepth > *maxDepth) {
    *ret = root.Val
    *maxDepth = curDepth
  }
  doDFS(root.Left, ret, maxDepth, curDepth+1)
  doDFS(root.Right, ret, maxDepth, curDepth+1)
}

func printTree(root *TreeNode) {
  //Convert Tree to an array and print it 

  arr := make([]int, 15)
  for i := range arr {
    arr[i] = -1
  }
  preOrderTraversal(root, 0, arr)
  fmt.Println(arr)
}

func preOrderTraversal(root *TreeNode, idx int, arr []int) {
  if (root == nil) {
    return
  }
  //fmt.Println("Setting idx", idx, " to:", root.Val)
  arr[idx] = root.Val
/*
  if (len(*arr) < 2*idx+1) {
    // need to realloc
    newarr := make([]int, 2*idx+1+1)
    for i:= range *arr {
      newarr[i] = *arr[i]
    }
    arr = &newarr
  }
*/
  preOrderTraversal(root.Left, 2*idx+1, arr)
  preOrderTraversal(root.Right, 2*idx+2, arr)  
}

func main () {
  tr1 := make([]TreeNode, 3)
  tr1[0].Val = 1
  tr1[0].Left, tr1[0].Right = nil, nil
  
  tr1[2].Val = 3
  tr1[2].Left, tr1[2].Right = nil, nil
  
  tr1[1].Val = 2
  tr1[1].Left, tr1[1].Right = &tr1[0], &tr1[2]
  //printTree(&tr1[1])
  fmt.Println(findBottomLeftValue(&tr1[1]))
  
  tr2 := make([]TreeNode, 3)
  tr2[0].Val = 10
  tr2[0].Left, tr2[0].Right = nil, &tr1[2]
  
  tr2[2].Val = 30
  tr2[2].Left, tr2[2].Right = nil, &tr1[1]
  
  tr2[1].Val = 20
  tr2[1].Left, tr2[1].Right = &tr2[0], &tr2[2]
  //printTree(&tr2[1])  
  fmt.Println(findBottomLeftValue(&tr2[1]))

}
