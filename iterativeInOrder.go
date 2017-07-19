/*

Non-leetcode

How to do inorder traversal of a binary tree without 
using a stack or using recursion.

*/
package main

import "fmt"

type TreeNode struct {
  Val int
  Left *TreeNode
  Right *TreeNode
}

//we'll use one of the 'nil' right pointers to 
//backtrack - we don't need to use stack or recursion

//This func returns the inOrderPredecessor of a given Node - 
//it stops before hitting a nil or looping back to the original node 
//The looping back can happen if we had previously set the return pointer
func inOrderPred (root *TreeNode) *TreeNode {
  var cur *TreeNode = nil
  if (root.Left != nil) {
    cur = root.Left
    for {
      if ((cur.Right == nil) || (cur.Right == root)) {
        break
      }
      cur = cur.Right
    }
  }
  return cur
}

//logic is simple - 
//for us to get back to a particular node after traversing all left nodes 
//we need to keep a memory - lets store that in the right child of our inorder
//predecessor - we'll clear this temp pointer as soon as we walked back using
//that pointer
func inOrderTraversal (root *TreeNode) {
  cur := root
  for cur != nil {
    if (cur.Left != nil) {
      //If we are going to walk to the left tree - we 
      //need to find the inorder pred and store the back pointer there
      pred := inOrderPred(cur)
      if (pred.Right == nil) {
        pred.Right = cur
        cur = cur.Left
      } else {
        //This the case we have already walked back using the 
        //back pointer - clear the right pointer where we had
        //temporarily stored the walk back.
        pred.Right = nil
        fmt.Print(cur.Val,", ")
        cur = cur.Right
      }
    } else {
      //We can simple walk the right - since left is not there
      fmt.Print(cur.Val,", ")
      cur = cur.Right
    }
  }
  fmt.Println("")
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
  inOrderTraversal(&tr1[1])

  tr2 := make([]TreeNode, 3)
  tr2[0].Val = 10
  tr2[0].Left, tr2[0].Right = nil, &tr1[2]

  tr2[2].Val = 30
  tr2[2].Left, tr2[2].Right = nil, &tr1[1]

  tr2[1].Val = 20
  tr2[1].Left, tr2[1].Right = &tr2[0], &tr2[2]
  printTree(&tr2[1])
  inOrderTraversal(&tr2[1])
/*   20
  10    30
    3      2
         1   3
*/

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
