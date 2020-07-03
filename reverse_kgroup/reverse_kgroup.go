package main

import "fmt"

/**
 * Definition for singly-linked list.
**/

// ListNode data structure
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	dataset := []int{1, 2, 3, 4, 5, 6}
	var prev, head *ListNode
	for _, data := range dataset {
		temp := &ListNode{data, nil}
		if head == nil {
			head = temp
			prev = temp
		} else {
			prev.Next = temp
			prev = temp
		}
	}
	head = reverseKGroup(head, 2)
	for {
		fmt.Printf("%+v\n", head)
		if head.Next != nil {
			head = head.Next
		} else {
			break
		}
	}
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k > 1 {
		counter := 0
		var prev *ListNode
		var groupHead *ListNode
		var prevGroupTail *ListNode
		var newHead *ListNode
		var fallback *ListNode
		for {
			counter += 1
			temp := &ListNode{head.Val, nil}

			if counter%k == 0 {
				if newHead == nil {
					newHead = temp
				}
				temp.Next = prev
				counter = 0
				if prevGroupTail != nil {
					prevGroupTail.Next = temp
				}
				prevGroupTail = groupHead
			} else if counter%k == 1 {
				fallback = head
				groupHead = temp
				prev = temp
			} else if prev != nil {
				temp.Next = prev
				prev = temp
			}

			if head.Next != nil {
				head = head.Next
			} else {
				if counter%k > 0 {
					prevGroupTail.Next = fallback
				}
				break
			}

		}
		return newHead
	} else {
		return head
	}
}
