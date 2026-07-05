package main 

import "fmt"

func unique(s []int) []int {
	res := make([]int, 0, len(s))
	seen := make(map[int]struct{})

	for _, v := range s {
		if _, ok := seen[v]; ok {
			continue
		}
		res = append(res, v)
		seen[v] = struct{}{}	
	}
	return res
}

func main() {
	s := []int{1,2,3,4,5,5}
	s = unique(s)
	fmt.Println(s)
}
