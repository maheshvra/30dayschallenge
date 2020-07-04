package main

import "regexp"
import "fmt"
import "strconv"

func main() {
	input := []string{"onepiece", "onepiece(1)", "onepiece(2)", "onepiece(3)", "onepiece"}
	fmt.Println(getFolderNames(input))
}

type bound struct {
	current, max int
}

func getFolderNames(names []string) []string {

	result := make([]string, len(names))
	nameMap := make(map[string]bound)
	validPattern := regexp.MustCompile(`^(?P<base>.*)\((?P<id>[0-9]+)\)$`)

	for i, name := range names {
		var base, idString string
		var id int
		if validPattern.MatchString(name) {
			base = validPattern.ReplaceAllString(name, "$base")
			idString = validPattern.ReplaceAllString(name, "$id")
			id, _ = strconv.Atoi(idString)
		} else {
			base = name
		}
		b, ok := nameMap[base]
		if ok {
			// fmt.Println(name, id, b.current, b.max)
			if id <= b.current {
				nameMap[name] = bound{0, 0}
				result[i] = name
			} else if id > b.current {
				nameMap[base] = bound{b.current + 1, id}
				result[i] = name
			} else {
				if b.current+1 >= b.max {
					nameMap[base] = bound{b.max + 1, b.max + 1}
					result[i] = fmt.Sprintf("%s(%d)", base, b.max+1)
				} else {
					//fmt.Println(base, b)
					nameMap[base] = bound{b.current + 1, b.max}
					result[i] = fmt.Sprintf("%s(%d)", base, b.current+1)
				}
			}
		} else {
			nameMap[base] = bound{0, 0}
			result[i] = base
		}
	}
	return result
}
