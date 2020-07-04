package main

import "fmt"

func main() {
	fmt.Println(maxProduct([]int{1, 5, 4, 5}))
}

func maxProduct(nums []int) int {
	var first, second int
	index := 0
	bound := len(nums) / 2

	oneList := nums[0:bound]
	twoList := nums[bound:]

	for {
		var one, two int
		if index >= len(oneList) && index >= len(twoList) {
			break
		}
		if index < len(oneList) {
			one = oneList[index]
		}
		if index < len(twoList) {
			two = twoList[index]
		}
		if one < two {
			temp := one
			one = two
			two = temp
		}

		if first < one {
			if two > first {
				second = two
			} else {
				second = first
			}
			first = one
		} else if second < one {
			second = one
		} else if second <= two {
			second = two
		}
		index += 1
	}
	return (first - 1) * (second - 1)
}
