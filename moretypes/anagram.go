package main

import (
"fmt"
"slices"
)

func sortedKey(s string) string {
	runes := []rune(s)
	slices.Sort(runes)
	return string(runes)
}

func groupAnagrams(s []string) [][]string {
	groups := make(map[string][]string)
	for _, v := range s {
		key := sortedKey(v)
		groups[key] = append(groups[key], v)
	}
	
	res := make([][]string, 0, len(groups))

	for _, v := range groups {
		res = append(res, v)
	}
	
	return res
	
}


func main() {
	a := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	res := groupAnagrams(a)
	fmt.Println(res)
}
