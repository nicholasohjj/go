package main

import (
"fmt"
"sort"
)

func sortedKey(w string) string {
	runes := []rune(w)
	sort.Slice(runes, func(x,y int) bool {
		return runes[x] < runes[y]
	})
	return string(runes)
}

func groupAnagrams(s []string) [][]string {
groups := make(map[string][]string)

for _, v := range s {
	groups[sortedKey(v)] = append(groups[sortedKey(v)], v)
}

res := make([][]string, 0, len(groups))

for _, v := range groups {
	res = append(res, v)
}	

return res
}

func main() {
s := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
res := groupAnagrams(s)
fmt.Println(res)
}


