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

	// We may not route index too
	routeIndex := make(map[string]bool)
	for _, path := range paths {
		_, ok := routeIndex[path[0]]
		if !ok {
			routeIndex[path[0]] = true
		} else {
			delete(tailMap, path[0])
		}
		_, ok2 := routeIndex[path[1]]
		if !ok2 {
			routeIndex[path[1]] = true
			tailMap[path[1]] = true
		}
	}
	for key, _ := range tailMap {
		return key
	}
	return ""
}
