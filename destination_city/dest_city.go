package main

import "fmt"

// https://leetcode.com/problems/destination-city/
func main() {

	paths := [][]string{
		{"jMgaf WaWA", "iinynVdmBz"},
		{" QCrEFBcAw", "wRPRHznLWS"},
		{"iinynVdmBz", "OoLjlLFzjz"},
		{"OoLjlLFzjz", " QCrEFBcAw"},
		{"IhxjNbDeXk", "jMgaf WaWA"},
		{"jmuAYy vgz", "IhxjNbDeXk"}}

	fmt.Println(destCity(paths))
}

func destCity(paths [][]string) string {
	tailMap := make(map[string]bool)
	for _, path := range paths {
		_, ok := tailMap[path[0]]
		if !ok {
			tailMap[path[0]] = false
		} else {
			delete(tailMap, path[0])
		}
		_, ok2 := tailMap[path[1]]
		if !ok2 {
			tailMap[path[1]] = true
		} else {
			delete(tailMap, path[1])
		}
	}
	for key, value := range tailMap {
		if value {
			return key
		}
	}
	return ""
}
