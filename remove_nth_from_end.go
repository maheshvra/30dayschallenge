package main

/*
* Leetcode: https://leetcode.com/problems/remove-nth-node-from-end-of-list/
 */

import (
	"fmt"
)

// ListNode type
type ListNode struct {
	Val  int
	Next *ListNode
}

func seek(node *ListNode, n int) *ListNode {
	var temp *ListNode = node
	for n > 1 {
		temp = temp.Next
		n -= 1
	}
	fmt.Println("temp: ", temp)
	return temp
}

// One pass solution with two pointers
// seek first pointer to n+1 position
// keep second pointer at the beginning
// This will keep distance between first
// and second N nodes apart.
// Increment nodes by one position until
// second pointer reaches the last node

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// Assuming count of nodes equal to N
	assume_last_node := seek(head, n)
	if assume_last_node.Next == nil {
		return head.Next
	}

	// Assuming count of nodes always greater than n
	first := seek(head, n+1)
	second := head
	for {
		if first.Next != nil {
			first = first.Next
			second = second.Next

		} else {
			fmt.Println(second)
			temp := second.Next.Next
			second.Next = temp
			break
		}
	}
	return head
}
