package main

import "fmt"

type Node struct {
	data int
	prev *Node
	next *Node
}

func firstUniqChar(s string) int {
	countMap := make(map[string]int)
	length := len(s)
	for i := 0; i < length/2+1; i++ {
		firstIndex := i
		secondIndex := length - i - 1
		if firstIndex > secondIndex {
			break
		}
		firstChar := string(s[firstIndex])
		secondChar := string(s[secondIndex])

		fval, _ := countMap[firstChar]
		sval, _ := countMap[secondChar]
		if firstChar == secondChar {
			countMap[firstChar] = fval + 2
		} else {
			countMap[firstChar] = fval + 1
			countMap[secondChar] = sval + 1
		}
	}
	for i, char := range s {
		if value, _ := countMap[string(char)]; value == 1 {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(firstUniqChar("dddccdbba"))
}
