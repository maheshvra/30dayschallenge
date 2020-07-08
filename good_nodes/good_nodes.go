package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(currentMax int, root *TreeNode) int {
	var count, leftMax, rightMax int
	leftMax = currentMax
	rightMax = currentMax

	if root.Left != nil && root.Left.Val >= currentMax {
		count += 1
		leftMax = root.Left.Val
		fmt.Println(root.Left.Val)
	}
	if root.Right != nil && root.Right.Val >= currentMax {
		count += 1
		rightMax = root.Right.Val
		fmt.Println(root.Right.Val)
	}
	if root.Left != nil {
		count += countNodes(leftMax, root.Left)
	}
	if root.Right != nil {
		count += countNodes(rightMax, root.Right)
	}
	return count
}

func goodNodes(root *TreeNode) int {
	return countNodes(root.Val, root) + 1
}
