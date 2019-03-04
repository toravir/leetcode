package funcs

import (
    "fmt"
)

var _ = fmt.Println

/*
https://leetcode.com/problems/remove-nth-node-from-end-of-list/
Given a linked list, remove the n-th node from the end of list and return its head.

 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
    Val int
    Next *ListNode
}
func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
    ret := head
    for i:=0; i < n ; i++ {
        head = head.Next
    }
    if head == nil {
        //++we have to remove the first Node
        return ret.Next
    }
    head = head.Next
    trailer := ret
    for head != nil {
        head = head.Next
        trailer = trailer.Next
    }
    trailer.Next = trailer.Next.Next
    return ret
}

func ReverseBetween(head *ListNode, m int, n int) *ListNode {
    if m == n {
        return head
    }
    ret := ListNode {Val: 0, Next: head}
    head = &ret
    toRev := n - m
    for i:=0 ; i < m-1; i++ {
        head = head.Next
    }
    revStart := head.Next
    var prev *ListNode
    prev = nil
    cur := revStart
    for i:=0; i < toRev+1; i ++ {
        tmp := cur.Next
        cur.Next = prev
        prev = cur
        cur = tmp
    }
    head.Next = prev
    revStart.Next = cur
    return ret.Next
}