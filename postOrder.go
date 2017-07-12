/*

145. Binary Tree Postorder Traversal
Given a binary tree, return the postorder traversal of its nodes' values.

Note: Recursive solution is trivial, could you do it iteratively?

*/

package main

import "fmt"

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

var stack []*TreeNode

func pushNode (node *TreeNode) {
    if (node != nil) {
        stack = append(stack, node)
    }
}

func popNode () *TreeNode {
    if (!isStackEmpty()) {
        retval := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        return retval
    }
    return nil
}

func peekNode () *TreeNode {
    if (!isStackEmpty()) {
        retval := stack[len(stack)-1]
        return retval
    }
    return nil
}

func isStackEmpty () bool {
    if (len(stack) == 0) {
        return true
    }
    return false
}

func isLeafNode (node *TreeNode) bool {
    if (node.Left == nil && node.Right == nil) {
        return true
    }
    return false
}

func postOrderTraversal (root *TreeNode) []int {
    retVal := make([]int, 0)
    stack = make([]*TreeNode, 0)
    pushNode(root)
    poppedChild := false
    for !isStackEmpty() {
        top := peekNode()
        if (poppedChild || isLeafNode(top)) {
            top = popNode()
            retVal = append(retVal, top.Val)
            nextTop := peekNode()
            if (nextTop != nil && ((nextTop.Right != nil && nextTop.Right == top) ||
                (nextTop.Right == nil && nextTop.Left  == top))) {
                poppedChild = true
            } else {
                poppedChild = false
            }
        } else {
            if (top.Right != nil) {
                pushNode(top.Right)
            }
            if (top.Left != nil) {
                pushNode(top.Left)
            }
            poppedChild = false
        }
    }
    return retVal
}


/* utility functions */
func printTree(root *TreeNode) {
  //Convert Tree to an array and print it 

  arr := make([]int, 15)
  for i := range arr {
    arr[i] = -1
  }
  preOrderTraversal(root, 0, arr)
  fmt.Println(arr)
}

/* utility functions */
func preOrderTraversal(root *TreeNode, idx int, arr []int) {
  if (root == nil) {
    return
  }
  //fmt.Println("Setting idx", idx, " to:", root.Val)
  arr[idx] = root.Val
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
  printTree(&tr1[1])
  fmt.Println(postOrderTraversal(&tr1[1]))

  tr2 := make([]TreeNode, 3)
  tr2[0].Val = 10
  tr2[0].Left, tr2[0].Right = nil, &tr1[2]

  tr2[2].Val = 30
  tr2[2].Left, tr2[2].Right = nil, &tr1[1]

  tr2[1].Val = 20
  tr2[1].Left, tr2[1].Right = &tr2[0], &tr2[2]
  printTree(&tr2[1])
  fmt.Println(postOrderTraversal(&tr2[1]))
  fmt.Println(postOrderTraversal(nil))
}
