package main

import "fmt"

func main() {
	fmt.Println(maxProduct([]int{3, 4, 5, 2}))
}

func maxProduct(nums []int) int {
	var first, second int
	for _, num := range nums {
		if first < num {
			second = first
			first = num
		} else if second < num {
			second = num
		}
	}
	return (first - 1) * (second - 1)
}
