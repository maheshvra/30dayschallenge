package main

// https://leetcode.com/problems/valid-parentheses/

import (
	"fmt"
	"strings"
)

func isValid(s string) bool {
	s = strings.Trim(s, " ")
	if len(s) == 0 {
		return true
	} else if len(s) < 2 {
		return false
	}

	openParantheses := make(map[string]string)
	openParantheses["{"] = "}"
	openParantheses["["] = "]"
	openParantheses["("] = ")"

	closeParantheses := make(map[string]string)
	closeParantheses["}"] = "{"
	closeParantheses["]"] = "["
	closeParantheses[")"] = "("

	cache := []string{}

	for i := 0; i < len(s); i++ {
		char := string(s[i])
		if _, ok := closeParantheses[char]; ok {
			if len(cache) < 1 {
				return false
			}
			last := cache[len(cache)-1]
			if openVal, ok := openParantheses[last]; ok && char == openVal {
				cache = cache[0 : len(cache)-1]
			} else {
				return false
			}

		} else {
			cache = append(cache, char)
		}
	}
	return len(cache) == 0
}

func main() {
	fmt.Println(isValid("(])"))
}
